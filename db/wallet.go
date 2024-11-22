package db

import (
	"database/sql"
	"fmt"
	"go-wallet/api/models"
	"go-wallet/db/postgres"
	"go-wallet/util/convert"
	"go-wallet/util/log"
	"math/big"
	"strings"
	"time"
)

func GetAddressByEmailAndChain(email, chain string) (string, error) {
	chain = strings.ToLower(chain)

	query := []string{
		"SELECT address",
		"FROM user_info",
		"WHERE email = $1 AND chain = $2 LIMIT 1",
	}

	args := []interface{}{
		email,
		chain,
	}

	var address string
	err := postgres.QueryRow(postgres.Compose(query), args, &address)
	if err != nil {
		if postgres.IsRecordNotFoundError(err) {
			return "", nil
		}
		log.Error(err)
		return "", err
	}

	return address, nil
}

func Withdraw(req models.WithdrawRequest) (string, error) {
	chain := strings.ToLower(req.Chain)
	coinType := strings.ToLower(req.CoinType)

	balance, err := GetBalance(req.Email, chain, coinType)
	if err != nil {
		return "0", err
	}

	if convert.Zero().Cmp(convert.ToDecimal(req.Value)) == 0 {
		return "0", fmt.Errorf("value shouldn't be zero")
	}

	bigFloatResp := convert.Zero().Sub(balance, convert.ToDecimal(req.Value))
	cmpResp := convert.Zero().Cmp(bigFloatResp)
	if cmpResp == 0 || cmpResp == 1 {
		return "0", fmt.Errorf("invalid amount")
	}

	err = postgres.Trans(func(sqlTx *sql.Tx) error {
		err := updateUserBalance(sqlTx, req.Email, chain, coinType, bigFloatResp)
		if err != nil {
			return err
		}
		err = insertTxHistory(sqlTx, req.Email, req.Destination, chain, coinType, req.Value)
		return err
	})

	if err == nil {
		log.Infof("%s withdraw %s %s on %s", req.Email, req.Value, req.CoinType, req.Chain)
	} else {
		log.Error(err)
	}

	return convert.BigFloatToString(bigFloatResp), err
}

func Transfer(from, to, chain, coinType, value string) (string, error) {
	chain = strings.ToLower(chain)
	coinType = strings.ToLower(coinType)

	fromBalance, err := GetBalance(from, chain, coinType)
	if err != nil {
		return "0", err
	}

	toBalance, err := GetBalance(to, chain, coinType)
	if err != nil {
		return "0", err
	}

	if convert.Zero().Cmp(convert.ToDecimal(value)) == 0 {
		return "0", fmt.Errorf("value shouldn't be zero")
	}

	fromBigFloatResp := convert.Zero().Sub(fromBalance, convert.ToDecimal(value))
	cmpResp := convert.Zero().Cmp(fromBigFloatResp)
	if cmpResp == 0 || cmpResp == 1 {
		return "0", fmt.Errorf("invalid amount")
	}

	toBigFloatResp := convert.Zero().Add(toBalance, convert.ToDecimal(value))
	toCmpResp := convert.Zero().Cmp(toBigFloatResp)
	if toCmpResp == 0 || toCmpResp == 1 {
		return "0", fmt.Errorf("invalid amount")
	}

	err = postgres.Trans(func(sqlTx *sql.Tx) error {
		err := updateUserBalance(sqlTx, from, chain, coinType, fromBigFloatResp)
		if err != nil {
			return err
		}

		err = updateUserBalance(sqlTx, to, chain, coinType, toBigFloatResp)
		if err != nil {
			return err
		}

		err = insertTxHistory(sqlTx, from, to, chain, coinType, value)
		return err
	})

	if err == nil {
		log.Infof("%s transfer %s %s on %s to %s", from, value, coinType, chain, to)
	} else {
		log.Error(err)
	}

	return convert.BigFloatToString(fromBigFloatResp), err
}

func updateUserBalance(sqlTx *sql.Tx, email, chain, coinType string, newValue *big.Float) error {
	query := []string{
		"UPDATE user_balance",
		"SET balance = $1",
		"WHERE email = $2 AND chain = $3 AND coin_type = $4",
	}

	args := []interface{}{
		convert.BigFloatToString(newValue),
		email,
		chain,
		coinType,
	}

	_, err := sqlTx.Exec(postgres.Compose(query), args...)
	if err == nil {
		log.Infof("update balance:%s, chain:%s, coinType:%s, newvalue:%s", email, chain, coinType, convert.BigFloatToString(newValue))
	} else {
		log.Error(err)
		return err
	}

	return nil
}

func insertTxHistory(sqlTx *sql.Tx, from, to, chain, coinType, value string) error {
	query := []string{
		`INSERT INTO transaction_history ("from", "to", value, chain, coin_type, timestamp)`,
		"VALUES ($1, $2, $3, $4, $5, $6)",
	}

	args := []interface{}{
		from,
		to,
		value,
		chain,
		coinType,
		time.Now().Unix(),
	}

	_, err := sqlTx.Exec(postgres.Compose(query), args...)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func GetBalance(email, chain, coinType string) (*big.Float, error) {
	chain = strings.ToLower(chain)
	coinType = strings.ToLower(coinType)

	query := []string{
		"SELECT balance",
		"FROM user_balance",
		"WHERE email = $1 AND chain = $2 AND coin_type = $3 LIMIT 1",
	}

	args := []interface{}{
		email,
		chain,
		coinType,
	}

	var balance string
	err := postgres.QueryRow(postgres.Compose(query), args, &balance)
	if err != nil {
		if postgres.IsRecordNotFoundError(err) {
			return convert.Zero(), nil
		}
		log.Error(err)
		return convert.Zero(), err
	}

	return convert.ToDecimal(balance), nil
}

func GetBalances(email string) ([]*models.BalanceResponse, error) {
	query := []string{
		"SELECT chain, coin_type, balance AS value",
		"FROM user_balance",
		"WHERE email = $1",
	}

	args := []interface{}{
		email,
	}

	rows, err := postgres.Query(postgres.Compose(query), args...)
	if err != nil {
		log.Error(postgres.Compose(query))
		log.Panic(err)
	}

	defer rows.Close()

	result := []*models.BalanceResponse{}
	for rows.Next() {
		m := &models.BalanceResponse{}

		err := rows.Scan(
			&m.Chain,
			&m.CoinType,
			&m.Value,
		)
		if err != nil {
			log.Panic(err)
		}
		result = append(result, m)
	}

	return result, nil
}

func GetTxHistory(email string) ([]*models.TxHistoryResponse, error) {
	query := []string{
		`SELECT "from", "to", value, chain, coin_type, timestamp`,
		"FROM transaction_history",
		`WHERE "from" = $1 OR "to" = $2`,
	}

	args := []interface{}{
		email,
		email,
	}

	rows, err := postgres.Query(postgres.Compose(query), args...)
	if err != nil {
		log.Error(postgres.Compose(query))
		log.Panic(err)
	}

	defer rows.Close()

	result := []*models.TxHistoryResponse{}
	for rows.Next() {
		m := &models.TxHistoryResponse{}

		err := rows.Scan(
			&m.From,
			&m.To,
			&m.Value,
			&m.Chain,
			&m.CoinType,
			&m.Timestamp,
		)
		if err != nil {
			log.Panic(err)
		}
		result = append(result, m)
	}

	return result, nil
}

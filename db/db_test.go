package db

import (
	"go-wallet/api/models"
	"go-wallet/config"
	"go-wallet/util/log"
	"testing"
)

func init() {
	config.Load(false, false, false)
	log.Init(config.IsDebugMode())
	Init()
}

func Test_GetAddressByEmailAndChain(t *testing.T) {
	tests := []struct {
		email string
		chain string
		want  string
	}{
		{email: "a@gmail.com", chain: "ethereum", want: "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"},
		{email: "b@gmail.com", chain: "ethereum", want: "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"},
		{email: "c@gmail.com", chain: "ethereum", want: "0x90F79bf6EB2c4f870365E785982E1f101E93b906"},
		{email: "d@gmail.com", chain: "ethereum", want: "0x90F79bf6EB2c4f870365E785982E1f101E93b906"},
	}

	for _, tt := range tests {
		if got, _ := GetAddressByEmailAndChain(tt.email, tt.chain); got != tt.want {
			t.Errorf("GetAddressByEmailAndChain() = %s, want %s", got, tt.want)
		}
	}
}

func Test_Withdraw(t *testing.T) {
	tests := []struct {
		req  models.WithdrawRequest
		want string
	}{
		{req: models.WithdrawRequest{Email: "a@gmail.com", Chain: "ethereum",
			CoinType:    "eth",
			Destination: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			Value:       "0.01"},
			want: ""},
		{req: models.WithdrawRequest{Email: "a@gmail.com", Chain: "ethereum",
			CoinType:    "eth",
			Destination: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			Value:       "0"},
			want: ""},
		{req: models.WithdrawRequest{Email: "dd@gmail.com", Chain: "ethereum",
			CoinType:    "eth",
			Destination: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			Value:       "0"},
			want: ""},
		{req: models.WithdrawRequest{Email: "a@gmail.com", Chain: "ethereum",
			CoinType:    "eth",
			Destination: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			Value:       "100"},
			want: ""},
	}

	for _, tt := range tests {
		if got, err := Withdraw(tt.req); err != nil {
			t.Errorf("Withdraw() = %s, want %s", got, tt.want)
		}
	}
}

func Test_Transfer(t *testing.T) {
	tests := []struct {
		from     string
		to       string
		chain    string
		coinType string
		value    string
		want     string
	}{
		{from: "a@gmail.com",
			to:       "b@gmail.com",
			chain:    "ethereum",
			coinType: "eth",
			value:    "0.015",
			want:     ""},
		{from: "aa@gmail.com",
			to:       "b@gmail.com",
			chain:    "ethereum",
			coinType: "eth",
			value:    "0.015",
			want:     ""},
		{from: "a@gmail.com",
			to:       "b@gmail.com",
			chain:    "ethereum",
			coinType: "eth",
			value:    "0",
			want:     ""},
	}

	for _, tt := range tests {
		if got, err := Transfer(tt.from, tt.to, tt.chain, tt.coinType, tt.value); err != nil {
			t.Errorf("Transfer() = %s, want %s", got, tt.want)
		}
	}
}

func Test_GetBalance(t *testing.T) {
	tests := []struct {
		email    string
		chain    string
		coinType string
		want     string
	}{
		{email: "a@gmail.com",
			chain:    "ethereum",
			coinType: "eth",
			want:     ""},
	}

	for _, tt := range tests {
		if got, err := GetBalance(tt.email, tt.chain, tt.coinType); err != nil {
			t.Errorf("GetBalance() = %s, want %s", got, tt.want)
		}
	}
}

func Test_GetBalances(t *testing.T) {
	tests := []struct {
		email string
		want  string
	}{
		{email: "a@gmail.com",
			want: ""},
	}

	for _, tt := range tests {
		if got, err := GetBalances(tt.email); err != nil {
			t.Errorf("GetBalances() = %v, want %s", got, tt.want)
		}
	}
}

func Test_GetTxHistory(t *testing.T) {
	tests := []struct {
		email string
		want  string
	}{
		{email: "a@gmail.com",
			want: ""},
	}

	for _, tt := range tests {
		if got, err := GetTxHistory(tt.email); err != nil {
			t.Errorf("GetTxHistory() = %v, want %s", got, tt.want)
		}
	}
}

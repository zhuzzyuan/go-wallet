package convert

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const decimalPrecision = 1024

// Zero returns zero value of *big.Float.
var Zero = newBigFloat

func ParseUint(s string) uint64 {
	s = strings.TrimPrefix(s, "0x")
	val, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		panic(err)
	}

	return val
}

func StringToHex(raw string) []byte {
	raw = strings.TrimPrefix(raw, "0x")
	bytes, err := hex.DecodeString(raw)
	if err != nil {
		panic(err)
	}

	return bytes
}

func HexToString(raw string) string {
	raw = strings.TrimPrefix(raw, "0x")
	bytesArr, err := hex.DecodeString(raw)
	if err != nil {
		panic(err)
	}

	return string(bytes.Trim(bytesArr, "\x00")[:])
}

func HexToIntegerString(raw string) string {
	raw = strings.TrimPrefix(raw, "0x")
	bytes, err := hex.DecodeString(raw)
	if err != nil {
		panic(err)
	}

	return new(big.Int).SetBytes(bytes).String()
}

// GetAddress returns address with 0x prefix.
func GetAddress(raw string) common.Address {
	raw = strings.TrimPrefix(raw, "0x")

	return common.HexToAddress(raw)
}

func Hex0xPrefix(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return hex
	}

	return "0x" + hex
}

func Trim0xPrefix(hex string) string {
	hex = strings.TrimSpace(hex)
	return strings.TrimPrefix(hex, "0x")
}

// TrimHexPrefixZeros trims "0000001dcd6500" to "1dcd6500".
func TrimHexPrefixZeros(hex string) string {
	return strings.TrimLeft(hex, "0")
}

func ZeroPeddingLeft(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return "0x" + strings.Repeat("0", 64-2-len(hex)) + hex[2:]
	}

	return strings.Repeat("0", 64-len(hex)) + hex
}

func ZeroPeddingRight(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return hex + strings.Repeat("0", 64-2-len(hex))
	}

	return hex + strings.Repeat("0", 64-len(hex))
}

func BytesToBigInt(bytes []byte) *big.Int {
	return new(big.Int).SetBytes(bytes)
}

// ToDecimal returns *big.Float format of given decimal string,
// will return nil if input string is empty.
func ToDecimal(valueStr string) *big.Float {
	value, _ := newBigFloat().SetString(valueStr)
	return value
}

func ToBigInt(valueStr string, base int) (*big.Int, bool) {
	value, ok := new(big.Int).SetString(Trim0xPrefix(valueStr), base)
	return value, ok
}

func MustToBigInt(valueStr string, base int) *big.Int {
	if valueStr == "" {
		return nil
	}

	value, ok := new(big.Int).SetString(Trim0xPrefix(valueStr), base)
	if !ok {
		panic("failed to convert to *big.Int")
	}
	return value
}

func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}

func DecimalFromInt(val *big.Int) *big.Float {
	if val == nil {
		return nil
	}

	return newBigFloat().SetInt(val)
}

// BigFloatToString converts big.Float to string.
func BigFloatToString(value *big.Float) string {
	if value == nil {
		return ""
	}

	valueStr := value.Text('f', 32)
	valueStr = strings.TrimRight(valueStr, "0")
	valueStr = strings.TrimSuffix(valueStr, ".")

	return valueStr
}

func newBigFloat() *big.Float {
	return new(big.Float).SetPrec(decimalPrecision)
}

// AmountReadable returns decimals-formatted amount.
// E.g., 100000000 unit of GAS with 8 decimals will return 1.
func AmountReadable(amount *big.Float, decimals uint8) *big.Float {
	if amount == nil {
		return nil
	}

	decimalsFactor := newBigFloat().SetInt(BigPow(10, int64(decimals)))
	readableAmount := newBigFloat().Quo(amount, decimalsFactor)

	return readableAmount
}

func BigFloatToBigInt(val *big.Float, decimals uint8) *big.Int {
	if val == nil {
		return nil
	}

	newValue := Zero().Mul(val, Zero().SetInt(BigPow(10, int64(decimals))))
	newValue = ToDecimal(newValue.Text('f', 32))

	newInt, _ := newValue.Int(nil)
	return newInt
}

func BigIntReadable(amount *big.Int, decimals uint8) *big.Float {
	return AmountReadable(DecimalFromInt(amount), decimals)
}

func BigIntReadableString(amount *big.Int, decimals uint8) string {
	return BigFloatToString(BigIntReadable(amount, decimals))
}

func Truncate(amount *big.Float, decimals uint8) *big.Float {
	decimalsFactor := newBigFloat().SetInt(BigPow(10, int64(decimals)))

	value := newBigFloat().Mul(amount, decimalsFactor)

	numStr := BigFloatToString(value)
	idx := strings.Index(numStr, ".")
	if idx == -1 {
		idx = len(numStr)
	}
	i, _ := new(big.Int).SetString(numStr[:idx], 10)
	value.SetInt(i)

	return value.Quo(value, decimalsFactor)
}

func TruncateToString(amount *big.Float, decimals uint8) string {
	return BigFloatToString(Truncate(amount, decimals))
}

func FormatBigFloat(val *big.Float, decimals uint8) string {
	if val == nil {
		return "-"
	}

	return BigFloatToString(Truncate(val, decimals))
}

func FormatValue(val *big.Float) string {
	if val == nil {
		return "-"
	}

	return BigFloatToString(Truncate(val, 2))
}

func TruncateFeeValueToString(value *big.Float) string {
	return BigFloatToString(TruncateFeeValue(value))
}

func TruncateFeeValue(value *big.Float) *big.Float {
	if value == nil {
		return nil
	}

	decimals := uint8(6)

	valTemp := Zero().Abs(Zero().Mul(value, Zero().SetInt(BigPow(10, 6))))
	for valTemp.Cmp(Zero().SetInt64(100)) < 0 || valTemp.Cmp(Zero()) == 0 {
		decimals++
		valTemp.Mul(valTemp, Zero().SetInt64(10))
	}

	return Truncate(value, decimals)
}

func SmartFormatBigFloat(val *big.Float) string {
	if val == nil {
		return "-"
	}

	if val.Cmp(Zero()) == 0 {
		return "0"
	}

	decimals := uint8(2)
	valTemp := newBigFloat().Abs(newBigFloat().Mul(val, Zero().SetInt64(100)))
	for valTemp.Cmp(newBigFloat().SetInt64(10)) < 0 || valTemp.Cmp(Zero()) == 0 {
		decimals++
		valTemp.Mul(valTemp, newBigFloat().SetInt64(10))
	}

	result, _ := Truncate(val, decimals).Float64()
	valueStr := fmt.Sprintf("%.*f", decimals, result)
	valueStr = strings.TrimRight(valueStr, "0")
	valueStr = strings.TrimSuffix(valueStr, ".")

	return valueStr
}

// BalanceToString returns a string balance.
func BalanceToString(value interface{}, decimals uint8) string {
	var balance string
	switch value := value.(type) {
	case string:
		balance = BigFloatToString(UintStr2Decimal(value, decimals))
	case *big.Int:
		balance = BigFloatToString(BigIntReadable(value, decimals))
	case *big.Float:
		balance = BigFloatToString(UintStr2Decimal(BigFloatToString(value), decimals))
	case float64:
		v := big.NewFloat(value)
		balance = BigFloatToString(UintStr2Decimal(BigFloatToString(v), decimals))
	default:
		panic("not support this type")
	}
	return balance
}

// UintStr2Decimal returns the *big.Float which is rounded quotient valueStr/10^decimals.
func UintStr2Decimal(valueStr string, decimals uint8) *big.Float {
	return new(big.Float).Quo(ToDecimal(valueStr), big.NewFloat(math.Pow10(int(decimals))))
}

func RemoveZeros(valueStr string) string {
	return BigFloatToString(ToDecimal(valueStr))
}

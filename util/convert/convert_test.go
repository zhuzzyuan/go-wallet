package convert

import (
	"math/big"
	"testing"
)

func TestSmartFormatBigFloat(t *testing.T) {
	testCases := map[*big.Float]string{
		ToDecimal("2"):           "2",
		ToDecimal("2.12345"):     "2.12",
		ToDecimal("2.123456789"): "2.12",
		ToDecimal("0.0066"):      "0.0066",
		ToDecimal("0.00668888"):  "0.0066",
		ToDecimal("5.00668888"):  "5",
		ToDecimal("15.0052"):     "15",
		ToDecimal("2.1"):         "2.1",
		ToDecimal("0"):           "0",
		ToDecimal("0.1"):         "0.1",
		ToDecimal("0.01"):        "0.01",
		ToDecimal("0.001"):       "0.001",
		ToDecimal("0.001234"):    "0.0012",
		ToDecimal("0.0001"):      "0.0001",
		ToDecimal("0.000106"):    "0.0001",
		ToDecimal("0.0001234"):   "0.00012",
		ToDecimal("3.299999"):    "3.29",
		ToDecimal("9.999999"):    "9.99",
		ToDecimal("9.090909"):    "9.09",
		ToDecimal("9.909090"):    "9.9",
		ToDecimal("-9.909090"):   "-9.9",
		ToDecimal("-1"):          "-1",
		ToDecimal("-0.001234"):   "-0.0012",
		ToDecimal("-999.888"):    "-999.88",
		ToDecimal("-1234"):       "-1234",
		ToDecimal("-0.99999"):    "-0.99",
	}

	for input, want := range testCases {
		get := SmartFormatBigFloat(input)
		if get != want {
			t.Errorf("get = %s, want = %s", get, want)
		}
	}
}

func TestFormatBigFloat(t *testing.T) {
	testCases := map[string]string{
		FormatBigFloat(ToDecimal("0.123456789"), 6):            "0.123456",
		FormatBigFloat(ToDecimal("0.987654321"), 6):            "0.987654",
		FormatBigFloat(ToDecimal("0.11111111"), 3):             "0.111",
		FormatBigFloat(ToDecimal("0.999999999"), 0):            "0",
		FormatBigFloat(ToDecimal("0.999999999"), 1):            "0.9",
		FormatBigFloat(ToDecimal("0.999999999"), 9):            "0.999999999",
		FormatBigFloat(ToDecimal("0.999999999"), 18):           "0.999999999",
		FormatBigFloat(ToDecimal("0"), 18):                     "0",
		FormatBigFloat(ToDecimal("999"), 18):                   "999",
		FormatBigFloat(ToDecimal("999"), 9):                    "999",
		FormatBigFloat(ToDecimal("999.9"), 18):                 "999.9",
		FormatBigFloat(ToDecimal("999.99"), 18):                "999.99",
		FormatBigFloat(ToDecimal("999.99"), 0):                 "999",
		FormatBigFloat(ToDecimal("-999.99"), 18):               "-999.99",
		FormatBigFloat(ToDecimal("-999.99"), 0):                "-999",
		FormatBigFloat(ToDecimal("-0.000000000000000001"), 18): "-0.000000000000000001",
		FormatBigFloat(ToDecimal("-0.000000000000000009"), 18): "-0.000000000000000009",
		FormatBigFloat(ToDecimal("0.999"), 18):                 "0.999",
		FormatBigFloat(ToDecimal("88"), 6):                     "88",
		FormatBigFloat(ToDecimal("88.888888888888888888"), 6):  "88.888888",
		FormatBigFloat(ToDecimal("0.1"), 18):                   "0.1",
	}

	for get, want := range testCases {
		if get != want {
			t.Errorf("get=%s, want=%s", get, want)
		}
	}
}

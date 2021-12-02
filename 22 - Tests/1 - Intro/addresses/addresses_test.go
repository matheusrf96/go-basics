package addresses

import "testing"

type testCase struct {
	insertedAddress string
	expectedReturn  string
}

func TestAddresType(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{"Street ABC", "Street"},
		{"Avenue Fábio Santos", "Avenue"},
		{"Road ZXY", "Road"},
		// {"Plaza A", "Street"},
		// {"", "Street"},
		// {"Road 123", "Street"},
		// {"Road 321", "ROAD"},
	}

	for _, tc := range testCases {
		receivedAddressType := AddressType(tc.insertedAddress)

		if receivedAddressType != tc.expectedReturn {
			t.Errorf(
				"The address received (%s) is different from the expected (%s). Case: `%s`",
				receivedAddressType,
				tc.expectedReturn,
				tc.insertedAddress,
			)
		}
	}
}

func TestAny(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("1 is not greater than 2")
	}
}

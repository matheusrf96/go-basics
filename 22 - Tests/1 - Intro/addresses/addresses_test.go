package addresses

import "testing"

func TestAddresType(t *testing.T) {
	addressTest := "Avenue Fábio Santos"
	expectedAddressType := "Avenue"
	receivedAddressType := AddressType(addressTest)

	if expectedAddressType != receivedAddressType {
		t.Errorf(
			"The address received (%s) is different from the expected (%s) :(",
			expectedAddressType,
			receivedAddressType,
		)
	}
}

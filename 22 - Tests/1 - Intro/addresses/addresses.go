package addresses

import "strings"

// AddressType checks if address type is valid
func AddressType(address string) (addrType string) {
	validTypes := []string{"street", "avenue", "road"}
	firstWordAddress := strings.Split(strings.ToLower(address), " ")[0]
	addressHasValidType := false

	for _, aType := range validTypes {
		if aType == firstWordAddress {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return strings.Title(firstWordAddress)
	}

	return "Invalid type"
}

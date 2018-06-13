package main

import "strconv"

// IntToHex convert integer bytes of hexadecima array
func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

package main

import "strconv"

// IntToHex convert integer to array of bytes of hexadecima number
func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

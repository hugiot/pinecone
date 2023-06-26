package random

import "bytes"

const baseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// String concurrent security
func String(length int) string {
	buff := bytes.NewBuffer(nil)
	for i := 0; i < length; i++ {
		index := Int(0, 61)
		buff.WriteByte(baseChars[index])
	}
	return buff.String()
}

package base64

import "encoding/base64"

// Base64Encode - a function to encode bytes to base64 string
func Base64Encode(message []byte) []byte {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return b
}

// Base64Decode - a function to вусщву base64 string to bytes
func Base64Decode(message string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

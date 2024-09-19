package util

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// GenerateOtp generates a random 6-digit OTP
func GenerateOtp() (int, error) {
	const otpLength = 6
	const charset = "0123456789"

	otp := make([]byte, otpLength)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return 0, err
		}
		otp[i] = charset[num.Int64()]
	}

	otpInt, err := strconv.Atoi(string(otp))
	if err != nil {
		return 0, err
	}

	return otpInt, nil
}

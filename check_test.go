package ethcheck

import (
	"testing"
)

func Test_PrivateKeyMatchesAddress(t *testing.T) {
	tests := []struct {
		address, privateKey string
		expected            bool
	}{
		{"0xff4de05706Ad132c221A5a755267AC5FcBF52B44", "1ed25e23e9cc186e8e3a8c02ecd825b40f82ebda4a51d94c5dce911609340bdf", true},
		{"0xff4de05706Ad132c221A5a755267AC5FcBF52B44", "1ed25e23e9cc186e8e3a8c02ecd825b40f82ebda4a51d94c5dce911609340bda", false},
	}

	for _, test := range tests {
		isMatch, err := PrivateKeyMatchesAddress(test.privateKey, test.address)
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err.Error())
		}
		if test.expected != isMatch {
			t.Errorf("expected %t, got %t", test.expected, isMatch)
		}
	}
}

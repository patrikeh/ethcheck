package ethcheck

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeyMatchesAddress checks whether private key derives given ethereum address
func PrivateKeyMatchesAddress(privateKey, addr string) (bool, error) {
	if privateKey[:2] != "0x" {
		privateKey = fmt.Sprintf("0x%s", privateKey)
	}

	res, err := hexutil.Decode(privateKey)
	if err != nil {
		return false, fmt.Errorf("error hex decoding private key - %w", err)
	}

	pk, err := crypto.ToECDSA([]byte(res))
	if err != nil {
		return false, fmt.Errorf("error loading private key - %w", err)
	}

	derived, err := privateKeyToAddress(*pk)
	if err != nil {
		return false, fmt.Errorf("error converting priv key to addr - %w", err)
	}

	return addr == derived, nil
}

func privateKeyToAddress(privateKey ecdsa.PrivateKey) (string, error) {
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("could not cast public key as ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}

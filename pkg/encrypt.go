package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

func securityEncode(password string) (result string) {
	if password != "" {
		var limitLength int

		passLength := len(password)
		shortLength := len(EncryptShortToken)
		longLength := len(EncryptLongToken)

		if passLength > shortLength {
			limitLength = passLength
		} else {
			limitLength = shortLength
		}

		for i := 0; limitLength > i; i++ {
			n1 := 187
			n2 := 187
			if passLength <= i {
				n1 = int(EncryptShortToken[i])
			} else if shortLength <= i {
				n2 = int(password[i])
			} else {
				n1 = int(EncryptShortToken[i])
				n2 = int(password[i])
			}
			result += string(EncryptLongToken[(n1^n2)%longLength])
		}
	}
	return
}

func convertRSAKey(key string) (*rsa.PublicKey, error) {
	var publicKeyBytes []byte
	var err error

	if strings.HasPrefix(key, "-----BEGIN") {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return nil, fmt.Errorf("failed to parse PEM block containing the public key")
		}
		publicKeyBytes = block.Bytes
	} else {
		// 解码Base64编码的公钥字符串
		publicKeyBytes, err = base64.StdEncoding.DecodeString(key)
		if err != nil {
			return nil, errors.New("failed to decode base64 encoded public key")
		}
	}

	// 将字节切片解析为 ASN.1 DER 格式的公钥
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		fmt.Println("Failed to parse ASN.1 DER format public key:", err)
		return nil, err
	}
	// golang x509 limit that RSA Modulus must be Positive

	// 将公钥转换为 RSA 公钥
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Failed to convert to RSA public key")
	}

	return rsaPublicKey, nil
}

func decodeBigInt(s string) (*big.Int, error) {
	z := new(big.Int)
	_, flag := z.SetString(s, 16)
	if !flag {
		return nil, errors.New("invalid hex string")
	}
	return z, nil
}

func newPubKey(n, e string) (*rsa.PublicKey, error) {
	nInt, err := decodeBigInt(n)
	if err != nil {
		return nil, err
	}
	eInt, err := decodeBigInt(e)
	if err != nil {
		return nil, err
	}
	return &rsa.PublicKey{N: nInt, E: int(eInt.Int64())}, nil
}

func encrypt(key *rsa.PublicKey, message string) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(message))
	if err != nil {
		fmt.Println("Failed to encrypt:", err)
		return "", err
	}

	// 将加密后的数据进行Base64编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

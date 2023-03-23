package pkg

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
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
	// 解码Base64编码的公钥字符串
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	// 解析DER编码的公钥
	block, _ := pem.Decode(decodedKey)
	if block == nil {
		return nil, err
	}

	// 解析公钥
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

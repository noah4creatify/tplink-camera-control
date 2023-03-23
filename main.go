package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

const (
	shortToken = "RDpbLfCPsJZ7fiv"
	longToken  = "yLwVl0zKqws7LgKPRQ84Mdt708T1qQ3Ha7xv3H7NyU84p21BriUWBU43odz3iP4rBL3cD02KZciXTysVXiV8ngg6vL48rPJyAUw0HurW20xqxv9aYb4M9wK1Ae0wlro510qXeU07kV57fQMc8L6aLgMLwygtc0F10a0Dg70TOoouyFhdysuRMO51yY5ZlOZZLEal1h0t9YQW0Ko7oBwmCAHoic4HYbUyVeU3sfQ1xtXcPcf1aT303wAQhv66qzW"
)

func securityEncode(password string) (result string) {
	if password != "" {
		var limitLength int

		passLength := len(password)
		shortLength := len(shortToken)
		longLength := len(longToken)

		if passLength > shortLength {
			limitLength = passLength
		} else {
			limitLength = shortLength
		}

		for i := 0; limitLength > i; i++ {
			n1 := 187
			n2 := 187
			if passLength <= i {
				n1 = int(shortToken[i])
			} else if shortLength <= i {
				n2 = int(password[i])
			} else {
				n1 = int(shortToken[i])
				n2 = int(password[i])
			}
			result += string(longToken[(n1^n2)%longLength])
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
func main() {}

package firebase

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenDecoder struct {
	TokenString string
}

func (td *TokenDecoder) Decode() (map[string]interface{}, error) {
	parts := strings.Split(td.TokenString, ".")
	headerJson, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	var header map[string]interface{}
	err = json.Unmarshal(headerJson, &header)
	if err != nil {
		return nil, err
	}
	return header, nil
}

type CertificateParser struct {
	CertString string
}

func (cp *CertificateParser) Parse() (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(cp.CertString))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert.PublicKey.(*rsa.PublicKey), nil
}

type TokenVerifier struct {
	TokenString string
	PublicKey   *rsa.PublicKey
}

func (tv *TokenVerifier) Verify() (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tv.TokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tv.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if time.Unix(claims.Exp, 0).Before(time.Now()) {
			return nil, errors.New("Token is valid. But token is expired.")
		} else {
			return claims, nil
		}
	} else {
		return nil, errors.New("Token is not valid")
	}
}

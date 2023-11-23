package middleware

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Hack-Hack-geek-Vol10/graph-gateway/utils/google"
	"github.com/golang-jwt/jwt"
)

type tokenKey struct{}

const (
	tokenPrefix  = "Bearer"
	authTokenKey = "Authorization"
)

func FirebaseAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(authTokenKey)
		// 未認証の場合はunauthorizedを返す
		if token == "" {
			log.Println("token is empty")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(token, " ")
		if len(authHeaderParts) != 2 {
			log.Println("token is invalid1")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if authHeaderParts[0] != tokenPrefix {
			log.Println("token is invalid2")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		td := &TokenDecoder{tokenString: authHeaderParts[1]}
		header, err := td.Decode()
		if err != nil {
			log.Println("token is invalid3")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		kid := header["kid"].(string)
		certString := google.GoogleJWks[kid].(string)

		cp := &CertificateParser{certString: certString}
		publicKey, err := cp.Parse()
		if err != nil {
			log.Println("token is invalid4")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tv := &TokenVerifier{tokenString: authHeaderParts[1], publicKey: publicKey}
		claims, err := tv.Verify()
		if err != nil {
			log.Println("token is invalid5")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), tokenKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type TokenDecoder struct {
	tokenString string
}

func (td *TokenDecoder) Decode() (map[string]interface{}, error) {
	parts := strings.Split(td.tokenString, ".")
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
	certString string
}

func (cp *CertificateParser) Parse() (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(cp.certString))
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
	tokenString string
	publicKey   *rsa.PublicKey
}

func (tv *TokenVerifier) Verify() (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tv.tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tv.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// if time.Unix(claims.Exp, 0).Before(time.Now()) {
		// 	return nil, errors.New("Token is valid. But token is expired.")
		// } else {
		// 	return claims, nil
		// }
		return claims, nil
	} else {
		return nil, errors.New("Token is not valid")
	}
}

package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/schema-creator/graph-gateway/pkg/firebase"
	"github.com/schema-creator/graph-gateway/pkg/google"
)

const (
	tokenPrefix  = "Bearer"
	authTokenKey = "Authorization"
	TokenKey     = "custom_claims"
)

func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(authTokenKey)

			if c.IsWebSocket() {
				next(c)
			}

			// 未認証の場合はunauthorizedを返す
			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is empty")
			}

			authHeaderParts := strings.Split(token, " ")
			if len(authHeaderParts) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid1")
			}

			if authHeaderParts[0] != tokenPrefix {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid2")
			}

			td := &firebase.TokenDecoder{TokenString: authHeaderParts[1]}
			header, err := td.Decode()
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid3")
			}

			kid := header["kid"].(string)
			certString := google.GoogleJWks[kid].(string)

			cp := &firebase.CertificateParser{CertString: certString}
			publicKey, err := cp.Parse()
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid4")
			}

			tv := &firebase.TokenVerifier{TokenString: authHeaderParts[1], PublicKey: publicKey}
			claims, err := tv.Verify()
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "token is invalid5")
			}

			c.Set(TokenKey, claims)

			return next(c)
		}
	}
}

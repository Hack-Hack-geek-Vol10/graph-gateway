package token

import (
	"time"

	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
)

type Maker interface {
	// トークンを作る
	CreateToken(projectId, userID string, authority model.Auth, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}

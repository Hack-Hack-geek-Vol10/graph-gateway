package token

import (
	"errors"
	"time"

	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
)

var (
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidToken     = errors.New("token is invalid")
	ErrProjectIDIsEmpty = errors.New("project id is empty")
)

type Payload struct {
	ProjectID string     `json:"project_id"`
	Authority model.Auth `json:"authority"`
	IssuedAt  time.Time  `josn:"issuedat"`
	ExpiredAt time.Time  `json:"expiredat"`
}

func NewPayload(projectID string, authority model.Auth, duration time.Duration) (*Payload, error) {
	if len(projectID) == 0 {
		return nil, ErrProjectIDIsEmpty
	}

	payload := &Payload{
		ProjectID: projectID,
		Authority: authority,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

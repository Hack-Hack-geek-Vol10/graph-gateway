package domain

type CreateUserParams struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

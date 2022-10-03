package helpers

import "time"

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUser struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAT time.Time `json:"createdAt"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  *ResponseUser
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

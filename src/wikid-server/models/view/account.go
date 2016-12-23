package view

import "time"

type AccountPostRootRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountPostRootResponse struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	CreatedTime time.Time `json:"createdTime"`
}

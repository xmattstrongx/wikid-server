package account

type PostRootRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostRootResponseModel struct {
	SessionId string `json:"sessionId"`
	AccountId string `json:"accountId"`
}

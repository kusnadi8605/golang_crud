package datastruct

//Token ..
type Token struct {
	Token string `json:"token,omitempty"`
}

//TokenRequest ..
type TokenRequest struct {
	Key string `json:"key,omitempty"`
}

//TokenResponse data
type TokenResponse struct {
	ResponseCode string  `json:"responseCode"`
	ResponseDesc string  `json:"responseDesc"`
	Payload      []Token `json:"payload"`
}

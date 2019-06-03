package datastruct

//PromoResponse data
type MiddlewareResponse struct {
	ResponseCode string  `json:"responseCode"`
	ResponseDesc string  `json:"responseDesc"`
	Payload      []Promo `json:"payload"`
}

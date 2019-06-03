package datastruct

//Promo ..
type Promo struct {
	//* tanda bintang supaya tidak err jika datanya null
	ProductID   int
	ProductName string
	//ProductDescription string
	Picture *string
	///CategoryID         *string
	UnitPrice *string
	//MarketPlaceID      *string
	MarketPlaceName *string
	//CategoryName    *string
	PromoID *int
	//PromoDesc  *string
	PromoStart *string
	PromoEnd   *string
	PromoValue *string
	//TypeID          *int
}

//PromoRequest ..
type PromoRequest struct {
	PromoID string `json:"promoID,omitempty"`
	Page    string `json:"page,omitempty"`
	PerPage string `json:"perPage,omitempty"`
	Order   string `json:"order,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Search  string `json:"search,omitempty"`
}

//PromoResponse data
type PromoResponse struct {
	ResponseCode string  `json:"responseCode"`
	ResponseDesc string  `json:"responseDesc"`
	Payload      []Promo `json:"payload"`
}

//PromDetail ..
type PromoDetail struct {
	//* tanda bintang supaya tidak err jika datanya null
	ProductID          int
	ProductName        string
	ProductDescription string
	Picture            *string
	CategoryID         *string
	UnitPrice          *string
	MarketPlaceID      *string
	MarketPlaceName    *string
	CategoryName       *string
	PromoID            *int
	PromoDesc          *string
	PromoStart         *string
	PromoEnd           *string
	PromoValue         *string
	//TypeID          *int
}

//PromoDetailRequest ..
type PromoDetailRequest struct {
	PromoID string `json:"promoID,omitempty"`
}

//PromoDetailResponse data
type PromoDetailResponse struct {
	ResponseCode string        `json:"responseCode"`
	ResponseDesc string        `json:"responseDesc"`
	Payload      []PromoDetail `json:"payload"`
}

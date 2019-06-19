package handler

import (
	"encoding/json"
	conf "golang_crud/config"
	dts "golang_crud/datastruct"
	mdl "golang_crud/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//GetPromoHandler with paging
func GetPromoHandler(conn *conf.Connection) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqPromo dts.PromoRequest
		var PromoResponse dts.PromoResponse

		logDate := time.Now().Format("20060102")
		conf.SetFilename(conf.Param.LogDir + conf.Param.LogsFile["promo"] + logDate + ".txt")

		body, err := ioutil.ReadAll(req.Body)

		err = json.Unmarshal(body, &reqPromo)
		if err != nil {
			PromoResponse.ResponseCode = "500"
			PromoResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(PromoResponse)

			conf.Logf("Decode Product : %s", err)

			return
		}

		// Initialize paging
		perPage := 10
		offset := 0
		paramPage := reqPromo.Page
		paramSort := reqPromo.Sort
		paramSearch := reqPromo.Search
		paramOrder := reqPromo.Order

		if paramPage != "" {
			paramPage, err := strconv.Atoi(paramPage)

			if err != nil {
				paramPage = 0
				offset = 0
			}

			if paramPage == 0 {
				offset = 0
			}

			offset = perPage * paramPage
		}

		listPromo, err := mdl.GetPromo(conn, paramSearch, paramOrder, paramSort, offset, perPage)

		if err != nil {
			PromoResponse.ResponseCode = "301"
			PromoResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(PromoResponse)

			conf.Logf("Response Product : %v", PromoResponse.ResponseDesc)

			return
		}

		if len(listPromo) < 1 {
			PromoResponse.ResponseCode = "301"
			PromoResponse.ResponseDesc = "data not found"
			json.NewEncoder(w).Encode(PromoResponse)

			conf.Logf("Response Product : %v", PromoResponse.ResponseDesc)

			return
		}

		PromoResponse.ResponseCode = "000"
		PromoResponse.ResponseDesc = "Success"
		PromoResponse.Payload = listPromo
		json.NewEncoder(w).Encode(PromoResponse)

		conf.Logf("Response Product : %v", PromoResponse.ResponseDesc)
	}
}

//GetPromoDetailHandler single product
func GetPromoDetailHandler(conn *conf.Connection) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqPromoDetail dts.PromoDetailRequest
		var PromoDetailResponse dts.PromoDetailResponse

		logDate := time.Now().Format("20060102")
		conf.SetFilename(conf.Param.LogDir + conf.Param.LogsFile["promo"] + logDate + ".txt")

		body, err := ioutil.ReadAll(req.Body)

		err = json.Unmarshal(body, &reqPromoDetail)
		if err != nil {
			PromoDetailResponse.ResponseCode = "500"
			PromoDetailResponse.ResponseDesc = err.Error()

			json.NewEncoder(w).Encode(PromoDetailResponse)

			conf.Logf("Decode Product : %s", err)

			return
		}

		promoID := reqPromoDetail.PromoID

		detailPromo, err := mdl.GetPromoDetail(conn, promoID)

		if err != nil {
			PromoDetailResponse.ResponseCode = "301"
			PromoDetailResponse.ResponseDesc = err.Error()
			PromoDetailResponse.Payload = detailPromo
			json.NewEncoder(w).Encode(PromoDetailResponse)

			conf.Logf("Response Product : %v", PromoDetailResponse.ResponseDesc)

			return
		}

		if len(detailPromo) < 1 {
			PromoDetailResponse.ResponseCode = "301"
			PromoDetailResponse.ResponseDesc = "data not found"
			//PromoDetailResponse.Payload = detailPromo
			json.NewEncoder(w).Encode(PromoDetailResponse)

			conf.Logf("Response Product : %v", PromoDetailResponse.ResponseDesc)

			return
		}

		PromoDetailResponse.ResponseCode = "000"
		PromoDetailResponse.ResponseDesc = "Success"
		PromoDetailResponse.Payload = detailPromo
		json.NewEncoder(w).Encode(PromoDetailResponse)

		conf.Logf("Response Product : %v", PromoDetailResponse.ResponseDesc)
	}
}

package models

import (
	"encoding/json"
	"fmt"
	conf "promo_api/config"
	dts "promo_api/datastruct"
	logger "promo_api/logging"
	utl "promo_api/utils"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

/*
func CountPromo(conn *conf.Connection) (int, error) {
	strPromo := dts.Promo{}

	row, err := conn.Query("select count(1) as  from mtr_Promo")

	defer row.Close()

	row.Scan(&strPromo.Total)
	if err != nil {
		return 0, err
	}

	strCount, _ := strconv.Atoi(strPromo.Total)

	return strCount, nil
}*/

//GetPromo as
func GetPromo(conn *conf.Connection, search string, order string, sort string, offset int, limmit int) ([]dts.Promo, error) {
	arrPromo := []dts.Promo{}
	strPromo := dts.Promo{}

	var arrOrder = [...]string{"a.createdDate", "marketPlaceName", "unitPrice"}

	_order, err := strconv.Atoi(order)
	strQuery := `select 
					a.productId,
					a.productName,
					picture,
					unitPrice,
					marketPlaceName,
					c.promoId,
					promoStart,
					promoEnd,
					promoValue
				from mtr_product a left join mtr_category b
					on a.categoryId=b.categoryId
				join mtr_promo  c
					on a.productId=c.productId
				left join mtr_market_place d
					on c.marketPlaceId=d.marketPlaceId
				where 1=1 `

	if search != "" {
		strQuery += `and PromoName like '%` + search + `%' `
	}

	if order != "" {
		strQuery += "order by " + arrOrder[_order] + " " + sort
	}

	strQuery += " limit ?, ?"
	//fmt.Println(strQuery, " offset :", offset, " limit:", limmit)

	rows, err := conn.Query(strQuery, offset, limmit)

	logger.Logf("Query Promo : %s %v %v", strQuery, offset, limmit)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(
			&strPromo.ProductID,
			&strPromo.ProductName,
			&strPromo.Picture,
			&strPromo.UnitPrice,
			&strPromo.MarketPlaceName,
			&strPromo.PromoID,
			&strPromo.PromoStart,
			&strPromo.PromoEnd,
			&strPromo.PromoValue,
		)

		if err != nil {
			return nil, err
		}

		arrPromo = append(arrPromo, strPromo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return arrPromo, nil
}

//GetPromoDetail ..
func GetPromoDetail(conn *conf.Connection, promoID string) ([]dts.PromoDetail, error) {
	arrPromoDetail := []dts.PromoDetail{}
	strPromoDetail := dts.PromoDetail{}
	//_promoID, err := strconv.Atoi(promoID)

	// key redis
	key := "PROMO" + promoID
	jsonString, err := utl.Get(key)

	fmt.Println("Return redis: ", jsonString)

	// data not found
	if err == redis.ErrNil {
		strQuery := `select 
					a.productId,
					a.productName,
					ProductDescription,
					picture,
					b.categoryId,
					categoryName,
					unitPrice,
					c.marketPlaceId,
					marketPlaceName,
					c.promoId,
					promoDesc,
					promoStart,
					promoEnd,
					promoValue
				from mtr_product a left join mtr_category b
					on a.categoryId=b.categoryId
				join mtr_promo  c
					on a.productId=c.productId
				left join mtr_market_place d
					on c.marketPlaceId=d.marketPlaceId
				where c.promoId='` + promoID + `'`

		rows, err := conn.Query(strQuery)

		logger.Logf("Query Promo  detail: %s", strQuery)

		defer rows.Close()
		if err != nil {
			return nil, err
		}

		for rows.Next() {

			err := rows.Scan(&strPromoDetail.ProductID,
				&strPromoDetail.ProductName,
				&strPromoDetail.ProductDescription,
				&strPromoDetail.Picture,
				&strPromoDetail.CategoryID,
				&strPromoDetail.CategoryName,
				&strPromoDetail.UnitPrice,
				&strPromoDetail.MarketPlaceID,
				&strPromoDetail.MarketPlaceName,
				&strPromoDetail.PromoID,
				&strPromoDetail.PromoDesc,
				&strPromoDetail.PromoStart,
				&strPromoDetail.PromoEnd,
				&strPromoDetail.PromoValue,
			)

			if err != nil {
				return nil, err
			}

			if err = rows.Err(); err != nil {
				return nil, err
			}

			arrPromoDetail = append(arrPromoDetail, strPromoDetail)
		}

	} else if err != nil {
		return nil, err

		// data found in redis
	} else {
		jsonData := []byte(jsonString)
		var err = json.Unmarshal(jsonData, &arrPromoDetail)

		if err != nil {
			return nil, err
		}

		fmt.Println("Return redis: ", arrPromoDetail)
	}

	return arrPromoDetail, nil
}

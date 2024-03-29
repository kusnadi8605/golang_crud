package models

import (
	"crypto/sha256"
	"fmt"
	conf "golang_crud/config"
	dts "golang_crud/datastruct"

	"net/http"
	//utl "golang_crud/model"
)

//GetToken d
func GetToken(conn *conf.Connection, req dts.TokenRequest, channel string, timestamp string) ([]dts.Token, error) {
	arrToken := []dts.Token{}
	strToken := dts.Token{}
	key := req.Key

	h := sha256.New()
	h.Write([]byte(timestamp + channel))

	encryptKey := fmt.Sprintf("%x", h.Sum(nil))

	if key != encryptKey {
		return arrToken, fmt.Errorf("Invalid key or channel")
	}

	strQuery := "SELECT password FROM mtr_channel WHERE channel=?"
	rows, err := conn.Query(strQuery, channel)

	conf.Logf("Query GetToken : %s %s", strQuery, channel)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&strToken.Token)
		if err != nil {
			return nil, err
		}
		//arrToken = append(arrToken, strToken)
	}
	//generate Token
	token, _ := GetTokenJwt(strToken.Token + channel)

	strToken.Token = token
	arrToken = append(arrToken, strToken)

	return arrToken, nil
}

//CheckToken ...
func CheckToken(conn *conf.Connection, r *http.Request, channel string) (bool, error) {
	strToken := dts.Token{}

	strQuery := "SELECT password FROM mtr_channel WHERE channel=?"
	rows, err := conn.Query(strQuery, channel)

	conf.Logf("Query checkToken : %s %s", strQuery, channel)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	//fmt.Println("datanya ", rows)

	for rows.Next() {

		err := rows.Scan(&strToken.Token)
		if err != nil {
			return false, err
		}
		//arrToken = append(arrToken, strToken)
	}

	//check Token
	token, err := ValidTokenJwt(r, strToken.Token+channel)
	//fmt.Println("param token : ", token, " ", err)
	return token, err
}

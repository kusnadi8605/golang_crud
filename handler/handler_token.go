package handler

import (
	"encoding/json"
	conf "golang_crud/config"
	dts "golang_crud/datastruct"
	mdl "golang_crud/models"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//GetTokenHandler return single data
func GetTokenHandler(conn *conf.Connection) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqToken dts.TokenRequest
		var TokenResponse dts.TokenResponse
		channel := req.Header.Get("channel")
		timestamp := req.Header.Get("timestamp")

		logDate := time.Now().Format("20060102")
		conf.SetFilename(conf.Param.LogDir + conf.Param.LogsFile["getToken"] + logDate + ".txt")

		body, err := ioutil.ReadAll(req.Body)
		//ip client
		ip, _, _ := net.SplitHostPort(req.RemoteAddr)
		conf.Logf("Header & Body GetToken: %s %s %s", ip, string(body[:]), req.Header)

		if err != nil {
			panic(err)
		}

		//log.Println(string(body))

		err = json.Unmarshal(body, &reqToken)
		if err != nil {
			TokenResponse.ResponseCode = "500"
			TokenResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(TokenResponse)

			conf.Logf("Decode GetToken : %s", err)

			return
		}

		Token, err := mdl.GetToken(conn, reqToken, channel, timestamp)

		if err != nil {
			TokenResponse.ResponseCode = "301"
			TokenResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(TokenResponse)

			conf.Logf("Response GetToken : %s", TokenResponse.ResponseDesc)

			return
		}

		//fmt.Println(Token)

		if len(Token) < 1 {
			TokenResponse.ResponseCode = "301"
			TokenResponse.ResponseDesc = "data not found"
			json.NewEncoder(w).Encode(TokenResponse)

			conf.Logf("Response GetToken : %s", TokenResponse.ResponseDesc)

			return
		}

		TokenResponse.ResponseCode = "000"
		TokenResponse.ResponseDesc = "Success"
		TokenResponse.Payload = Token

		conf.Logf("Response GetToken : %s", TokenResponse.ResponseDesc)

		json.NewEncoder(w).Encode(TokenResponse)
	}
}

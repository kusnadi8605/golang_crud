package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	conf "promo_api/config"
	hdlr "promo_api/handler"
	logger "promo_api/logging"
	mdw "promo_api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// load config file
	configFile := flag.String("conf", "config/config.yml", "main configuration file")

	flag.Parse()
	conf.LoadConfigFromFile(configFile)

	//logDate := time.Now().Format("20060102")
	//fmt.Println(logDate)
	//logger.SetFilename(conf.Param.LogDir + conf.Param.LogsFile["runningApp"] + logDate + ".txt")

	//logger.Logf("Load Conf File %s ", *configFile)

	conf.RedisDbInit(conf.Param.RedisURL)
	//logger.Logf("Load Redis Conf...")

}

func main() {
	// open mysql connection

	conn, err := conf.New(conf.Param.DBType, conf.Param.DBUrl)

	logger.Logf("Load Database Conf: %s ", conf.Param.DBType)
	logger.Logf("running App on port: %s ", conf.Param.ListenPort)

	if err != nil {
		logger.Logf("Load Database Conf: %s ", err)
		log.Fatal(err)
	}

	http.HandleFunc("/api/token", mdw.Chain(hdlr.GetTokenHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.Logging()))
	http.HandleFunc("/api/promo", mdw.Chain(hdlr.GetPromoHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.Logging(), mdw.IsvalidToken(conn)))
	http.HandleFunc("/api/promoDetail", mdw.Chain(hdlr.GetPromoDetailHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST"), mdw.Logging(), mdw.IsvalidToken(conn)))

	var errors error
	errors = http.ListenAndServe(conf.Param.ListenPort, nil)

	if errors != nil {
		logger.Logf("Unable to start the server: %s ", conf.Param.ListenPort)
		os.Exit(1)
	}
}

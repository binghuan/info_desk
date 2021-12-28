package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_bankHttpDelivery "infodesk/bank/delivery/http"
	_bankHttpDeliveryMiddleware "infodesk/bank/delivery/http/middleware"
	_bankRepo "infodesk/bank/repository/mysql"
	_bankUcase "infodesk/bank/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func createDBConnection(keyOfConfig string) *sql.DB {

	dbHost := viper.GetString(fmt.Sprint(keyOfConfig, ".host"))
	dbPort := viper.GetString(fmt.Sprint(keyOfConfig, ".port"))
	dbUser := viper.GetString(fmt.Sprint(keyOfConfig, ".user"))
	dbPass := viper.GetString(fmt.Sprint(keyOfConfig, ".pass"))
	dbName := viper.GetString(fmt.Sprint(keyOfConfig, ".name"))
	connection4Agent := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Taipei")
	dsn := fmt.Sprintf("%s?%s", connection4Agent, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}

func main() {

	dbConn1 := createDBConnection("database_1")
	defer func() {
		err := dbConn1.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	dbConn2 := createDBConnection("database_2")
	defer func() {
		err := dbConn2.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	middL := _bankHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	bankRepo := _bankRepo.NewMysqlBankRepository(dbConn1, dbConn2)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	bankUsecase := _bankUcase.NewBankUsecase(bankRepo, timeoutContext)
	_bankHttpDelivery.NewBankHandler(e, bankUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

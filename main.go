package main

import (
	"database/sql"
	"fmt"
	"log"

	_middleware "github.com/aqua-farm/config/middleware"
	httpDeliver "github.com/aqua-farm/farm/delivery/http"

	farmRepo "github.com/aqua-farm/farm/repository"
	farmUsecase "github.com/aqua-farm/farm/usecase"
	httpPondDeliver "github.com/aqua-farm/pond/delivery/http"
	pondRepo "github.com/aqua-farm/pond/repository"
	pondUsecase "github.com/aqua-farm/pond/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open(`postgres`, psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	middL := _middleware.InitMiddleware()

	e.Use(middL.CORS)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.DELETE, echo.PATCH, echo.PUT},
	}))

	farmRepo := farmRepo.NewPostgresqlFarmRepository(dbConn)
	pondRepo := pondRepo.NewPostgresqlPondRepository(dbConn)

	farmUsecase := farmUsecase.NewFarmUsecase(farmRepo, pondRepo)
	pondUsecase := pondUsecase.NewPondUsecase(pondRepo, farmUsecase)

	httpDeliver.NewFarmHandler(e, farmUsecase)
	httpPondDeliver.NewPondHandler(e, pondUsecase)

	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint
}

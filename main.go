package main

import (
	"log"
	"paradise-booking/config"
	accounthandler "paradise-booking/modules/account/handler"
	accountstorage "paradise-booking/modules/account/storage"
	accountusecase "paradise-booking/modules/account/usecase"
	mysqlprovider "paradise-booking/provider/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Get config error", err)
		return
	}

	//utils.RunDBMigration(cfg)

	// Declare DB
	db, err := mysqlprovider.NewMySQL(cfg)
	if err != nil {
		log.Fatalln("Can not connect mysql: ", err)
	}

	// declare dependencies
	accountRepo := accountstorage.NewAccountStorage(db)
	accountUseCase := accountusecase.NewUserUseCase(cfg, accountRepo)
	accountHdl := accounthandler.NewAccountHandler(accountUseCase)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.POST("/register", accountHdl.RegisterAccount())
	v1.POST("/login", accountHdl.LoginAccount())
	v1.PATCH("/update/:id", accountHdl.UpdatePersonalInfoAccountById())
	v1.GET("/get-profile", accountHdl.GetAccountByEmail())
	router.Run(":" + cfg.App.Port)
}

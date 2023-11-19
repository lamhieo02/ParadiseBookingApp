package main

import (
	"log"
	"paradise-booking/config"
	"paradise-booking/constant"
	accounthandler "paradise-booking/modules/account/handler"
	accountstorage "paradise-booking/modules/account/storage"
	accountusecase "paradise-booking/modules/account/usecase"
	"paradise-booking/modules/middleware"
	placehandler "paradise-booking/modules/place/handler"
	placestorage "paradise-booking/modules/place/storage"
	placeusecase "paradise-booking/modules/place/usecase"
	mysqlprovider "paradise-booking/provider/mysql"
	"paradise-booking/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Get config error", err)
		return
	}

	// Declare DB
	db, err := mysqlprovider.NewMySQL(cfg)
	if err != nil {
		log.Fatalln("Can not connect mysql: ", err)
	}

	utils.RunDBMigration(cfg)

	// declare dependencies
	accountRepo := accountstorage.NewAccountStorage(db)
	accountUseCase := accountusecase.NewUserUseCase(cfg, accountRepo)
	accountHdl := accounthandler.NewAccountHandler(accountUseCase)

	// pepare for place
	placeRepo := placestorage.NewPlaceStorage(db)
	placeUseCase := placeusecase.NewPlaceUseCase(cfg, placeRepo, accountRepo)
	placeHdl := placehandler.NewPlaceHandler(placeUseCase)
	router := gin.Default()

	// fix error CORS
	configCORS := cors.DefaultConfig()
	configCORS.AllowOrigins = []string{"http://localhost:3000"}
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	router.Use(cors.New(configCORS))

	middlewares := middleware.NewMiddlewareManager(cfg, accountRepo)
	router.Use(middlewares.Recover())

	v1 := router.Group("/api/v1")

	// User
	v1.POST("/register", accountHdl.RegisterAccount())
	v1.POST("/login", accountHdl.LoginAccount())
	v1.PATCH("/account/:id", accountHdl.UpdatePersonalInfoAccountById())
	v1.GET("/profile", accountHdl.GetAccountByEmail())
	v1.GET("/profile/:id", middlewares.RequiredAuth(), accountHdl.GetAccountByID())
	v1.PATCH("/account/role/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.AdminRole), accountHdl.UpdateAccountRoleByID())

	// Place
	v1.POST("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.CreatePlace())
	v1.PUT("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.UpdatePlace())
	v1.GET("/places/:id", placeHdl.GetPlaceByID())
	v1.GET("/places/owner", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.ListPlaceByVendor())
	v1.GET("/places/owner/:vendor_id", placeHdl.ListPlaceByVendorID())
	v1.DELETE("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.DeletePlaceByID())
	v1.GET("/places", placeHdl.ListAllPlace())

	router.Run(":" + cfg.App.Port)
}

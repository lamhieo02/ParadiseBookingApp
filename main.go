package main

import (
	"log"
	"net/http"
	cmdworker "paradise-booking/cmd/worker"
	"paradise-booking/config"
	"paradise-booking/constant"
	accounthandler "paradise-booking/modules/account/handler"
	accountstorage "paradise-booking/modules/account/storage"
	accountusecase "paradise-booking/modules/account/usecase"
	"paradise-booking/modules/middleware"
	placehandler "paradise-booking/modules/place/handler"
	placestorage "paradise-booking/modules/place/storage"
	placeusecase "paradise-booking/modules/place/usecase"
	uploadhandler "paradise-booking/modules/upload/handler"
	uploadusecase "paradise-booking/modules/upload/usecase"
	verifyemailshanlder "paradise-booking/modules/verify_emails/handler"
	verifyemailsstorage "paradise-booking/modules/verify_emails/storage"
	verifyemailsusecase "paradise-booking/modules/verify_emails/usecase"
	"paradise-booking/provider/cache"
	mysqlprovider "paradise-booking/provider/mysql"
	redisprovider "paradise-booking/provider/redis"
	s3provider "paradise-booking/provider/s3"
	"paradise-booking/utils"
	"paradise-booking/worker"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
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

	// Declare redis
	redis, err := redisprovider.NewRedisClient(cfg)
	if err != nil {
		log.Fatalln("Can not connect redis: ", err)
	}

	// declare redis client for asynq
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
	}

	// declare task distributor
	taskDistributor := worker.NewRedisTaskDistributor(&redisOpt)

	// declare dependencies account
	accountRepo := accountstorage.NewAccountStorage(db)
	accountCache := cache.NewAuthUserCache(accountRepo, cache.NewRedisCache(redis))
	accountUseCase := accountusecase.NewUserUseCase(cfg, accountRepo, taskDistributor)
	accountHdl := accounthandler.NewAccountHandler(cfg, accountUseCase)

	// declare verify email usecase
	verifyEmailsSto := verifyemailsstorage.NewVerifyEmailsStorage(db)
	verifyEmailsUseCase := verifyemailsusecase.NewVerifyEmailsUseCase(verifyEmailsSto, accountRepo)
	verifyEmailsHdl := verifyemailshanlder.NewVerifyEmailsHandler(verifyEmailsUseCase)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		cmdworker.RunTaskProcessor(&redisOpt, accountstorage.NewAccountStorage(db), cfg, verifyEmailsUseCase)
	}()
	wg.Wait()

	// declare dependencies

	// pepare for place
	placeRepo := placestorage.NewPlaceStorage(db)
	placeUseCase := placeusecase.NewPlaceUseCase(cfg, placeRepo, accountRepo)
	placeHdl := placehandler.NewPlaceHandler(placeUseCase)

	// upload file to s3
	s3Provider := s3provider.NewS3Provider(cfg)
	uploadUC := uploadusecase.NewUploadUseCase(cfg, s3Provider)
	uploadHdl := uploadhandler.NewUploadHandler(cfg, uploadUC)

	router := gin.Default()

	// config CORS
	configCORS := setupCors()
	router.Use(cors.New(configCORS))

	middlewares := middleware.NewMiddlewareManager(cfg, accountCache)
	router.Use(middlewares.Recover())

	v1 := router.Group("/api/v1")

	// health check
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Hello": "World"})
	})
	v1.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// User
	v1.POST("/register", accountHdl.RegisterAccount())
	v1.POST("/login", accountHdl.LoginAccount())
	v1.PATCH("/account/:id", accountHdl.UpdatePersonalInfoAccountById())
	v1.GET("/profile", accountHdl.GetAccountByEmail())
	v1.GET("/profile/:id", accountHdl.GetAccountByID())
	v1.GET("/accounts", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.AdminRole), accountHdl.GetAllAccountUserAndVendor())
	v1.PATCH("/account/role/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.AdminRole), accountHdl.UpdateAccountRoleByID())
	v1.POST("/change/password", middlewares.RequiredAuth(), accountHdl.ChangePassword())
	v1.POST("/change/status", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.AdminRole), accountHdl.ChangeStatusAccount())

	// Place
	v1.POST("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.CreatePlace())
	v1.PUT("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.UpdatePlace())
	v1.GET("/places/:id", placeHdl.GetPlaceByID())
	v1.GET("/places/owner", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.ListPlaceByVendor())
	v1.GET("/places/owner/:vendor_id", placeHdl.ListPlaceByVendorID())
	v1.DELETE("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.DeletePlaceByID())
	v1.GET("/places", placeHdl.ListAllPlace())

	// verify email
	v1.GET("/verify_email", verifyEmailsHdl.CheckVerifyCodeIsMatching())

	// upload file to s3
	v1.POST("/upload", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole, constant.UserRole, constant.AdminRole), uploadHdl.UploadFile())

	// google login
	//v1.GET("/google/login")
	router.Run(":" + cfg.App.Port)
}

func setupCors() cors.Config {
	configCORS := cors.DefaultConfig()
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
	configCORS.AllowOrigins = []string{"http://localhost:3000"}

	return configCORS
}

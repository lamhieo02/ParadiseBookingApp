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
	amenityhandler "paradise-booking/modules/amenity/handler"
	amenitystorage "paradise-booking/modules/amenity/storage"
	amenityusecase "paradise-booking/modules/amenity/usecase"
	bookinghandler "paradise-booking/modules/booking/handler"
	bookingstorage "paradise-booking/modules/booking/storage"
	bookingusecase "paradise-booking/modules/booking/usecase"
	bookingdetailstorage "paradise-booking/modules/booking_detail/storage"
	bookingguiderhandler "paradise-booking/modules/booking_guider/handler"
	bookingguiderstorage "paradise-booking/modules/booking_guider/storage"
	bookingguiderusecase "paradise-booking/modules/booking_guider/usecase"
	bookingratinghandler "paradise-booking/modules/booking_rating/handler"
	bookingratingstorage "paradise-booking/modules/booking_rating/storage"
	bookingratingusecase "paradise-booking/modules/booking_rating/usecase"
	calendarguiderhandler "paradise-booking/modules/calendar_guider/handler"
	calendarguiderstorage "paradise-booking/modules/calendar_guider/storage"
	calendarguiderusecase "paradise-booking/modules/calendar_guider/usecase"
	commenthandler "paradise-booking/modules/comment/handler"
	commentstorage "paradise-booking/modules/comment/storage"
	commentusecase "paradise-booking/modules/comment/usecase"
	likepostreviewhandler "paradise-booking/modules/like_post_review/handler"
	likepostreviewstorage "paradise-booking/modules/like_post_review/storage"
	likepostreviewusecase "paradise-booking/modules/like_post_review/usecase"
	mediahandler "paradise-booking/modules/media/handler"
	mediausecase "paradise-booking/modules/media/usecase"
	"paradise-booking/modules/middleware"
	paymenthandler "paradise-booking/modules/payment/handler"
	paymentstorage "paradise-booking/modules/payment/store"
	paymentusecase "paradise-booking/modules/payment/usecase"
	placehandler "paradise-booking/modules/place/handler"
	placestorage "paradise-booking/modules/place/storage"
	placeusecase "paradise-booking/modules/place/usecase"
	placewishlisthandler "paradise-booking/modules/place_wishlist/handler"
	placewishliststorage "paradise-booking/modules/place_wishlist/storage"
	placewishlistusecase "paradise-booking/modules/place_wishlist/usecase"
	policieshandler "paradise-booking/modules/policy/handler"
	policiesstorage "paradise-booking/modules/policy/storage"
	policiesusecase "paradise-booking/modules/policy/usecase"
	postguidehandler "paradise-booking/modules/post_guide/handler"
	postguidestorage "paradise-booking/modules/post_guide/storage"
	postguideusecase "paradise-booking/modules/post_guide/usecase"
	postreviewhandler "paradise-booking/modules/post_review/handler"
	postreviewstorage "paradise-booking/modules/post_review/storage"
	postreviewusecase "paradise-booking/modules/post_review/usecase"
	replycommenthandler "paradise-booking/modules/reply_comment/handler"
	replycommentstorage "paradise-booking/modules/reply_comment/storage"
	replycommentusecase "paradise-booking/modules/reply_comment/usecase"
	requestguiderhandler "paradise-booking/modules/request_guider/handler"
	requestguiderstorage "paradise-booking/modules/request_guider/storage"
	requestguiderusecase "paradise-booking/modules/request_guider/usecase"
	verifyemailshanlder "paradise-booking/modules/verify_emails/handler"
	verifyemailsstorage "paradise-booking/modules/verify_emails/storage"
	verifyemailsusecase "paradise-booking/modules/verify_emails/usecase"
	wishlisthandler "paradise-booking/modules/wishlist/handler"
	wishliststorage "paradise-booking/modules/wishlist/storage"
	wishlistusecase "paradise-booking/modules/wishlist/usecase"
	"paradise-booking/provider/cache"
	googlemapprovider "paradise-booking/provider/googlemap"
	mediaprovider "paradise-booking/provider/media"
	momoprovider "paradise-booking/provider/momo"
	mysqlprovider "paradise-booking/provider/mysql"
	redisprovider "paradise-booking/provider/redis"
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

	// utils.RunDBMigration(cfg)

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

	// cache residis
	cacheRedis := cache.NewRedisCache(redis)

	// declare task distributor
	taskDistributor := worker.NewRedisTaskDistributor(&redisOpt)

	// google map
	googleMap := googlemapprovider.NewGoogleMap(cfg)

	// momo
	momo := momoprovider.NewMomo(cfg)

	// media
	mediaProvider := mediaprovider.NewMediaProvider(cfg)

	// declare dependencies account
	accountSto := accountstorage.NewAccountStorage(db)
	accountCache := cache.NewAuthUserCache(accountSto, cacheRedis)

	// declare verify email usecase
	verifyEmailsSto := verifyemailsstorage.NewVerifyEmailsStorage(db)
	verifyEmailsUseCase := verifyemailsusecase.NewVerifyEmailsUseCase(verifyEmailsSto, accountSto)
	verifyEmailsHdl := verifyemailshanlder.NewVerifyEmailsHandler(verifyEmailsUseCase)

	accountUseCase := accountusecase.NewUserUseCase(cfg, accountSto, verifyEmailsUseCase, taskDistributor, cacheRedis)
	accountHdl := accounthandler.NewAccountHandler(cfg, accountUseCase)

	// prepare for placewishlist storeage
	placeWishListSto := placewishliststorage.NewPlaceWishListStorage(db)
	// declare cache for place_wishlist
	placeWishListCache := cache.NewPlaceWishListCache(placeWishListSto, cacheRedis)

	// declare dependencies
	// prepare for wish list
	wishListSto := wishliststorage.NewWishListStorage(db)
	wishListUseCase := wishlistusecase.NewWishListUseCase(wishListSto, placeWishListSto, cacheRedis)
	wishListHdl := wishlisthandler.NewWishListHandler(wishListUseCase)

	// prepare for payment
	paymentSto := paymentstorage.NewPaymentStorage(db)
	paymentUC := paymentusecase.NewPaymentUseCase(paymentSto)
	paymentHdl := paymenthandler.NewPaymentHandler(paymentUC)
	// prepare for place
	bookingSto := bookingstorage.NewBookingStorage(db)

	placeSto := placestorage.NewPlaceStorage(db)
	placeCache := cache.NewPlaceStoCache(placeSto, cacheRedis)
	placeUseCase := placeusecase.NewPlaceUseCase(cfg, placeSto, accountCache, googleMap, placeWishListCache, placeCache, bookingSto)
	placeHdl := placehandler.NewPlaceHandler(placeUseCase)

	// prepare for booking detail
	bookingDetailSto := bookingdetailstorage.NewBookingDetailStorage(db)

	// prepare for booking
	bookingUseCase := bookingusecase.NewBookingUseCase(bookingSto, bookingDetailSto, cfg, taskDistributor, accountSto, placeSto, momo, paymentSto)
	bookingHdl := bookinghandler.NewBookingHandler(bookingUseCase)

	// prepare place wish list
	placeWishListUseCase := placewishlistusecase.NewPlaceWishListUseCase(placeWishListSto, placeSto, cacheRedis)
	placeWishListHdl := placewishlisthandler.NewPlaceWishListHandler(placeWishListUseCase)

	// prepare for place rating
	postGuideSto := postguidestorage.NewPostGuideStorage(db)

	bookingRatingSto := bookingratingstorage.Newbookingratingstorage(db)
	bookingRatingUC := bookingratingusecase.Newbookingratingusecase(bookingRatingSto, accountSto, placeSto, cacheRedis, postGuideSto)
	bookingRatingHdl := bookingratinghandler.Newbookingratinghandler(bookingRatingUC)

	// prepare for amenities
	amenitySto := amenitystorage.NewAmenityStorage(db)
	amenityUC := amenityusecase.NewAmenityUseCase(amenitySto, cfg)
	amenityHdl := amenityhandler.NewAmenityHandler(amenityUC)

	// upload file
	mediaUC := mediausecase.NewMediaUseCase(cfg, mediaProvider)
	mediaHdl := mediahandler.NewMediaHandler(cfg, mediaUC)

	// prepare for policy
	policySto := policiesstorage.NewPolicyStorage(db)
	policyUC := policiesusecase.NewPolicyUseCase(policySto)
	policyHdl := policieshandler.NewPolicyHandler(policyUC)

	// prepare for comment
	commentSto := commentstorage.NewCommentStorage(db)

	// prepare for like post review
	likePostReviewSto := likepostreviewstorage.NewLikePostReviewStorage(db)
	likePostReviewUC := likepostreviewusecase.NewLikePostReviewUseCase(likePostReviewSto)
	likePostReviewHdl := likepostreviewhandler.NewLikePostReviewHandler(likePostReviewUC)

	// declare cache for like_post_review
	// likePostReviewCache := cache.NewLikePostReviewStoCache(likePostReviewSto, cacheRedis)

	// declare cache for comment
	// commentCache := cache.NewCommentStoCache(commentSto, cacheRedis)
	// prepare reply comment
	replyCommentSto := replycommentstorage.NewReplyCommentStorage(db)
	replyCommentUC := replycommentusecase.NewReplyCommentUsecase(replyCommentSto, commentSto)
	replyCommentHdl := replycommenthandler.NewReplyCommentHandler(replyCommentUC)

	// prepare for post review
	postReviewSto := postreviewstorage.NewPostReviewStorage(db)
	postReviewUC := postreviewusecase.NewPostReviewUseCase(postReviewSto, commentSto, accountCache, likePostReviewSto, replyCommentSto, *googleMap)
	postReviewHdl := postreviewhandler.NewPostReviewHandler(postReviewUC)

	// declare cache for comment
	commentUC := commentusecase.NewCommentUseCase(commentSto, replyCommentSto, accountCache)
	commentHdl := commenthandler.NewCommentHandler(commentUC)

	// declare for post_guide
	postGuideCache := cache.NewPostGuideStoCache(postGuideSto, cacheRedis)
	postGuideUC := postguideusecase.NewPostGuideUsecase(postGuideSto, postGuideCache, accountCache, *googleMap, cacheRedis)
	postGuideHdl := postguidehandler.NewPostGuideHandler(postGuideUC)

	// declare for calendar guider
	bookingGuiderSto := bookingguiderstorage.NewBookingGuiderStorage(db)
	calendarGuiderSto := calendarguiderstorage.NewCalendarGuiderStorage(db)
	calendarGuiderUC := calendarguiderusecase.NewCalendarGuiderUseCase(calendarGuiderSto, bookingGuiderSto)
	calendarGuiderHdl := calendarguiderhandler.NewCalendarGuiderHandler(calendarGuiderUC)

	// declare for booking guider
	bookingGuiderUC := bookingguiderusecase.NewBookingGuiderUseCase(bookingGuiderSto, taskDistributor, momo, paymentSto, calendarGuiderSto, postGuideUC)
	bookingGuiderHdl := bookingguiderhandler.NewBookingGuiderHandler(bookingGuiderUC)

	// declare for request guider
	requestGuiderSto := requestguiderstorage.NewRequestGuiderStorage(db)
	requestGuiderUC := requestguiderusecase.NewRequestGuiderUC(requestGuiderSto, accountSto)
	requestGuiderHdl := requestguiderhandler.NewRequestGuiderHandler(requestGuiderUC)

	// run task processor
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmdworker.RunTaskProcessor(&redisOpt, accountSto, cfg, verifyEmailsUseCase, bookingSto, bookingUseCase)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmdworker.RunTaskScheduler(&redisOpt, cfg)
	}()

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
	v1.POST("/forgot/password", accountHdl.ForgotPassword())
	v1.POST("/reset/password", accountHdl.ResetPassword())

	// Place
	v1.POST("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.CreatePlace())
	v1.PUT("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.UpdatePlace())
	v1.GET("/places/:id", placeHdl.GetPlaceByID())
	v1.GET("/places/owner", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.ListPlaceByVendor())
	v1.GET("/places/owner/:vendor_id", placeHdl.ListPlaceByVendorID())
	v1.DELETE("/places", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), placeHdl.DeletePlaceByID())
	v1.POST("/places/list", placeHdl.ListAllPlace())
	v1.GET("/places/dates_booked", placeHdl.GetDatesBookedPlace())
	v1.GET("/places/check_date_available", placeHdl.CheckDateBookingAvailable())
	v1.GET("/places/status_booking", placeHdl.GetStatusPlaceToBook())

	// booking
	v1.POST("/bookings", bookingHdl.CreateBooking())
	v1.GET("/confirm_booking", bookingHdl.ConfirmStatusBooking())
	v1.POST("/booking_list", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole, constant.GuiderRole), bookingHdl.ListBooking())
	v1.GET("/bookings/:id", bookingHdl.GetBookingByID())
	v1.GET("/bookings", middlewares.RequiredAuth(), bookingHdl.GetBookingByPlaceID())
	v1.GET("/bookings_list/manage_reservation", middlewares.RequiredAuth(), bookingHdl.ListBookingNotReservation())
	v1.DELETE("/bookings/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole, constant.GuiderRole), bookingHdl.DeleteBookingByID())
	v1.POST("/cancel_booking", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole, constant.GuiderRole), bookingHdl.CancelBookingByID())
	v1.PUT("/bookings", middlewares.RequiredAuth(), bookingHdl.UpdateStatusBooking())

	// wish list
	v1.POST("/wish_lists", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.CreateWishList())
	v1.GET("/wish_lists/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.GetWishListByID())
	v1.GET("/wish_lists/user/:user_id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.GetWishListByUserID())
	v1.PUT("/wish_lists/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.UpdateWishListByID())
	v1.DELETE("/wish_lists/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), wishListHdl.DeleteWishListByID())

	// place wish list
	v1.POST("/place_wish_lists", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), placeWishListHdl.CreatePlaceWishList())
	v1.DELETE("/place_wish_lists/:place_id/:wishlist_id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), placeWishListHdl.DeletePlaceWishList())
	v1.GET("/place_wish_lists/place", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole), placeWishListHdl.ListPlaceByWishListID())

	// booking rating
	v1.POST("/booking_ratings", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.UserRole, constant.VendorRole, constant.GuiderRole), bookingRatingHdl.MakeComment())
	v1.GET("/booking_ratings/comments", bookingRatingHdl.GetCommentByObjectID())
	v1.GET("/booking_ratings/bookings", bookingRatingHdl.GetCommentByBookingID())
	v1.GET("/booking_ratings/users", bookingRatingHdl.GetCommentByUserID())
	v1.GET("/booking_ratings/vendors", bookingRatingHdl.GetCommentByVendorID())
	v1.GET("/booking_ratings/statistics", bookingRatingHdl.GetStatisTicsByObjectId())

	// verify email
	v1.GET("/verify_email", verifyEmailsHdl.CheckVerifyCodeIsMatching())

	// verify reset code password
	v1.GET("/verify_reset_password", verifyEmailsHdl.CheckResetCodePasswordIsMatching())

	// upload file to s3
	v1.POST("/upload", middlewares.RequiredAuth(), mediaHdl.UploadFile())
	v1.GET("/images/:path", mediaHdl.GetImage())

	// amenities
	v1.POST("/amenities", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole, constant.GuiderRole), amenityHdl.CreateAmenity())
	v1.DELETE("/amenities/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole, constant.GuiderRole), amenityHdl.DeleteAmenityByID())
	v1.GET("/amenities/config", amenityHdl.GetAllConfigAmenity())
	v1.GET("/amenities/object", amenityHdl.ListAmenityByObjectId())
	v1.POST("/amenities/object/remove", amenityHdl.DeleteAmenityByListID())

	// policies
	v1.POST("/policies", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole, constant.GuiderRole), policyHdl.UpsertPolicy())
	v1.GET("/policies", policyHdl.GetPolicyByObjectId())

	// payment
	v1.POST("/payments/list_by_vendor", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.VendorRole), paymentHdl.ListPaymentByVendorID())

	// post review
	v1.POST("/post_reviews", middlewares.RequiredAuth(), postReviewHdl.CreatePostReview())
	v1.PUT("/post_reviews", middlewares.RequiredAuth(), postReviewHdl.UpdatePostReview())
	v1.POST("/post_reviews/list/:account_id", postReviewHdl.ListPostReviewByAccountID())
	v1.DELETE("/post_reviews/:post_review_id", middlewares.RequiredAuth(), postReviewHdl.DeletePostReviewByID())
	v1.GET("/post_reviews/:post_review_id", postReviewHdl.GetPostReviewByID())
	v1.POST("/post_reviews/list", postReviewHdl.ListPostReviewByFilter())

	// post review rating
	v1.POST("/post_review/comment", middlewares.RequiredAuth(), postReviewHdl.CommentPostReview())

	// like post review
	v1.POST("/like_post_reviews", middlewares.RequiredAuth(), likePostReviewHdl.LikePostReview())

	// reply comment
	v1.POST("/reply_comments", middlewares.RequiredAuth(), replyCommentHdl.ReplySourceComment())
	v1.DELETE("/reply_comments/:reply_comment_id", middlewares.RequiredAuth(), replyCommentHdl.DeleteReplyComment())

	// comment
	v1.DELETE("/comments/:comment_id", middlewares.RequiredAuth(), commentHdl.DeleteCommentByID())
	v1.GET("/comments/:post_review_id", commentHdl.GetCommentByPostReviewID())

	// post guide
	v1.POST("/post_guides", middlewares.RequiredAuth(), postGuideHdl.CreatePostGuide())
	v1.GET("/post_guides/:id", postGuideHdl.GetPostGuideByID())
	v1.GET("/post_guides/list", postGuideHdl.ListPostGuideByFilter())
	v1.PUT("/post_guides", middlewares.RequiredAuth(), postGuideHdl.UpdatePostGuideByID())
	v1.DELETE("/post_guides/:id", middlewares.RequiredAuth(), postGuideHdl.DeletePostGuideByID())

	// calendar guider
	v1.POST("/calendar_guiders", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.GuiderRole, constant.AdminRole), calendarGuiderHdl.CreateCalendarGuider())
	v1.GET("/calendar_guiders/:id", calendarGuiderHdl.GetCalendarGuiderByID())
	v1.POST("/calendar_guiders/list", calendarGuiderHdl.ListCalendarGuiderByFilter())
	v1.DELETE("/calendar_guiders/:id", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.GuiderRole, constant.AdminRole), calendarGuiderHdl.DeleteCalendarGuiderByID())
	v1.PUT("/calendar_guiders", middlewares.RequiredAuth(), middlewares.RequiredRoles(constant.GuiderRole, constant.AdminRole), calendarGuiderHdl.UpdateCalendarGuiderByID())

	// booking guider
	v1.POST("/booking_guiders", bookingGuiderHdl.CreateBookingGuider())
	v1.GET("/confirm_booking_guider", bookingGuiderHdl.ConfirmStatusBookingGuider())
	v1.GET("/booking_guiders/:id", bookingGuiderHdl.GetBookingGuiderByID())
	v1.GET("/booking_guiders/user/:user_id", bookingGuiderHdl.GetBookingGuiderByUser())
	v1.POST("/booking_guiders/list", middlewares.RequiredAuth(), bookingGuiderHdl.ListBookingGuider())
	v1.DELETE("/booking_guiders/:id", middlewares.RequiredAuth(), bookingGuiderHdl.DeleteBookingGuiderByID())
	v1.PUT("/booking_guiders", middlewares.RequiredAuth(), bookingGuiderHdl.UpdateStatusBookingGuider())

	// request guider
	v1.POST("/request_guiders", middlewares.RequiredAuth(), requestGuiderHdl.UpsertRequestGuider())
	v1.GET("/request_guiders/list", requestGuiderHdl.ListRequestGuiderByUserID())
	v1.GET("/request_guiders/user/:user_id", requestGuiderHdl.GetRequestGuiderByUserID())
	v1.POST("/confirm_request_guider", requestGuiderHdl.ConfirmRequestGuider())

	// google login
	//v1.GET("/google/login")
	router.Run(":" + cfg.App.Port)
	wg.Wait()

}

func setupCors() cors.Config {
	configCORS := cors.DefaultConfig()
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"}
	configCORS.AllowCredentials = true
	//configCORS.AllowOrigins = []string{"http://localhost:3000"}
	configCORS.AllowAllOrigins = true

	return configCORS
}

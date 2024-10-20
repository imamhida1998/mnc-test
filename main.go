package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"mnc-test/model/request"
	"mnc-test/service/controller"
	"mnc-test/service/repository"
	"mnc-test/service/usecase"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	//dsn := `root:@tcp(127.0.0.1:3306)/testmnc?charset=utf8mb4&parseTime=True&loc=Local`
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open("mysql", dsn)

	db.LogMode(true)

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	topUpRepository := repository.NewTopUpRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	paymentRepository := repository.NewPaymentRepository(db)
	transferRepository := repository.NewTransferRepository(db)

	AuthJwt := usecase.NewAuthService()
	userUsecase := usecase.NewUserUsecase(userRepository)
	topupUsecase := usecase.NewTransactionUsecase(userRepository, topUpRepository, paymentRepository, transferRepository, transactionRepository)

	userController := controller.NewUserController(userUsecase, AuthJwt)
	topupController := controller.NewTopupController(topupUsecase)

	router := gin.Default()
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.POST("/topup", authMiddleware(AuthJwt, userUsecase), topupController.TopUp)
	router.POST("/pay", authMiddleware(AuthJwt, userUsecase), topupController.Payment)
	router.POST("/transfer", authMiddleware(AuthJwt, userUsecase), topupController.Transfer)
	router.GET("/transactions", authMiddleware(AuthJwt, userUsecase), topupController.Transaction)
	router.PUT("/profile", authMiddleware(AuthJwt, userUsecase), userController.UpdateProfile)

	server := http.Server{
		Addr:    ":7001",
		Handler: router,
	}
	server.ListenAndServe()
	if err != nil {
		return
	}

}

func authMiddleware(authService usecase.AuthService, userService usecase.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			//response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			response := gin.H{
				"status": "FAILED",
				"result": "Unauthorized",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := gin.H{
				"status": "FAILED",
				"result": err.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := gin.H{
				"status": "FAILED",
				"result": "Unauthorized",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["phone_number"].(string)
		pin := claim["pin"].(string)

		user, err := userService.GetAcoountByPhoneNumber(request.Login{
			userID,
			pin,
		})
		if err != nil {
			response := gin.H{
				"status": "FAILED",
				"result": err.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

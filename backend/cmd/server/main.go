package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	authHandlers "github.com/eskokado/startup-auth-go/backend/internal/handlers/auth"
	billingHandlers "github.com/eskokado/startup-auth-go/backend/internal/handlers/billing"
	orgHandlers "github.com/eskokado/startup-auth-go/backend/internal/handlers/org"
	taskHandlers "github.com/eskokado/startup-auth-go/backend/internal/handlers/task"
	"github.com/eskokado/startup-auth-go/backend/internal/middleware"
	"github.com/eskokado/startup-auth-go/backend/internal/providers"
	provider "github.com/eskokado/startup-auth-go/backend/internal/providers"
	repository "github.com/eskokado/startup-auth-go/backend/internal/repositories"
	authUsecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/auth"
	orgUsecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/org"
	taskUsecase "github.com/eskokado/startup-auth-go/backend/internal/usecase/task"
	service "github.com/eskokado/startup-auth-go/backend/pkg/domain/services"
)

func main() {
	// Carregar variáveis de ambiente
	_ = godotenv.Load(".env")
	sender := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		parsePort(os.Getenv("SMTP_PORT")),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	// 1. Configurar o banco de dados (SQLite para exemplo)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&repository.GormUser{}, &repository.GormOrganization{}, &repository.GormMembership{}, &repository.GormInvitation{}, &repository.GormTask{})

	// 2. Inicializar repositórios
	userRepo := repository.NewGormUserRepository(db)
	orgRepo := repository.NewGormOrganizationRepository(db)
	memberRepo := repository.NewGormMembershipRepository(db)
	invRepo := repository.NewGormInvitationRepository(db)
	taskRepo := repository.NewGormTaskRepository(db)

	// 3. Inicializar serviços
	emailService := service.NewEmailService(sender)

	// 4. Inicializar redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Endereço do Redis
		Password: "",               // Senha
		DB:       0,                // Banco padrão
	})

	// 4. Inicializar provedores
	cryptoProvider := provider.NewBcryptProvider(bcrypt.DefaultCost)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret"
	}
    tokenProvider := provider.NewJWTProvider(jwtSecret, 1*time.Hour)
	blacklistProvider := providers.NewRedisBlacklist(rdb)
	stripeProvider := providers.NewStripeProvider()

	// 5. Inicializar casos de uso
    personalOrgUC := orgUsecase.NewCreatePersonalOrgUsecase(orgRepo, memberRepo)
    registerUseCase := authUsecase.NewRegisterUsecase(userRepo, cryptoProvider, personalOrgUC)
    loggerUseCase := authUsecase.NewLoginUsecase(
        userRepo, cryptoProvider, tokenProvider, blacklistProvider, orgRepo,
    )
	logoutUseCase := authUsecase.NewLogoutUsecase(blacklistProvider)
	requestPasswordResetUC := authUsecase.NewRequestPasswordReset(userRepo, emailService)
	resetPasswordUC := authUsecase.NewResetPassword(userRepo)
	updateNameUC := authUsecase.NewUpdateNameUseCase(userRepo)
	updatePasswordUC := authUsecase.NewUpdatePasswordUseCase(userRepo, cryptoProvider)

	// Casos de uso de organização e tarefas
	inviteUC := orgUsecase.NewInviteMemberUsecase(invRepo, memberRepo, orgRepo, 72*time.Hour, os.Getenv("APP_BASE_URL"))
	acceptInviteUC := orgUsecase.NewAcceptInviteUsecase(invRepo, memberRepo)
	createTaskUC := taskUsecase.NewCreateTaskUsecase(taskRepo)
	listTasksUC := taskUsecase.NewListTasksUsecase(taskRepo)
	updateTaskUC := taskUsecase.NewUpdateTaskStatusUsecase(taskRepo)
	deleteTaskUC := taskUsecase.NewDeleteTaskUsecase(taskRepo)

	// 6. Criar handlers HTTP
	registerHTTPHandler := authHandlers.NewRegisterHandler(registerUseCase, userRepo)
	loggerHTTPHandler := authHandlers.NewLoginHandler(loggerUseCase)
	logoutHTTPHandler := authHandlers.NewLogoutHandler(logoutUseCase)
	forgotPasswordHandler := authHandlers.NewForgotPasswordHandler(requestPasswordResetUC)
	resetPasswordHandler := authHandlers.NewResetPasswordHandler(resetPasswordUC)
	updateNameHandler := authHandlers.NewUpdateNameHandler(updateNameUC)
	updatePasswordHandler := authHandlers.NewUpdatePasswordHandler(updatePasswordUC)

	inviteHandler := orgHandlers.NewInviteHandler(inviteUC)
	acceptInviteHandler := orgHandlers.NewAcceptInviteHandler(acceptInviteUC)
	getPersonalOrgHandler := orgHandlers.NewGetPersonalOrgHandler(orgRepo)
	createTaskHandler := taskHandlers.NewCreateTaskHandler(createTaskUC)
	listTasksHandler := taskHandlers.NewListTasksHandler(listTasksUC)
	updateTaskHandler := taskHandlers.NewUpdateTaskStatusHandler(updateTaskUC)
	deleteTaskHandler := taskHandlers.NewDeleteTaskHandler(deleteTaskUC)

	checkoutHandler := billingHandlers.NewCheckoutHandler(stripeProvider)

	// 7. Configurar roteador Gin
	router := gin.Default()

	// 7.1 Configurar CORS (ANTES dos middlewares de autenticação)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 7.2 Criar middleware de autenticação (DEPOIS do CORS)
	authMiddleware := middleware.JWTAuthMiddleware(tokenProvider, blacklistProvider)

	// 8. Registrar rotas
	router.POST("/auth/register", registerHTTPHandler.Handle)
	router.POST("/auth/login", loggerHTTPHandler.Handle)
	router.DELETE("/auth/logout", authMiddleware, logoutHTTPHandler.Handle)
	router.POST("/auth/forgot-password", forgotPasswordHandler.Handle)
	router.POST("/auth/reset-password", resetPasswordHandler.Handle)
	router.PUT("/user/name/:userID", authMiddleware, updateNameHandler.Handle)
	router.PUT("/user/password/:userID", authMiddleware, updatePasswordHandler.Handle)

	// Organização e tarefas
	router.GET("/org/personal/:ownerID", authMiddleware, getPersonalOrgHandler.Handle)
	router.POST("/org/invite", authMiddleware, inviteHandler.Handle)
	router.POST("/org/invite/accept", authMiddleware, acceptInviteHandler.Handle)
	router.POST("/tasks", authMiddleware, createTaskHandler.Handle)
	router.GET("/tasks", authMiddleware, listTasksHandler.Handle)
	router.PUT("/tasks/status", authMiddleware, updateTaskHandler.Handle)
	router.DELETE("/tasks/:id", authMiddleware, deleteTaskHandler.Handle)

	// Billing
	router.POST("/billing/checkout", authMiddleware, checkoutHandler.Handle)

	// 9. Iniciar o servidor
	router.Run(":8080")
}

func parsePort(port string) int {
	var p int
	fmt.Sscanf(port, "%d", &p)
	return p
}

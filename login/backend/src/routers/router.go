package router

import (
	"database/sql"
	"esystem/src/api/middleware"
	"fmt"
	"net/http"
	"os"

	"esystem/src/cmd/config"
	"esystem/src/pkg/auth"
	user "esystem/src/pkg/user"

	"esystem/src/pkg/utils/decode"
	utils "esystem/src/pkg/utils/encode"

	"esystem/src/pkg/utils/logger"

	"github.com/gorilla/mux"

	authHandler "esystem/src/api/auth/handler"

	userHandler "esystem/src/api/user/handler"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	defaultPort   = "8000"
	defaultAppEnv = "LOCAL"
)

var (
	authService *auth.Service
)

type RouterConfig struct {
	DB     *sql.DB
	Logs   *logger.ELK
	Config config.Configuration
}

func InitRouter(routerConfig RouterConfig) http.Handler {
	fmt.Println("InitRouter :", routerConfig)

	/* Repository */
	userRepository := user.NewRepository("TB_MAS_User", routerConfig.DB, routerConfig.Logs)
	authRepository := auth.NewRepository("TB_MAS_UserRole", routerConfig.DB, routerConfig.Logs)

	/* Service */
	utilService := utils.NewService(routerConfig.Logs)
	utilDecodeService := decode.NewService(routerConfig.Logs)

	userServiceParam := user.ServiceParam{
		UtilService: utilService,
		Repo:        userRepository,
		Logs:        routerConfig.Logs,
	}

	userService := user.NewService(userServiceParam)

	authServiceParam := auth.ServiceParam{
		UserService:       userService,
		UtilDecodeService: utilDecodeService,
		Repo:              authRepository,
		Logs:              routerConfig.Logs,
	}
	authService = auth.NewService(authServiceParam)

	/* Handler */

	userHandlerParam := userHandler.HandlerParam{
		Service: userService,
		Logs:    routerConfig.Logs,
	}
	userHandler := userHandler.MakeUserHandler(userHandlerParam)

	authHandlerParam := authHandler.HandlerParam{
		Service: authService,
		Logs:    routerConfig.Logs,
	}
	authHandler := authHandler.MakeAuthHandler(authHandlerParam)

	/* Router */
	middleware := middleware.NewMiddleware(routerConfig.Logs)
	router := mux.NewRouter()
	router.Use(middleware.Recover)
	router.HandleFunc("/version", versionHandler)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")

	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("", http.HandlerFunc(userHandler.Add)).Methods("POST")

	return router
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	appVersion := getEnvString("APP_VERSION", defaultPort)
	fmt.Fprintln(w, appVersion)
}

func getEnvString(env, fallback string) string {
	result := os.Getenv(env)
	if result == "" {
		return fallback
	}
	return result
}

type DBConnectURL struct {
	UserName string
	Password string
	DBHost   string
	Port     int
	DbName   string
}

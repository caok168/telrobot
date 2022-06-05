package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	controller "telrobot/controller"
	"telrobot/middlewares/jwt"
	logmw "telrobot/middlewares/log"
	repo "telrobot/repositories"
	app "telrobot/util/common"
)

func main(){
	fmt.Println("Hello World!")

	var (
		httpPort = app.EnvInt("OC_HTTP_PORT", 4000)
		dbURL = app.Env("OC_DB_URL", "host=localhost port=5432 user=robot password=robot dbname=robot sslmode=disable")
		cacheAddr = app.Env("REDIS_ADDR", "localhost:6379")
	)

	fmt.Println("hostPort:", httpPort, ", cacheAddr:", cacheAddr)

	gin.SetMode(gin.DebugMode)

	db, err := app.ConnectionPostgres(dbURL, false)
	if err != nil {
		fmt.Printf("db init failed: %v\n", err)
		return
	}

	defer db.Close()


	users := repo.NewUserRepository(db)

	svc := &RobotService{
		HTTPPort: httpPort,
		Users:    users,
	}

	errs := make(chan error, 2)

	go svc.startHTTPServer(errs)

	errinfo := <- errs
	fmt.Printf("service failed: %v\n", errinfo)
}

type RobotService struct {
	HTTPPort int
	Users    *repo.UserRepository
}

func (svc *RobotService) startHTTPServer(errs chan error) {
	p := fmt.Sprintf(":%d", svc.HTTPPort)
	errs <- http.ListenAndServe(p, svc.httpHandlers())
}

func (svc *RobotService) httpHandlers() *gin.Engine {
	r := gin.New()

	r.Use(logmw.Logger())
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{
		"GET", "POST", "PUT",
		"HEAD", "OPTIONS", "PATCH",
	}
	config.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type",
		"Tus-Resumable", "Upload-Length",
		"Upload-Metadata", "Upload-Offset",
	}
	r.Use(cors.New(config))

	svc.SetUserHandlers(r)

	return r
}

func (svc *RobotService) SetUserHandlers(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(jwt.JWT())
	{
		u := controller.NewUserController(svc.Users)
		{
			api.POST("/users", u.Create)
			api.GET("/users", u.List)
			api.GET("/users/:id", u.Get)
			api.PUT("/users/:id", u.Update)
			api.DELETE("/users/:id", u.Remove)
		}
	}
}

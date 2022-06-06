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
		token = app.Env("TOKEN", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjUxOWVmMmI0NTIwZDJhMzU3ZDNmZTdkZjE1Y2RmOGQ3NDJmOGNjNGViODVlNWI1ZWRmYjcyYjYyODg2NDU2NmYwOWFiZjM2OGRkNjViYjYwIn0.eyJhdWQiOiIxIiwianRpIjoiNTE5ZWYyYjQ1MjBkMmEzNTdkM2ZlN2RmMTVjZGY4ZDc0MmY4Y2M0ZWI4NWU1YjVlZGZiNzJiNjI4ODY0NTY2ZjA5YWJmMzY4ZGQ2NWJiNjAiLCJpYXQiOjE2NTM3MDE2ODQsIm5iZiI6MTY1MzcwMTY4NCwiZXhwIjoxNjg1MjM3Njg0LCJzdWIiOiJiNzM0NTYyMy0xMGRjLTRjNjktOGFhNy1jODdmMThmODljMzciLCJzY29wZXMiOltdfQ.cLqXS2KEOEvfdYOGFMeCqtphHb_JLubwjI1oKYxFgS0bIeMRolTSSWi1BE0HWO5DRmGAxrDOBS4iGqWtsb1FdUhbQnLYIe1aa9jOwGHm6kR_GJHHF3ETai8UGWqaTjnuiC74IsexyIlgr3Qj7w4kDuJuO60SnYcJTrxsIRRh736P71tux61PZrokM9UjAFqrD9Pn7kIHcsCuMGdExQm8MmFkDGC9VIMgfWrdtEAyZd-eEhTpqLlyA-ov_LCQeFiZHhfm18zWH_KFmvYj1ftqv_o8mQ5GcF8Q_43umiPZXH70rEaQutp79KNmt0g08_-Y3TwG2PY57P9WB0_boeRAeycoCWwUk_F0CmulM0eyBs-OIyI6HIxXQOQ2O53OYrPglGhoBklVa6Y7_ie3n94xu_I2x1QiDk6Uf_0NqH2V8CIUu_i1v-e-vQmYWAL_K9TGDOPBiq_J2soyuA7qHfd1kswF2O2nYQdJ6TlS2Fe-WYtN1PiJ3M9UPgKGm8JHekib-XONmhBCl8cdEhV0SXNPFhXx49CgD65CT8tVDB5VRYPWqPLztoRdkZ5LmrTv2XbpNSMBwT4CwXhaeLMPKxtl2HHTHmvPYMCZGaib3TAPdkXkpHKLAMNb91XMu9u8lz1Fq5PRlx5GssCdU_GpOCgww7gyWatFgvzArvaJf_dg60k")
	)

	fmt.Println("hostPort:", httpPort, ", cacheAddr:", cacheAddr)
	fmt.Println("token:", token)

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

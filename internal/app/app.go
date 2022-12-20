package app

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yudgxe/sima-rest-api/internal/delivery/rest/handler"
	"github.com/yudgxe/sima-rest-api/internal/service/basic"
	"github.com/yudgxe/sima-rest-api/internal/store/psql"
	"github.com/yudgxe/sima-rest-api/pkg/database"

	_ "github.com/lib/pq"
)

var (
	singingkey = "as4KWkw3491k3jdIIWl21"
	tokenTTL   = 2 * time.Hour
)

func Start(c *Config) {
	e := echo.New()

	db, err := database.NewPostgres(database.PostgresConnInfo{
		Host:     c.DB.Host,
		Port:     c.DB.Port,
		User:     c.DB.User,
		Password: c.DB.Password,
		Name:     c.DB.Name,
		SSLMode:  database.ModeDisable,
	})
	if err != nil {
		e.Logger.Fatal(err)
	}

	store := psql.New(db)

	handler.NewHandler(&handler.Deps{
		UserService: basic.NewUserService(store),
		AuthService: basic.NewAuthService(store, singingkey, tokenTTL),
	}).Init(e, singingkey)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Logger.Fatal(e.Start(c.BindAddr))

	/*
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		a := <-done

		fmt.Println(a)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	*/
}

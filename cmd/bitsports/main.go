package main

import (
	"bitsports"
	"bitsports/driver-framework/middlewares"
	"bitsports/ent"
	"bitsports/ent/migrate"
	"bitsports/usecase/dataprovider"
	"bitsports/usecase/jwt"
	passwordvalidator "bitsports/usecase/password_validator"
	"bitsports/utils"
	"context"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error
	var client *ent.Client

	//simulation for a non prodcution stage, her we also could load another data source information
	// using dataprovider.SetDataSource(setDatasource with env values)
	if os.Getenv("PRODUCTION") != "" {
		client, err = dataprovider.GetPostgresClient()
	} else {
		client, err = dataprovider.GetSqliteClient()
	}

	if err != nil {
		log.Fatal("opening ent client", err)
	}
	e := echo.New()

	// e.Logger = logrus.New().Formatter
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	pv := passwordvalidator.NewPasswordValidator(
		passwordvalidator.WithMinLen(6),
	)
	srv := handler.NewDefaultServer(bitsports.NewSchema(client, pv))

	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		operation := graphql.GetPath(ctx).String()
		ctx, err = jwt.ValidateOperation(ctx, client, operation)
		if err != nil {
			println(err.Error())
			graphql.AddError(ctx, err)
			return nil, nil
		}
		return next(ctx)

	})

	playgroundHandler := playground.Handler("Bitsports", "/query")
	e.GET("/", middlewares.EchoWrapper(playgroundHandler))

	e.POST("/query", middlewares.EchoWrapper(srv.ServeHTTP),
		echo.WrapMiddleware(middlewares.InjectToContext),
		echo.WrapMiddleware(jwt.SetJwtTokenMiddleware),
	)

	if err := e.Start(":" + utils.GetEnvOrDefault("WEB_PORT", "8081")); err != nil {
		log.Fatal("http server terminated", err)
	}
}

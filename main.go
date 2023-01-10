// Copyright (c) 2023 ...

package main

import (
	"golang-rest/controllers"
	"golang-rest/globals"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)

	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		var cfg aws.Config = aws.Config{
			Region:   aws.String(os.Getenv("REGION")),
			Endpoint: aws.String(os.Getenv("DYNAMO_ENDPOINT")),
			LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody),
		}
		db = dynamodb.New(sess, &cfg)
	}

	globals.TableName = aws.String(os.Getenv("TABLE"))
	globals.DynamoDbSvc = new(globals.DBService)
	globals.DynamoDbSvc.Session = sess
	globals.DynamoDbSvc.DbSvc = dynamodbiface.DynamoDBAPI(db)
}

func options(c *gin.Context) {
	c.Status(http.StatusNoContent)
	return
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	// https://github.com/gin-contrib/cors
	// replace * with some real value in the future for usage behind google ApiGateway
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, OPTIONS, GET, PUT, DELETE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour}))

	api := router.Group("/api")

	// Ping to check health
	api.GET("/healthcheck", controllers.HealthCheck)

	v1 := api.Group("/v1")
	{
		// Entity
		v1.GET("/entity/:entityId", controllers.GetEntity)

	}
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}

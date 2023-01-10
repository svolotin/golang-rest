package globals

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var TableName *string

var DynamoDbSvc *DBService

type DBService struct {
	DbSvc dynamodbiface.DynamoDBAPI
	Session *session.Session
}
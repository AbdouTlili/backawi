package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DBManager struct {
	AwsSession          *session.Session
	DynamoServiceClient *dynamodb.DynamoDB
	DynamodbTable       string
}

type ProductItem struct {
	ProductID string
	Name      string
	Quantity  int
	Discount  float64
	Price     float64
}

// type Item struct {
// 	Product
// }

// item := map[string]*dynamodb.AttributeValue{
// 	"ProductID": {
// 		S: aws.String("xxx"),
// 	},
// 	"Name": {
// 		S: aws.String("Sample Product"),
// 	},
// 	"Quantity": {
// 		N: aws.String("10"),
// 	},
// 	"Discount": {
// 		S: aws.String("100%"),
// 	},
// 	"Price": {
// 		N: aws.String("50.00"),
// 	},
// }

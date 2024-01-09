package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	log "github.com/sirupsen/logrus"
)

type DBManager struct {
	AwsSession          *session.Session
	DynamoServiceClient *dynamodb.DynamoDB
}

func (dbm *DBManager) Init() {
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("eu-west-3"), // Specify your AWS region
	// })

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-3"),
		Credentials: credentials.NewStaticCredentials("AKIARFMYROC2IDSBYY7B", "lgfcZmI855cn1z93lNk/dWd9t9XAhkH6eXEKSsdW", ""),
	})

	if err != nil {
		log.Fatalln("Error creating session:", err)
		return
	}

	dbm.AwsSession = sess

	// Create a new DynamoDB service client.
	dbm.DynamoServiceClient = dynamodb.New(sess)
}

func main() {
	// Create a new session using environment variables and shared credentials.
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("eu-west-3"), // Specify your AWS region
	// })

	// if err != nil {
	// 	fmt.Println("Error creating session:", err)
	// 	return
	// }

	dbm := DBManager{}
	dbm.Init()

	// Create a new DynamoDB service client.
	// svc := dynamodb.New(sess)

	// Define the item you want to add to the DynamoDB table.
	item := map[string]*dynamodb.AttributeValue{
		"ProductID": {
			S: aws.String("xxx"),
		},
		"Name": {
			S: aws.String("Sample Product"),
		},
		"Quantity": {
			N: aws.String("10"),
		},
		"Discount": {
			S: aws.String("100%"),
		},
		"Price": {
			N: aws.String("50.00"),
		},
	}

	// Specify the name of your DynamoDB table.
	tableName := "products_table"

	// Build the input for the PutItem operation.
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	// Add the item to the DynamoDB table.
	_, err := dbm.DynamoServiceClient.PutItem(input)

	if err != nil {
		fmt.Println("Error adding item to DynamoDB:", err)
		return
	}

	fmt.Println("Item added to DynamoDB successfully.")
}

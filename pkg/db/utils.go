package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateItem(productID string,
	productName string,
	quantity string,
	discount string,
	price string) map[string]*dynamodb.AttributeValue {

	item := map[string]*dynamodb.AttributeValue{
		"ProductID": {
			S: aws.String(productID),
		},
		"Name": {
			S: aws.String(productName),
		},
		"Quantity": {
			N: aws.String(quantity),
		},
		"Discount": {
			S: aws.String(discount),
		},
		"Price": {
			N: aws.String(price),
		},
	}
	return item
}

// input := &dynamodb.PutItemInput{
// 	Item:      item,
// 	TableName: aws.String(tableName),
// }

func CreateItemPutPayload(item map[string]*dynamodb.AttributeValue, tableName string) *dynamodb.PutItemInput {

	return &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

}

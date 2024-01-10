package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (dbm *DBManager) Init(cfg *aws.Config, tableName string) {

	dbm.DynamodbTable = tableName

	var err error

	dbm.AwsSession, err = session.NewSession(cfg)

	if err != nil {
		log.Fatalln("Error creating session:", err)
		return
	}

	// Create a new DynamoDB service client.
	dbm.DynamoServiceClient = dynamodb.New(dbm.AwsSession)
}

func (dbm *DBManager) CreateItemInDB(productID string,
	productName string,
	quantity string,
	discount string,
	price string) error {
	item := CreateItem(productID, productName, quantity, discount, price)
	payload := CreateItemPutPayload(item, dbm.DynamodbTable)

	_, err := dbm.DynamoServiceClient.PutItem(payload)

	if err != nil {
		log.Warn("Failed to add element to DB ", err)
	} else {
		log.Info("Item added to DynamoDB successfully.", err)
	}
	return err
}

func (dbm *DBManager) GetAllItemsInDB() error {
	input := &dynamodb.ScanInput{
		TableName: aws.String(dbm.DynamodbTable),
	}

	// perform the scan to get all the elements
	result, err := dbm.DynamoServiceClient.Scan(input)
	if err != nil {
		log.Warn("Error scanning Dynamodb table ", err)
		return err
	}

	log.Info("Received ", *result.Count, " items from DynamoDB")
	for _, item := range result.Items {
		productItem := CastDbRawItemToProductObject(item)
		log.Info(productItem)
	}
	return nil
}

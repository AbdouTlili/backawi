package db

import (
	"strconv"

	log "github.com/sirupsen/logrus"

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
			N: aws.String(discount),
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

func CastDbRawItemToProductObject(item map[string]*dynamodb.AttributeValue) ProductItem {
	productItem := ProductItem{}

	exists := false
	var err error

	tmpPrice, exists := item["Price"]
	if exists {
		productItem.Price, err = strconv.ParseFloat(*tmpPrice.N, 64)
		if err != nil {
			log.Fatal("error parsing the Price received from dynamoDB", err)
		}
	} else {
		log.Fatal("Key does not exist in the Item object")
	}

	tmpQuantity, exists := item["Quantity"]
	if exists {
		productItem.Quantity, err = strconv.Atoi(*tmpQuantity.N)
		if err != nil {
			log.Fatal("error parsing the Quantity received from dynamoDB", err)
			return ProductItem{}
		}
	} else {
		log.Fatal("Key Quantity does not exist in the Item object")
		return ProductItem{}
	}

	tmpName, exists := item["Name"]
	if exists {
		productItem.Name = *tmpName.S
		if err != nil {
			log.Fatal("error parsing the Name received from dynamoDB", err)
			return ProductItem{}
		}
	} else {
		log.Fatal("Key Name does not exist in the Item object")
		return ProductItem{}
	}

	tmpDiscount, exists := item["Discount"]
	if exists {
		productItem.Discount, err = strconv.ParseFloat(*tmpDiscount.N, 64)
		if err != nil {
			log.Fatal("error parsing the Discount received from dynamoDB", err)
			return ProductItem{}
		}
	} else {
		log.Fatal("Key Discount does not exist in the Item object")
		return ProductItem{}
	}

	tmpProductID, exists := item["ProductID"]
	if exists {
		productItem.ProductID = *tmpProductID.S
		if err != nil {
			log.Fatal("error parsing the ProductID received from dynamoDB", err)
			return ProductItem{}
		}
	} else {
		log.Fatal("Key ProductID does not exist in the Item object")
		return ProductItem{}
	}

	// log.Info("Product Price received from db is : ", productItem.Price)
	return productItem
}

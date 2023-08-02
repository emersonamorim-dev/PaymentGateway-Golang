package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func SavePayment(payment interface{}) error {
	av, err := dynamodbattribute.MarshalMap(payment)
	if err != nil {
		return err
	}

	// Cria uma nova sessão do AWS SDK
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return err
	}

	db := dynamodb.New(sess)

	// Cria um novo item no DynamoDB
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Payments"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

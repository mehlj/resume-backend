package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  
  "github.com/aws/aws-lambda-go/lambda"

  "log"
)

// Increments atomic counter in DynamoDB
func incrementCounter() (string, error) {
  session := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))
  svc := dynamodb.New(session)

  tableName := "resume-counter"

  input := &dynamodb.UpdateItemInput {
    TableName: aws.String(tableName),
    Key: map[string]*dynamodb.AttributeValue{
      "primaryKey": {
          S: aws.String("VisitorCounter"),
      },
    },
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
      ":num": {
          N: aws.String("1"),
      },
    },
    ReturnValues:     aws.String("UPDATED_NEW"),
    UpdateExpression: aws.String("set counterValue = counterValue + :num"),
  }

  _, err := svc.UpdateItem(input)
  if err != nil {
      log.Fatalf("Got error calling UpdateItem: %s", err)
  }

  return "Incremented visitor counter", nil
}

func main() {
  lambda.Start(incrementCounter)
}
package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/events"

  "fmt"
  "log"
  "net/http"
  "strconv"
)

type Item struct {
  CounterValue int `json:"counterValue"`
  PrimaryKey   string `json:"primaryKey"`
}

// get current counter value
func getCounter() (int){
  session := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))
  svc := dynamodb.New(session)

  tableName := "resume-counter"

  result, err := svc.GetItem(&dynamodb.GetItemInput{
    TableName: aws.String(tableName),
    Key: map[string]*dynamodb.AttributeValue{
      "primaryKey": {
        S: aws.String("VisitorCounter"),
      },
    },
  })

  if err != nil {
    log.Fatalf("Got error calling GetItem: %s", err)
  }

  item := Item{}

  err = dynamodbattribute.UnmarshalMap(result.Item, &item)
  if err != nil {
    panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
  }
  
  return item.CounterValue
}

// Increments atomic counter in DynamoDB
func incrementCounter(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusOK,
    Body:       strconv.Itoa(getCounter()),
  }, nil
}

func main() {
  lambda.Start(incrementCounter)
}
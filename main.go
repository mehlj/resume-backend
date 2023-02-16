import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-lambda-go/lambda"

  "fmt"
  "log"
)


// custom types
type Request struct {
	Id int `json:"Id"`
}
  
type Response struct {
	Message string `json:"message"`
}


// handler function
func requestHandler(req Request) (Response, error) {

  session := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))

  // Create DynamoDB client
  svc := dynamodb.New(session)

  // PICK UP HERE
  result, err := dynamodbattribute.MarshalMap(req)
  if err != nil {
    fmt.Println("Failed to marshall request")
    return Response{}, err
  }

  input := &dynamodb.PutItemInput{
    Item:      result,
    TableName: aws.String("TestTable"),
  }

  _, err = svc.PutItem(input)
  if err != nil {
    fmt.Println("Failed to write to db")
    return Response{}, err
  }

  return Response{fmt.Sprintf("Request processed %d", req.Id)}, nil
}


// start Lambda routine and pass our handler function to it
func main() {
  lambda.Start(requestHandler)
}
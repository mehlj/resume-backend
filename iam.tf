resource "aws_iam_role" "lambda-role" {
 name = "lambda-role"

 assume_role_policy = jsonencode({
   "Version" : "2012-10-17",
   "Statement" : [
     {
       "Effect" : "Allow",
       "Principal" : {
         "Service" : "lambda.amazonaws.com"
       },
       "Action" : "sts:AssumeRole"
     }
   ]
  })
}
          
resource "aws_iam_policy" "dynamodb-lambda-policy" {
   name   = "dynamodb-lambda-policy"

   policy = jsonencode({
      "Version" : "2012-10-17",
      "Statement" : [
        {
           "Effect" : "Allow",
           "Action" : ["dynamodb:*"],
           "Resource" : aws_dynamodb_table.resume-counter.arn
        }
      ]
   })
}

resource "aws_iam_role_policy_attachment" "lambda-policy-attach" {
   role       = aws_iam_role.lambda-role.name
   policy_arn = aws_iam_policy.dynamodb-lambda-policy.arn
}
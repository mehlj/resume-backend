// TODO - complete (besides role line)

# resource "aws_lambda_function" "function" {
#   function_name = "MyLambdaFunction"

#   role             = aws_iam_role.lambda-role.arn
#   handler          = "index.handler"
#   filename         = "${path.module}/assets/index.js.zip"

#   runtime     = "nodejs16.x"
#   memory_size = 128
#   timeout     = 3
# }
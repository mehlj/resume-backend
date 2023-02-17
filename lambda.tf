resource "aws_lambda_function" "function" {
  filename         = "lambda-handler.zip"
  function_name    = "resume-visitor-counter"
  source_code_hash = filebase64sha256("lambda-handler.zip")
  role             = aws_iam_role.lambda-role.arn

  handler = "main"
  runtime = "go1.x"
}
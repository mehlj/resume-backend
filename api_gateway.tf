resource "aws_apigatewayv2_api" "lambda" {
  name          = "resume_lambda_gw"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "lambda" {
  api_id = aws_apigatewayv2_api.lambda.id
  name        = "resume_lambda_stage"

  auto_deploy = true
}

resource "aws_apigatewayv2_integration" "lambda" {
  api_id             = aws_apigatewayv2_api.lambda.id
  integration_type   = "AWS_PROXY"

  integration_uri    = aws_lambda_function.function.invoke_arn
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "increment" {
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "GET /increment"
  target    = "integrations/${aws_apigatewayv2_integration.lambda.id}"
}

resource "aws_lambda_permission" "api_gw" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.function.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}
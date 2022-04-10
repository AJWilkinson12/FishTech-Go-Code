data "aws_region" "current" {}

resource "aws_lambda_function" "fishTech" {
  function_name    = "fishTech"
  filename         = "fishTech.zip"
  handler          = "fishTech"
  source_code_hash = sha256(filebase64("fishTech.zip"))
  role             = aws_iam_role.fishTech.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 1
}

resource "aws_iam_role" "fishTech" {
  name               = "fishTech"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY
}

resource "aws_lambda_permission" "fishTech" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.fishTech.arn
  principal     = "apigateway.amazonaws.com"
}

resource "aws_api_gateway_resource" "fishTech" {
  rest_api_id = aws_api_gateway_rest_api.fishTech.id
  parent_id   = aws_api_gateway_rest_api.fishTech.root_resource_id
  path_part   = "fishTech"
}

resource "aws_api_gateway_rest_api" "fishTech" {
  name = "fishTech"
}

resource "aws_api_gateway_method" "fishTech" {
  rest_api_id   = aws_api_gateway_rest_api.fishTech.id
  resource_id   = aws_api_gateway_resource.fishTech.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "fishTech" {
  rest_api_id             = aws_api_gateway_rest_api.fishTech.id
  resource_id             = aws_api_gateway_resource.fishTech.id
  http_method             = aws_api_gateway_method.fishTech.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${aws_lambda_function.fishTech.arn}/invocations"
}

resource "aws_api_gateway_deployment" "fishTech_v1" {
  depends_on = [
    "aws_api_gateway_integration.fishTech"
  ]
  rest_api_id = aws_api_gateway_rest_api.fishTech.id
  stage_name  = "v1"
}

output "url" {
  value = "${aws_api_gateway_deployment.fishTech_v1.invoke_url}${aws_api_gateway_resource.fishTech.path}?ip=54.192.0.0"
}

resource "aws_api_gateway_rest_api" "lambda_api" {
  name               = "lambda_api"
  description        = "Lamba function to generate gmail-like profile images"
  binary_media_types = ["*/*"]
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = "${aws_api_gateway_rest_api.lambda_api.id}"
  parent_id   = "${aws_api_gateway_rest_api.lambda_api.root_resource_id}"
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = "${aws_api_gateway_rest_api.lambda_api.id}"
  resource_id   = "${aws_api_gateway_resource.proxy.id}"
  http_method   = "GET"
  authorization = "NONE"

  request_parameters = {
    "method.request.querystring.initials" = true
    "method.request.querystring.size"     = true
  }
}

resource "aws_api_gateway_integration" "lambda" {
  rest_api_id = "${aws_api_gateway_rest_api.lambda_api.id}"
  resource_id = "${aws_api_gateway_method.proxy.resource_id}"
  http_method = "${aws_api_gateway_method.proxy.http_method}"

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "${aws_lambda_function.lambda_api.invoke_arn}"

  request_parameters = {
    "integration.request.querystring.initials" = "'method.request.querystring.initials'"
    "integration.request.querystring.size"     = "'method.request.querystring.size'"
  }
}

resource "aws_api_gateway_method" "proxy_root" {
  rest_api_id   = "${aws_api_gateway_rest_api.lambda_api.id}"
  resource_id   = "${aws_api_gateway_rest_api.lambda_api.root_resource_id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda_root" {
  rest_api_id = "${aws_api_gateway_rest_api.lambda_api.id}"
  resource_id = "${aws_api_gateway_method.proxy_root.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_root.http_method}"

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "${aws_lambda_function.lambda_api.invoke_arn}"
}

resource "aws_api_gateway_deployment" "lambda_api" {
  depends_on = [
    "aws_api_gateway_integration.lambda",
    "aws_api_gateway_integration.lambda_root",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.lambda_api.id}"
  stage_name  = "test"
}

output "base_url" {
  value = "${aws_api_gateway_deployment.lambda_api.invoke_url}"
}

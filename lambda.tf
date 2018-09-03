provider "aws" {
  region = "us-west-2"
}

resource "aws_lambda_function" "lambda" {
  function_name    = "tf-default-profile-image"
  handler          = "default-profile-image-lambda"
  runtime          = "go1.x"
  filename         = "default-profile-image-lambda.zip"
  role             = "${aws_iam_role.lambda_exec.arn}"
  source_code_hash = "${base64sha256(file("default-profile-image-lambda.zip"))}"
}

resource "aws_iam_role" "lambda_exec" {
  name = "tf-default-profile-image"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.lambda.function_name}"
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_deployment.lambda.execution_arn}/*/*"
}

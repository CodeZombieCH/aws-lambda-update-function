# AWS Lambda Update Function

A simple tool to update an AWS Lambda function by uploading a ZIP archive.

## Usage

    update --function-name <function-name> --zip-file <zip-file>

with:
- `<function-name>`: Name of the AWS Lambda function to update (must exist)
- `<zip-file>`: Path to ZIP file to upload

Please make sure to configure AWS credentials, e.g. by setting the following environment variables:
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_REGION

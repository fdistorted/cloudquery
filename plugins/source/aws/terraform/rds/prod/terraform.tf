terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "rds"
    region = "us-east-1"
  }
}

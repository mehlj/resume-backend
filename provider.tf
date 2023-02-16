terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }

  backend "s3" {
    bucket         = "resume-backend-tfstate"
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "resumebackend_state_locks"
    encrypt        = true
  }
}

provider "aws" {
  alias  = "virginia"
  region = "us-east-1"
}
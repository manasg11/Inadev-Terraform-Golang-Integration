provider "aws" {
  region = "us-east-2"  
}

resource "aws_s3_bucket" "my_bucket" {
  bucket = "manasbucket7017safari"  
}

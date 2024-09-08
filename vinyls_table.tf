resource "aws_dynamodb_table" "vinyls-table" {
  name           = "vinyls"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "vinyl_id"

  attribute {
    name = "vinyl_id"
    type = "S"
  }

  tags = {
    Name        = "vinyls-table"
  }
}
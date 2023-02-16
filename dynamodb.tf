resource "aws_dynamodb_table_item" "init-counter" {
  table_name = aws_dynamodb_table.resume-counter.name
  hash_key   = aws_dynamodb_table.resume-counter.hash_key

  item = <<ITEM
{
  "VisitorCount": {"N": "0"}
}
ITEM
}

resource "aws_dynamodb_table" "resume-counter" {
  name         = "resume-counter"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "VisitorCount"

  attribute {
    name = "VisitorCount"
    type = "N"
  }
}
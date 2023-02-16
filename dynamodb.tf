resource "aws_dynamodb_table_item" "visitor-counter" {

  # ignore changes to counterValue - this will change over time
  lifecycle {
    ignore_changes = [
      item,
    ]
  }

  table_name = aws_dynamodb_table.resume-counter.name
  hash_key   = aws_dynamodb_table.resume-counter.hash_key

  item = <<ITEM
{
  "primaryKey": {"S": "VisitorCounter"},
  "counterValue": {"N": "0"}
}
ITEM
}

resource "aws_dynamodb_table" "resume-counter" {
  name         = "resume-counter"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "primaryKey"

  attribute {
    name = "primaryKey"
    type = "S"
  }
}
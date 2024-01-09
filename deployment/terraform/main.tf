resource "aws_dynamodb_table" "products_table" {
    name = "products_table"
    billing_mode = "PAY_PER_REQUEST"
    hash_key = "ProductID"

    attribute {
      name = "ProductID"
      type = "S"
    }

    tags = {
      "Name" = "dynamo-db-product-table"
      "Environment" = "test"
    }
    
  
}

output "dynamodb-arn" {
  value = aws_dynamodb_table.products_table.arn
}

output "dynamodb-table-id" {
  value = aws_dynamodb_table.products_table.id
}
package generator

var terraformProvider = `
provider "aws" {}
`

var terraformDynamoDB = `

resource "aws_dynamodb_table" "{{ .Cfg.TableName }}" {
  name           = "{{ .Cfg.TableName }}"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "{{ .Cfg.PrimaryIndex.PartitionKey.AttrName }}"
  {{- if .Cfg.PrimaryIndex.SortKey }}
  range_key      = "{{ .Cfg.PrimaryIndex.SortKey.AttrName }}"
  {{- end }}
  {{- range $name, $val := .Attrs }}

  attribute {
    name = "{{$name}}"
    type = "{{$val}}"
  }
  {{- end }}

  {{- range .Cfg.GlobalSecondaryIndexes }}

  global_secondary_index {
    name               = "{{ .IndexName }}"
    hash_key           = "{{ .PartitionKey.AttrName }}"
	range_key          = "{{ .SortKey.AttrName}}"
	projection_type = "ALL"
  }
  {{ end -}}

}

`

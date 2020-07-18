package templates

var DynamoDBUpdateMessage = `
func (d *DB) DDBUpdate{{ .MessageName }}(ctx context.Context, item *{{ .MessageName }}) (*{{ .MessageName }}, error) {

	_, err := d.DDBGet{{ .MessageName }}(ctx, item.Name)

	if err != nil {
		return item, err
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return item, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName{{ .MessageName }}),
	}

	_, err = d.Client.PutItemWithContext(ctx, input)

	if err != nil {
		return item, err
	}

	return item, nil
}
`

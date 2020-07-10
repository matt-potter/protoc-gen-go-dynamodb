package templates

// DynamoDBCreateFunc is the template for the create func
var DynamoDBCreateFunc string = `
func (d *db) Create(ctx context.Context, item *api.Activity) (*api.Activity, error) {

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return item, err
	}

	expr, err := expression.NewBuilder().WithCondition(expression.AttributeNotExists(expression.Name("id"))).Build()

	if err != nil {
		return item, err
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(TableNameActivity),
		ConditionExpression: expr.Condition(),
	}

	_, err = d.Client.PutItemWithContext(ctx, input)

	if err != nil {
		return item, err
	}

	return item, nil

}
`

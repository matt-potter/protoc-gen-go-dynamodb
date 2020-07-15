package templates

var DynamoDBHeaderAndClient string = `
package {{ .PackageName }}

import "github.com/aws/aws-sdk-go/service/dynamodb"

type DB struct {
	Client *dynamodb.DynamoDB
}

func NewDB(ddb *dynamodb.DynamoDB) *DB {
	return &DB{
		Client: ddb,
	}
}
`

var DynamoDBMessageHeaders string = `
package {{ .PackageName }}

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/pkg/errors"
)

const TableName{{ .MessageName }} = "{{ .TableName }}"
`

var DynamoDBGetMessageMulti = `
func (d *DB) DDBGet{{ .MessageName | Plural }}(ctx context.Context, prefix string, startKey string, limit int64) ({{ .MessageName | Plural | ToLower }} []*{{ .MessageName }}, lastEvaluatedKey string, err error) {

	type lastKeyResponse map[string]struct {
	}

	expr, err := expression.NewBuilder().WithKeyCondition(expression.KeyBeginsWith(expression.Key("collection"), prefix)).Build()

	if err != nil {
		return nil, "", err
	}

	input := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Limit:                     aws.Int64(limit),
		TableName:                 aws.String(TableName{{ .MessageName }}),
		IndexName:                 aws.String("collection-id-index"),
	}

	if startKey != "" {

		decoded, err := base64.StdEncoding.DecodeString(startKey)

		if err != nil {
			return nil, "", err
		}

		log.Println("Decoded", string(decoded))

		exclusiveStartKey, err := dynamodbattribute.MarshalMap(string(decoded))

		if err != nil {
			return nil, "", err
		}

		input.ExclusiveStartKey = exclusiveStartKey

		log.Println("EX:", exclusiveStartKey)
	}

	res, err := d.Client.QueryWithContext(ctx, input)

	if err != nil {
		return nil, "", err
	}

	lastEvaluatedKey = ""

	if len(res.LastEvaluatedKey) > 0 {

		log.Println(res.LastEvaluatedKey)

		pageResponseMap := make(map[string]interface{})

		for _, val := range res.LastEvaluatedKey {
			log.Println(val.GoString())
		}

		err = dynamodbattribute.UnmarshalMap(res.LastEvaluatedKey, &pageResponseMap)

		if err != nil {
			return nil, "", err
		}

		b, err := json.Marshal(pageResponseMap)

		if err != nil {
			return nil, "", err
		}

		lastEvaluatedKey = base64.StdEncoding.EncodeToString(b)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &{{ .MessageName | Plural | ToLower }})

	if err != nil {
		return nil, "", err
	}

	log.Println("LAST KEY", lastEvaluatedKey)

	return activities, lastEvaluatedKey, nil
}
`
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

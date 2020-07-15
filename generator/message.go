package generator

import (
	"log"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/matt-potter/protoc-gen-go-dynamodb/dynamopb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func (g *Generator) generateMessageHeader(f *protogen.GeneratedFile, msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	f.P(`package `, g.packageName)
	f.P(`import (`)
	f.P(`"context"`)
	f.P(`"encoding/base64"`)
	f.P(`"encoding/json"`)
	f.P(`"log"`)
	f.P(`"github.com/aws/aws-sdk-go/aws"`)
	f.P(`"github.com/aws/aws-sdk-go/service/dynamodb"`)
	f.P(`"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"`)
	f.P(`"github.com/aws/aws-sdk-go/service/dynamodb/expression"`)
	f.P(`"github.com/pkg/errors"`)
	f.P(`)`)
	f.P(`const TableName`, msg.GoIdent.GoName, ` = "`, cfg.GetTableName(), `"`)
	f.P()
}

func (g *Generator) generateMessageCreate(f *protogen.GeneratedFile, msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	f.P(`func (d *DB) DDBCreate`, msg.GoIdent.GoName, `(ctx context.Context, item *`, msg.GoIdent.GoName, `) (*`, msg.GoIdent.GoName, `, error) {`)
	f.P(`av, err := dynamodbattribute.MarshalMap(item)`)
	f.P(`if err != nil {`)
	f.P(`return item, err`)
	f.P(`}`)

	switch cfg.PrimaryIndex.SortKey {
	case nil:
		f.P(`expr, err := expression.NewBuilder().WithCondition(expression.AttributeNotExists(expression.Name("`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `"))).Build()`)
	default:
		f.P(`expr, err := expression.NewBuilder().WithCondition(expression.AttributeNotExists(expression.Name("`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `")).And(expression.AttributeNotExists(expression.Name("`, cfg.PrimaryIndex.SortKey.GetAttrName(), `")))).Build()`)
	}

	f.P(`if err != nil {`)
	f.P(`return item, err`)
	f.P(`}`)
	f.P(`input := &dynamodb.PutItemInput{`)
	f.P(`Item:                      av,`)
	f.P(`TableName:                 aws.String(TableName`, msg.GoIdent.GoName, `),`)
	f.P(`ConditionExpression:       expr.Condition(),`)
	f.P(`ExpressionAttributeNames:  expr.Names(),`)
	f.P(`ExpressionAttributeValues: expr.Values(),`)
	f.P(`}`)
	f.P(`_, err = d.Client.PutItemWithContext(ctx, input)`)
	f.P(`if err != nil {`)
	f.P(`return item, err`)
	f.P(`}`)
	f.P(`return item, nil`)
	f.P(`}`)
	f.P()
}

func (g *Generator) generateMessageGet(f *protogen.GeneratedFile, msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	switch cfg.PrimaryIndex.SortKey {
	case nil:
		f.P(`func (d *DB) DDBGet`, msg.GoIdent.GoName, `By`, strings.Title(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` (ctx context.Context, `, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.PartitionKey.GetAttrType()], `) (*`, msg.GoIdent.GoName, `, error) {`)
	default:
		f.P(`func (d *DB) DDBGet`, msg.GoIdent.GoName, `By`, strings.Title(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `And`, strings.Title(cfg.PrimaryIndex.SortKey.GetAttrName()), ` (ctx context.Context, `, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.PartitionKey.GetAttrType()], `, `, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.SortKey.GetAttrType()], `) (*`, msg.GoIdent.GoName, `, error) {`)
	}

	f.P(`input := &dynamodb.GetItemInput{`)
	f.P(`Key: map[string]*dynamodb.AttributeValue{`)
	f.P(`"`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `": {`)

	switch cfg.PrimaryIndex.PartitionKey.GetAttrType() {
	case dynamopb.KeyDefinition_STRING:
		f.P(`S: aws.String(`, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `),`)
	case dynamopb.KeyDefinition_NUMBER:
		f.P(`N: aws.String(`, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `),`)
	case dynamopb.KeyDefinition_BINARY:
		f.P(`B: `, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `,`)
	default:
		log.Fatal("unknown attribute type")
	}

	f.P(`},`)
	if cfg.PrimaryIndex.SortKey.GetAttrName() != "" {
		f.P(`"`, cfg.PrimaryIndex.SortKey.GetAttrName(), `": {`)
		switch cfg.PrimaryIndex.SortKey.GetAttrType() {
		case dynamopb.KeyDefinition_STRING:
			f.P(`S: aws.String(`, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `),`)
		case dynamopb.KeyDefinition_NUMBER:
			f.P(`N: aws.String(`, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `),`)
		case dynamopb.KeyDefinition_BINARY:
			f.P(`B: `, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `,`)
		default:
			log.Fatal("unknown attribute type")
		}
		f.P(`},`)
	}
	f.P(`},`)
	f.P(`TableName: aws.String(TableName`, msg.GoIdent.GoName, `),`)
	f.P(`}`)
	f.P(`res, err := d.Client.GetItemWithContext(ctx, input)`)
	f.P(`if err != nil {`)
	f.P(`return nil, err`)
	f.P(`}`)
	f.P(`if len(res.Item) == 0 {`)
	f.P(`return nil, errors.New("not found")`)
	f.P(`}`)
	f.P(``, strings.ToLower(msg.GoIdent.GoName), ` := &`, msg.GoIdent.GoName, `{}`)
	f.P(`err = dynamodbattribute.UnmarshalMap(res.Item, `, strings.ToLower(msg.GoIdent.GoName), `)`)
	f.P(`if err != nil {`)
	f.P(`return nil, err`)
	f.P(`}`)
	f.P(`return `, strings.ToLower(msg.GoIdent.GoName), `, nil`)
	f.P(`}`)
	f.P()
}

func (g *Generator) generateMessageQuery(f *protogen.GeneratedFile, msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	if cfg.PrimaryIndex.SortKey != nil {
		g.generateMessageQueryBody(f, msg, cfg.PrimaryIndex)
	}

	for _, gsi := range cfg.GlobalSecondaryIndexes {
		g.generateMessageQueryBody(f, msg, gsi)
	}

}

func (g *Generator) generateMessageQueryBody(f *protogen.GeneratedFile, msg *protogen.Message, index *dynamopb.Index) {
	f.P(`func (d *DB) DDBQuery`,
		inflection.Plural(msg.GoIdent.GoName),
		`By`,
		strings.Title(index.PartitionKey.GetAttrName()),
		`And`,
		strings.Title(index.SortKey.GetAttrName()),
		`(ctx context.Context, expr expression.Expression, startKey string, limit int64) (`, strings.ToLower(inflection.Plural(msg.GoIdent.GoName)),
		` []*`, msg.GoIdent.GoName, `, lastEvaluatedKey string, err error) {`)

	f.P(`type lastKeyResponse map[string]struct {`)
	f.P(`}`)
	f.P(``)
	f.P(``)
	f.P(``)
	f.P(`input := &dynamodb.QueryInput{`)
	f.P(`KeyConditionExpression:    expr.KeyCondition(),`)
	f.P(`ExpressionAttributeNames:  expr.Names(),`)
	f.P(`ExpressionAttributeValues: expr.Values(),`)
	f.P(`Limit:                     aws.Int64(limit),`)
	f.P(`TableName:                 aws.String(TableName`, msg.GoIdent.GoName, `),`)
	if index.GetIndexName() != "" {
		f.P(`IndexName:                 aws.String("`, index.GetIndexName(), `"),`)
	}
	f.P(`}`)
	f.P(``)
	f.P(`if startKey != "" {`)
	f.P(``)
	f.P(`decoded, err := base64.StdEncoding.DecodeString(startKey)`)
	f.P(``)
	f.P(`if err != nil {`)
	f.P(`return nil, "", err`)
	f.P(`}`)
	f.P(``)
	f.P(`log.Println("Decoded", string(decoded))`)
	f.P(``)
	f.P(`exclusiveStartKey, err := dynamodbattribute.MarshalMap(string(decoded))`)
	f.P(``)
	f.P(`if err != nil {`)
	f.P(`return nil, "", err`)
	f.P(`}`)
	f.P(``)
	f.P(`input.ExclusiveStartKey = exclusiveStartKey`)
	f.P(``)
	f.P(`log.Println("EX:", exclusiveStartKey)`)
	f.P(`}`)
	f.P(``)
	f.P(`res, err := d.Client.QueryWithContext(ctx, input)`)
	f.P(``)
	f.P(`if err != nil {`)
	f.P(`return nil, "", err`)
	f.P(`}`)
	f.P(``)
	f.P(`lastEvaluatedKey = ""`)
	f.P(``)
	f.P(`if len(res.LastEvaluatedKey) > 0 {`)
	f.P(``)
	f.P(`log.Println(res.LastEvaluatedKey)`)
	f.P(``)
	f.P(`pageResponseMap := make(map[string]interface{})`)
	f.P(``)
	f.P(`for _, val := range res.LastEvaluatedKey {`)
	f.P(`log.Println(val.GoString())`)
	f.P(`}`)
	f.P(``)
	f.P(`err = dynamodbattribute.UnmarshalMap(res.LastEvaluatedKey, &pageResponseMap)`)
	f.P(``)
	f.P(`if err != nil {`)
	f.P(`return nil, "", err`)
	f.P(`}`)
	f.P(``)
	f.P(`b, err := json.Marshal(pageResponseMap)`)
	f.P(``)
	f.P(`if err != nil {`)
	f.P(`return nil, "", err`)
	f.P(`}`)
	f.P(``)
	f.P(`lastEvaluatedKey = base64.StdEncoding.EncodeToString(b)`)
	f.P(`}`)
	f.P(``)
	f.P(`err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &`, strings.ToLower(inflection.Plural(msg.GoIdent.GoName)), `)`)
	f.P(``)
	f.P(`if err != nil {`)
	f.P(`return nil, "", err`)
	f.P(`}`)
	f.P(``)
	f.P(`log.Println("LAST KEY", lastEvaluatedKey)`)
	f.P(``)
	f.P(`return `, strings.ToLower(inflection.Plural(msg.GoIdent.GoName)), `, lastEvaluatedKey, nil`)
	f.P(`}`)
	f.P()
}

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

func (g *Generator) generateLastKeyTypes(f *protogen.GeneratedFile, msg *protogen.Message, index *dynamopb.Index) {

	var indexName string

	if index.GetIndexName() == "" {
		indexName = "primary"
	} else {
		indexName = index.GetIndexName()
	}

	pk := index.GetPartitionKey()

	if pk == nil {
		log.Fatalf("Primary key required for index %s", index.GetIndexName())
	}
	if index.GetIndexName() == "" {
		f.P(`type primaryLastKey struct {`)
	} else {
		f.P(`type `, sanitiseString(indexName), `LastKey struct {`)
	}
	f.P("	", strings.Title(sanitiseString(index.GetPartitionKey().GetAttrName())), " ", DynamoGoMap[index.GetPartitionKey().GetAttrType()], " `json:\"", strings.ToLower(sanitiseString(index.GetPartitionKey().GetAttrName())), "\"`")
	if index.GetSortKey() != nil {
		if index.GetSortKey().GetAttrName() != "" {
			f.P("	", strings.Title(sanitiseString(index.GetSortKey().GetAttrName())), " ", DynamoGoMap[index.GetSortKey().GetAttrType()], " `json:\"", strings.ToLower(sanitiseString(index.GetSortKey().GetAttrName())), "\"`")
		}
	}
	f.P(`}`)
	f.P(``)

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
	f.P(``)
	g.generateLastKeyTypes(f, msg, index)
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
	f.P(`	decoded, err := base64.StdEncoding.DecodeString(startKey)`)
	f.P(``)
	f.P(`	if err != nil {`)
	f.P(`		return nil, "", err`)
	f.P(`	}`)
	f.P(``)
	if index.GetIndexName() == "" {
		f.P(`	key := new(primaryLastKey)`)
	} else {
		f.P(`	key := new(`, sanitiseString(index.GetIndexName()), `LastKey)`)
	}
	f.P(``)
	f.P(`	err = json.Unmarshal(decoded, key)`)
	f.P(``)
	f.P(`	if err != nil {`)
	f.P(`		return nil, "", err`)
	f.P(`	}`)
	f.P(``)
	f.P(`	exclusiveStartKey := map[string]*dynamodb.AttributeValue{}`)
	f.P(``)

	switch index.GetPartitionKey().GetAttrType() {

	case dynamopb.KeyDefinition_STRING:
		f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], ` = aws.String(key.`, strings.Title(index.GetPartitionKey().GetAttrName()), `)`)
	case dynamopb.KeyDefinition_NUMBER:
		f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], ` = aws.String(strconv.Itoa(key.`, strings.Title(index.GetPartitionKey().GetAttrName()), `))`)
	case dynamopb.KeyDefinition_BINARY:
		f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], ` = []byte(key.`, strings.Title(index.GetPartitionKey().GetAttrName()), `)`)
	}

	if index.GetSortKey() != nil {
		switch index.GetSortKey().GetAttrType() {
		case dynamopb.KeyDefinition_STRING:
			f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()], ` = aws.String(key.`, strings.Title(index.GetSortKey().GetAttrName()), `)`)
		case dynamopb.KeyDefinition_NUMBER:
			f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()], ` = aws.String(strconv.Itoa(key.`, strings.Title(index.GetSortKey().GetAttrName()), `))`)
		case dynamopb.KeyDefinition_BINARY:
			f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()], ` = []byte(key.`, strings.Title(index.GetSortKey().GetAttrName()), `)`)
		}
	}

	f.P(``)
	f.P(`	input.ExclusiveStartKey = exclusiveStartKey`)
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
	if index.GetIndexName() == "" {
		f.P(`	outKey := new(`, "primary", `LastKey)`)
	} else {
		f.P(`	outKey := new(`, sanitiseString(index.GetIndexName()), `LastKey)`)
	}
	f.P(``)
	switch index.GetPartitionKey().GetAttrType() {

	case dynamopb.KeyDefinition_STRING:
		f.P(`	outKey.`, strings.Title(index.GetPartitionKey().GetAttrName()), ` = *res.LastEvaluatedKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()])
	case dynamopb.KeyDefinition_NUMBER:
		f.P(`	outKey.`, strings.Title(index.GetPartitionKey().GetAttrName()), `, err = strconv.Atoi(*res.LastEvaluatedKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], `)`)
	case dynamopb.KeyDefinition_BINARY:
		f.P(`	outKey.`, strings.Title(index.GetPartitionKey().GetAttrName()), ` = res.LastEvaluatedKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()])
	}

	if index.GetSortKey() != nil {
		f.P(`	outKey.`, strings.Title(index.GetSortKey().GetAttrName()), ` = *res.LastEvaluatedKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()])
	}
	f.P(``)
	f.P(`	b, err := json.Marshal(outKey)`)
	f.P(``)
	f.P(`	if err != nil {`)
	f.P(`		return nil, "", err`)
	f.P(`	}`)
	f.P(``)
	f.P(`	lastEvaluatedKey = base64.StdEncoding.EncodeToString(b)`)
	f.P(`}`)
	f.P(``)
	f.P(`err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &`, strings.ToLower(inflection.Plural(msg.GoIdent.GoName)), `)`)
	f.P(``)
	f.P(`if err != nil {`)
	f.P(`return nil, "", err`)
	f.P(`}`)
	f.P(``)
	f.P(``)
	f.P(`return `, strings.ToLower(inflection.Plural(msg.GoIdent.GoName)), `, lastEvaluatedKey, nil`)
	f.P(`}`)
	f.P()
}

package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

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
	f.P(`const TableName`, msg.GoIdent.GoName, ` = "`, cfg.GetTableName(), `"`)
	f.P()
}

func (g *Generator) generateMessageCreate(f *protogen.GeneratedFile, msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	f.P(`func (d *DB) Create`, msg.GoIdent.GoName, `(ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, item *`, msg.GoIdent.GoName, `) (*`, msg.GoIdent.GoName, `, error) {`)
	f.P(`av, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "MarshalMap(item)", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"}))
	f.P(`if err != nil {`)
	f.P(`return item, err`)
	f.P(`}`)

	switch cfg.PrimaryIndex.SortKey {
	case nil:
		f.P(`expr, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "NewBuilder()", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `.WithCondition(expression.AttributeNotExists(expression.Name("`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `"))).Build()`)
	default:
		f.P(`expr, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "NewBuilder()", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `.WithCondition(expression.AttributeNotExists(expression.Name("`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `")).And(expression.AttributeNotExists(expression.Name("`, cfg.PrimaryIndex.SortKey.GetAttrName(), `")))).Build()`)
	}

	f.P(`if err != nil {`)
	f.P(`return item, err`)
	f.P(`}`)
	f.P(`input := &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "PutItemInput", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{`)
	f.P(`Item:                      av,`)
	f.P(`TableName:                 `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(TableName`, msg.GoIdent.GoName, `),`)
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
		f.P(`func (d *DB) Get`, msg.GoIdent.GoName, `By`, strings.Title(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` (ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, `, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.PartitionKey.GetAttrType()], `) (*`, msg.GoIdent.GoName, `, error) {`)
	default:
		f.P(`func (d *DB) Get`, msg.GoIdent.GoName, `By`, strings.Title(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `And`, strings.Title(cfg.PrimaryIndex.SortKey.GetAttrName()), ` (ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, `, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.PartitionKey.GetAttrType()], `, `, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.SortKey.GetAttrType()], `) (*`, msg.GoIdent.GoName, `, error) {`)
	}

	f.P(`input := &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "GetItemInput", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{`)
	f.P(`Key: map[string]*`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{`)
	f.P(`"`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `": {`)

	switch cfg.PrimaryIndex.PartitionKey.GetAttrType() {
	case dynamopb.KeyDefinition_STRING:
		f.P(`S: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `),`)
	case dynamopb.KeyDefinition_NUMBER:
		f.P(`N: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `),`)
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
			f.P(`S: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `),`)
		case dynamopb.KeyDefinition_NUMBER:
			f.P(`N: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `),`)
		case dynamopb.KeyDefinition_BINARY:
			f.P(`B: `, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `,`)
		default:
			log.Fatal("unknown attribute type")
		}
		f.P(`},`)
	}
	f.P(`},`)
	f.P(`TableName: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(TableName`, msg.GoIdent.GoName, `),`)
	f.P(`}`)
	f.P(`res, err := d.Client.GetItemWithContext(ctx, input)`)
	f.P(`if err != nil {`)
	f.P(`return nil, err`)
	f.P(`}`)
	f.P(`if len(res.Item) == 0 {`)
	f.P(`return nil, `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "New", GoImportPath: "errors"}), `("not found")`)
	f.P(`}`)
	f.P(``, strings.ToLower(msg.GoIdent.GoName), ` := &`, msg.GoIdent.GoName, `{}`)
	f.P(`err = `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "UnmarshalMap", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"}), `(res.Item, `, strings.ToLower(msg.GoIdent.GoName), `)`)
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

	f.P(`func (d *DB) Query`,
		inflection.Plural(msg.GoIdent.GoName),
		`By`,
		strings.Title(index.PartitionKey.GetAttrName()),
		`And`,
		strings.Title(index.SortKey.GetAttrName()),
		`(ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, expr `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Expression", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `, startKey string, limit int64) (`, strings.ToLower(inflection.Plural(msg.GoIdent.GoName)),
		` []*`, msg.GoIdent.GoName, `, lastEvaluatedKey string, err error) {`)
	f.P(``)
	g.generateLastKeyTypes(f, msg, index)
	f.P(``)
	f.P(``)
	f.P(``)
	f.P(`input := &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "QueryInput", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{`)
	f.P(`KeyConditionExpression:    expr.KeyCondition(),`)
	f.P(`ExpressionAttributeNames:  expr.Names(),`)
	f.P(`ExpressionAttributeValues: expr.Values(),`)
	f.P(`Limit:                     `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Int64", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(limit),`)
	f.P(`TableName:                 `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(TableName`, msg.GoIdent.GoName, `),`)
	if index.GetIndexName() != "" {
		f.P(`IndexName:                 aws.String("`, index.GetIndexName(), `"),`)
	}
	f.P(`}`)
	f.P(``)
	f.P(`if startKey != "" {`)
	f.P(`	decoded, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "StdEncoding", GoImportPath: "encoding/base64"}), `.DecodeString(startKey)`)
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
	f.P(`	err = `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Unmarshal", GoImportPath: "encoding/json"}), `(decoded, key)`)
	f.P(``)
	f.P(`	if err != nil {`)
	f.P(`		return nil, "", err`)
	f.P(`	}`)
	f.P(``)
	f.P(`	exclusiveStartKey := map[string]*`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{}`)
	f.P(``)

	switch index.GetPartitionKey().GetAttrType() {

	case dynamopb.KeyDefinition_STRING:
		f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"] = &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{ `, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], `: aws.String(key.`, strings.Title(index.GetPartitionKey().GetAttrName()), `)}`)
	case dynamopb.KeyDefinition_NUMBER:
		f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"] = &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{ `, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], `: aws.String(strconv.Itoa(key.`, strings.Title(index.GetPartitionKey().GetAttrName()), `))}`)
	case dynamopb.KeyDefinition_BINARY:
		f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"] = &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{ `, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], `: []byte(key.`, strings.Title(index.GetPartitionKey().GetAttrName()), `)}`)
	}

	if index.GetSortKey() != nil {
		switch index.GetSortKey().GetAttrType() {
		case dynamopb.KeyDefinition_STRING:
			f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"] = &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{ `, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()], `: aws.String(key.`, strings.Title(index.GetSortKey().GetAttrName()), `)}`)
		case dynamopb.KeyDefinition_NUMBER:
			f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"] = &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{ `, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()], `: aws.String(strconv.Itoa(key.`, strings.Title(index.GetSortKey().GetAttrName()), `))}`)
		case dynamopb.KeyDefinition_BINARY:
			f.P(`	exclusiveStartKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"] = &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{ `, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()], `: []byte(key.`, strings.Title(index.GetSortKey().GetAttrName()), `)}`)
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
		f.P(`	outKey.`, strings.Title(index.GetPartitionKey().GetAttrName()), `, err = `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Atoi", GoImportPath: "strconv"}), `(*res.LastEvaluatedKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()], `)`)
	case dynamopb.KeyDefinition_BINARY:
		f.P(`	outKey.`, strings.Title(index.GetPartitionKey().GetAttrName()), ` = res.LastEvaluatedKey["`, strings.ToLower(index.GetPartitionKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetPartitionKey().GetAttrType()])
	}

	if index.GetSortKey() != nil {
		f.P(`	outKey.`, strings.Title(index.GetSortKey().GetAttrName()), ` = *res.LastEvaluatedKey["`, strings.ToLower(index.GetSortKey().GetAttrName()), `"].`, DynamoAttributeGoMap[index.GetSortKey().GetAttrType()])
	}
	f.P(``)
	f.P(`	b, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Marshal", GoImportPath: "encoding/json"}), `(outKey)`)
	f.P(``)
	f.P(`	if err != nil {`)
	f.P(`		return nil, "", err`)
	f.P(`	}`)
	f.P(``)
	f.P(`	lastEvaluatedKey = `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "StdEncoding", GoImportPath: "encoding/base64"}), `.EncodeToString(b)`)
	f.P(`}`)
	f.P(``)
	f.P(`err = `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "UnmarshalListOfMaps", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"}), `(res.Items, &`, strings.ToLower(inflection.Plural(msg.GoIdent.GoName)), `)`)
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

func (g *Generator) generateMessageUpdate(f *protogen.GeneratedFile, msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	f.P(`func (d *DB) Update`, msg.GoIdent.GoName, `(ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, item *`, msg.GoIdent.GoName, `) (*`, msg.GoIdent.GoName, `, error) {`)
	f.P(`av, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "MarshalMap", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"}), `(item)`)
	f.P(`if err != nil {`)
	f.P(`return item, err`)
	f.P(`}`)

	switch cfg.PrimaryIndex.SortKey {
	case nil:
		f.P(`expr, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "NewBuilder", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `().WithCondition(`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeExists", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `(expression.Name("`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `"))).Build()`)
	default:
		f.P(`expr, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "NewBuilder", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `().WithCondition(`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeExists", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `(expression.Name("`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `")).And(expression.AttributeExists(expression.Name("`, cfg.PrimaryIndex.SortKey.GetAttrName(), `")))).Build()`)
	}

	f.P(`if err != nil {`)
	f.P(`return item, err`)
	f.P(`}`)
	f.P(`input := &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "PutItemInput", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{`)
	f.P(`Item:                      av,`)
	f.P(`TableName:                 `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(TableName`, msg.GoIdent.GoName, `),`)
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

func (g *Generator) generateMessageDelete(f *protogen.GeneratedFile, msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	switch cfg.PrimaryIndex.SortKey {
	case nil:
		f.P(`func (d *DB) Delete`, msg.GoIdent.GoName, `By`, strings.Title(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` (ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, `, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.PartitionKey.GetAttrType()], `) (error) {`)
	default:
		f.P(`func (d *DB) Delete`, msg.GoIdent.GoName, `By`, strings.Title(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `And`, strings.Title(cfg.PrimaryIndex.SortKey.GetAttrName()), ` (ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, `, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.PartitionKey.GetAttrType()], `, `, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), ` `, DynamoGoMap[cfg.PrimaryIndex.SortKey.GetAttrType()], `) (error) {`)
	}

	f.P(`input := &`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "DeleteItemInput", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{`)
	f.P(`Key: map[string]*`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "AttributeValue", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `{`)
	f.P(`"`, cfg.PrimaryIndex.PartitionKey.GetAttrName(), `": {`)

	switch cfg.PrimaryIndex.PartitionKey.GetAttrType() {
	case dynamopb.KeyDefinition_STRING:
		f.P(`S: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `),`)
	case dynamopb.KeyDefinition_NUMBER:
		f.P(`N: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.PartitionKey.GetAttrName()), `),`)
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
			f.P(`S: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `),`)
		case dynamopb.KeyDefinition_NUMBER:
			f.P(`N: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(`, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `),`)
		case dynamopb.KeyDefinition_BINARY:
			f.P(`B: `, strings.ToLower(cfg.PrimaryIndex.SortKey.GetAttrName()), `,`)
		default:
			log.Fatal("unknown attribute type")
		}
		f.P(`},`)
	}
	f.P(`},`)
	f.P(`TableName: `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(TableName`, msg.GoIdent.GoName, `),`)
	f.P(`}`)
	f.P(`_, err := d.Client.DeleteItemWithContext(ctx, input)`)
	f.P(`return err`)
	f.P(`}`)
	f.P()
}

func generateMessageTimestampMarshal(f *protogen.GeneratedFile, msg *protogen.Message, fields []*protogen.Field) {

	f.P(`func (a *`, msg.GoIdent.GoName, `) MarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {`)
	f.P(``)
	f.P(`	type Copy `, msg.GoIdent.GoName)
	f.P(``)
	f.P(`	m, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Marshal", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"}), `(&struct {`)

	for _, field := range fields {
		f.P("		", field.GoName, " ", f.QualifiedGoIdent(protogen.GoIdent{GoName: "Time", GoImportPath: "time"}), " `json:\"", field.Desc.Name(), "\"`")
	}

	f.P(`		*Copy`)
	f.P(`	}{`)

	for _, field := range fields {
		f.P(`		`, field.GoName, `: a.`, field.GoName, `.AsTime(),`)
	}

	f.P(`		Copy:       (*Copy)(a),`)
	f.P(`	})`)
	f.P(``)
	f.P(`	*av = *m`)
	f.P(``)
	f.P(`	return err`)
	f.P(`}`)
	f.P(``)
	f.P(`func (a *`, msg.GoIdent.GoName, `) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {`)
	f.P(`	type Copy `, msg.GoIdent.GoName, ``)
	f.P(``)
	f.P(`	aux := &struct {`)

	for _, field := range fields {
		f.P("		", field.GoName, " ", f.QualifiedGoIdent(protogen.GoIdent{GoName: "Time", GoImportPath: "time"}), " `json:\"", field.Desc.Name(), "\"`")
	}

	f.P(`		*Copy`)
	f.P(`	}{`)
	f.P(`		Copy: (*Copy)(a),`)
	f.P(`	}`)
	f.P(``)
	f.P(`	err := dynamodbattribute.Unmarshal(av, aux)`)
	f.P(``)
	f.P(`	if err != nil {`)
	f.P(`		return err`)
	f.P(`	}`)
	f.P(``)

	for _, field := range fields {
		f.P(`	a.`, field.GoName, ` = `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "New", GoImportPath: "google.golang.org/protobuf/types/known/timestamppb"}), `(aux.`, field.GoName, `)`)
	}

	f.P(``)
	f.P(`	return nil`)
	f.P(`}`)
}

func (g *Generator) generateProviderTerraform() {

	t, err := template.New("dynamodb").Parse(terraformProvider)

	if err != nil {
		log.Fatalf("Unable to parse terraform: %s", err)
	}

	var outPath string

	switch g.pathType {
	case "SOURCE_RELATIVE":
		outPath = fmt.Sprintf("%s/terraform", g.protoPath)
	default:
		outPath = fmt.Sprintf("%s/terraform", g.importPath)
	}

	err = os.MkdirAll(outPath, os.ModePerm)

	if err != nil {
		log.Fatalf("Unable to create terraform dir: %s", err)
	}

	var out []byte

	buf := bytes.NewBuffer(out)

	err = t.Execute(buf, g)

	if err != nil {
		log.Fatalf("Unable to execute terraform: %s", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.tf", outPath, "provider"), buf.Bytes(), os.ModePerm)

	if err != nil {
		log.Fatalf("Unable to write terraform: %s", err)
	}

}

func (g *Generator) generateMessageTerraform(msg *protogen.Message) {

	cfg, ok := proto.GetExtension(msg.Desc.Options(), dynamopb.E_Config).(*dynamopb.Cfg)

	if !ok {
		log.Fatal("could not read config")
	}

	t, err := template.New("dynamodb").Parse(terraformDynamoDB)

	if err != nil {
		log.Fatalf("Unable to parse terraform: %s", err)
	}

	var outPath string

	switch g.pathType {
	case "SOURCE_RELATIVE":
		outPath = fmt.Sprintf("%s/terraform", g.protoPath)
	default:
		outPath = fmt.Sprintf("%s/terraform", g.importPath)
	}

	err = os.MkdirAll(outPath, os.ModePerm)

	if err != nil {
		log.Fatalf("Unable to create terraform dir: %s", err)
	}

	var out []byte

	buf := bytes.NewBuffer(out)

	attrMap := map[string]string{}

	attrMap[cfg.PrimaryIndex.PartitionKey.AttrName] = cfg.PrimaryIndex.PartitionKey.AttrType.String()[:1]

	for _, gsi := range cfg.GlobalSecondaryIndexes {
		attrMap[gsi.PartitionKey.AttrName] = gsi.PartitionKey.AttrType.String()[:1]
	}

	configMap := struct {
		Attrs map[string]string
		Cfg   *dynamopb.Cfg
	}{
		Attrs: attrMap,
		Cfg:   cfg,
	}

	err = t.Execute(buf, configMap)

	if err != nil {
		log.Fatalf("Unable to parse terraform: %s", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.tf", outPath, strings.ToLower(cfg.TableName)), buf.Bytes(), os.ModePerm)

	if err != nil {
		log.Fatalf("Unable to write terraform: %s", err)
	}

}

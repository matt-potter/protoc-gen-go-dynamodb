package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/matt-potter/protoc-gen-go-dynamodb/dynamopb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	opts := protogen.Options{}

	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	base := plugin.NewGeneratedFile("base_db_generated.go", ".")

	generateBaseClient(base)

	for _, pfile := range plugin.Files {

		if len(pfile.Messages) == 0 {
			continue
		}

		segments := strings.Split(pfile.GeneratedFilenamePrefix, "/")

		basename := segments[len(segments)-1]

		file := plugin.NewGeneratedFile(fmt.Sprintf("%s_db_generated.go", basename), ".")

		messages := []*protogen.Message{}

		for _, msg := range pfile.Messages {
			if proto.HasExtension(msg.Desc.Options(), dynamopb.E_Storable) {
				messages = append(messages, msg)
			}
		}

		if len(messages) == 0 {
			file.Skip()
			continue
		}

		for _, storable := range messages {
			generateFuncs(file, storable)
		}

		file.Skip()

		var buf = new(bytes.Buffer)

		file.Write(buf.Bytes())
	}

	// Generate a response from our plugin and marshall as protobuf
	stdout := plugin.Response()

	out, err := proto.Marshal(stdout)

	if err != nil {
		panic(err)
	}

	// Write the response to stdout, to be picked up by protoc
	fmt.Fprintf(os.Stdout, string(out))
}

func generateFuncs(file *protogen.GeneratedFile, msg *protogen.Message) {

	f := jen.NewFile("db")

	f.Comment("// Code generated by protoc-gen-go-dynamodb. DO NOT EDIT.")
	f.Func()

	file.P(`// Code generated by protoc-gen-go-dynamodb. DO NOT EDIT.`)
	file.P(`package db`)
	file.P(`func (d *db) Save(ctx context.Context, `, `a *api.Activity) (*api.Activity, error) {`)

	file.P(`av, err := dynamodbattribute.MarshalMap(a)`)
	file.P(`if err != nil {`)
	file.P(`	return a, err`)
	file.P(`}`)

	file.P(`tableName := "Movies"`)

	file.P(`input := &dynamodb.PutItemInput{`)
	file.P(`	Item:      av,`)
	file.P(`	TableName: aws.String(tableName),`)
	file.P(`}`)

	file.P(`_, err = d.Client.PutItem(input)`)

	file.P(`if err != nil {`)
	file.P(`	return a, err`)
	file.P(`}`)

	file.P(`return a, nil`)
	file.P(`}`)
}

func generateBaseClient(file *protogen.GeneratedFile) {
	file.P(`// Code generated by protoc-gen-go-dynamodb. DO NOT EDIT.`)
	file.P(`package db`)
	file.P(`import (`)
	file.P(`"github.com/aws/aws-sdk-go/aws/session"`)
	file.P(`"github.com/aws/aws-sdk-go/service/dynamodb"`)
	file.P(`)`)
	file.P(`type db struct {`)
	file.P(`Client *dynamodb.DynamoDB`)
	file.P(`}`)
	file.P(`func NewDB() *db {

	sess := session.New()
	d := dynamodb.New(sess)

	return &db{
		Client: d,
	}

	}`)

}

package generator

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/matt-potter/protoc-gen-go-dynamodb/dynamopb"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// Generator is a struct to hold our config values
type Generator struct {
	packageName string
	importPath  string
	protoPath   string
	pathType    string
	plugin      *protogen.Plugin
	importMap   map[string]string
}

// NewGenerator is a wrapper function to generate our dynamo code
func NewGenerator(plugin *protogen.Plugin, pathType string) (*Generator, error) {

	gen := &Generator{}

	gen.pathType = pathType

	gen.importMap = make(map[string]string)

	gen.importMap["context"] = "context"
	gen.importMap["aws"] = "github.com/aws/aws-sdk-go/aws"
	gen.importMap["dynamodb"] = "github.com/aws/aws-sdk-go/service/dynamodb"
	gen.importMap["dynamodbattribute"] = "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	gen.importMap["expression"] = "github.com/aws/aws-sdk-go/service/dynamodb/expression"

	pkg, err := getPackageName(plugin.Files)

	if err != nil {
		return nil, err
	}

	gen.packageName = pkg

	ip, err := getImportPath(plugin.Files)

	if err != nil {
		return nil, err
	}

	gen.importPath = ip

	pp, err := getProtoPath(plugin.Files)

	if err != nil {
		return nil, err
	}

	gen.protoPath = pp

	gen.plugin = plugin

	log.Printf("%+v", gen)

	return gen, nil
}

func (g *Generator) Generate() {

	g.generateHeaderAndClient()

	g.generateMethods()

	for _, pfile := range g.plugin.Files {

		if !pfile.Generate {
			continue
		}

		if len(pfile.Messages) == 0 {
			continue
		}

		// file := g.plugin.NewGeneratedFile(fmt.Sprintf("%s.db.go", pfile.GeneratedFilenamePrefix), pfile.GoImportPath)

		// messages := []*protogen.Message{}

		// for _, msg := range pfile.Messages {

		// 	if proto.HasExtension(msg.Desc.Options(), dynamopb.E_Storable) {

		// 		// for _, f := range msg.Fields {

		// 		// 	if f.Desc.Kind() == protoreflect.MessageKind {
		// 		// 		log.Printf("%+v", f.Message.Desc.FullName())
		// 		// 	}

		// 		// }
		// 		messages = append(messages, msg)
		// 	}

		// }

		// if len(messages) == 0 {
		// 	file.Skip()
		// 	continue
		// }

		// for _, storable := range messages {

		// 	GenerateFuncs(file, pfile, storable)

		// }

		// file.Skip()

	}

	stdout := g.plugin.Response()

	out, err := proto.Marshal(stdout)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, string(out))

}

func getProtoPath(files []*protogen.File) (string, error) {

	protoPath := ""

	for _, f := range files {

		if !f.Generate {
			continue
		}

		if protoPath == "" {
			if path.Dir(f.Desc.Path()) == "" {
				return "", errors.New("Package name could not be determined")
			}
			protoPath = path.Dir(f.Desc.Path())
		}

		if path.Dir(f.Desc.Path()) != protoPath {
			log.Printf("Mismatch found: %s, %s", path.Dir(f.Desc.Path()), protoPath)
			return "", errors.New("Package name mismatch in provided .proto files")
		}

	}

	return protoPath, nil

}

func getPackageName(files []*protogen.File) (string, error) {

	packageName := ""

	for _, f := range files {

		if !f.Generate {
			continue
		}

		if packageName == "" {
			if f.GoPackageName == "" {
				return "", errors.New("Package name could not be determined")
			}
			packageName = string(f.GoPackageName)
		}

		if string(f.GoPackageName) != packageName {
			log.Printf("Mismatch found: %s, %s", string(f.GoPackageName), packageName)
			return "", errors.New("Package name mismatch in provided .proto files")
		}

	}

	return packageName, nil

}

func getImportPath(files []*protogen.File) (string, error) {

	importPath := ""

	for _, f := range files {

		if !f.Generate {
			continue
		}

		if importPath == "" {
			if f.GoImportPath == "" {
				return "", errors.New("Package name could not be determined")
			}
			importPath = string(f.GoImportPath)
		}

		if string(f.GoImportPath) != importPath {
			log.Printf("Mismatch found: %s, %s", string(f.GoImportPath), importPath)
			return "", errors.New("Import path mismatch in provided .proto files")
		}

	}

	return importPath, nil

}

func (g *Generator) generateHeaderAndClient() {

	switch g.pathType {
	case "IMPORT":
		f := g.plugin.NewGeneratedFile(fmt.Sprintf("%s/%s.db.go", g.protoPath, g.packageName), protogen.GoImportPath(g.importPath))
		f.P(fmt.Sprintf("package %s", g.packageName))
	case "SOURCE_RELATIVE":
		f := g.plugin.NewGeneratedFile(fmt.Sprintf("%s/%s.db.go", g.importPath, g.packageName), protogen.GoImportPath(g.importPath))
		f.P(fmt.Sprintf("package %s", g.packageName))
	default:
		f := g.plugin.NewGeneratedFile(fmt.Sprintf("%s/%s.db.go", g.importPath, g.packageName), protogen.GoImportPath(g.importPath))
		f.P(fmt.Sprintf("package %s", g.packageName))
	}

}

func (g *Generator) generateMethods() {

	for _, file := range g.plugin.Files {

		if !file.Generate {
			continue
		}

		if len(file.Messages) == 0 {
			continue
		}

		for _, msg := range file.Messages {

			if proto.HasExtension(msg.Desc.Options(), dynamopb.E_Storable) {
				if !proto.GetExtension(msg.Desc.Options(), dynamopb.E_Storable).(bool) {
					continue
				}
			}

			var pk *dynamopb.Pk
			var gsis []*dynamopb.Gsi

			if proto.HasExtension(msg.Desc.Options(), dynamopb.E_PrimaryIndex) {
				pk = proto.GetExtension(msg.Desc.Options(), dynamopb.E_PrimaryIndex).(*dynamopb.Pk)
			}

			if proto.HasExtension(msg.Desc.Options(), dynamopb.E_GlobalSecondaryIndexes) {
				gsis = proto.GetExtension(msg.Desc.Options(), dynamopb.E_GlobalSecondaryIndexes).([]*dynamopb.Gsi)
			}

			if pk == nil {
				log.Fatalf("storable is set but no primary_key is defined in %s", file.Desc.Path())
			}

			f := g.plugin.NewGeneratedFile(fmt.Sprintf("%s.db.go", file.GeneratedFilenamePrefix), file.GoImportPath)

			log.Println(pk, gsis)

			g.generateBaseActions(f)

		}

	}

	// pi := &dynamopb.KeyDefinition{}
	// // gsis := []*dynamopb.Gsi{}

	// if proto.HasExtension(msg.Desc.Options(), dynamopb.E_PrimaryIndex) {
	// 	pi = proto.GetExtension(msg.Desc.Options(), dynamopb.E_PrimaryIndex).(*dynamopb.KeyDefinition)
	// }

	// // if proto.HasExtension(msg.Desc.Options(), dynamopb.E_GlobalSecondaryIndexes) {
	// // 	gsis = proto.GetExtension(msg.Desc.Options(), dynamopb.E_GlobalSecondaryIndexes).([]*dynamopb.Gsi)
	// // }

	// f.P(`package `, pf.GoPackageName)
	// f.P(`const TableName`, msg.GoIdent, ` = "`, msg.GoIdent, `"`)
	// f.P(`func Create`, msg.GoIdent, `(ctx `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "Context", GoImportPath: "context"}), `, db *`, f.QualifiedGoIdent(protogen.GoIdent{GoName: "DynamoDB", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb"}), `, item *`, msg.GoIdent, `) (*`, msg.GoIdent, `, error) {`)
	// f.P(`av, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "MarshalMap(item)", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"}))
	// f.P(`if err != nil {`)
	// f.P(`return item, err`)
	// f.P(`}`)
	// f.P(`expr, err := `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "NewBuilder()", GoImportPath: "github.com/aws/aws-sdk-go/service/dynamodb/expression"}), `.WithCondition(expression.AttributeNotExists(expression.Name("`, pi.GetName(), `"))).Build()`)
	// f.P(`if err != nil {`)
	// f.P(`	return item, err`)
	// f.P(`}`)
	// f.P(`input := &dynamodb.PutItemInput{`)
	// f.P(`Item:                      av,`)
	// f.P(`TableName:                 `, f.QualifiedGoIdent(protogen.GoIdent{GoName: "String", GoImportPath: "github.com/aws/aws-sdk-go/aws"}), `(TableName`, msg.GoIdent, `),`)
	// f.P(`ConditionExpression:       expr.Condition(),`)
	// f.P(`ExpressionAttributeNames:  expr.Names(),`)
	// f.P(`ExpressionAttributeValues: expr.Values(),`)
	// f.P(`}`)
	// f.P(`_, err = db.PutItemWithContext(ctx, input)`)
	// f.P(`if err != nil {`)
	// f.P(`	return item, err`)
	// f.P(`}`)
	// f.P(`return item, nil`)
	// f.P(`}`)
}

func (g *Generator) generateBaseActions(f *protogen.GeneratedFile) {
	f.P(`package `, g.packageName)
}

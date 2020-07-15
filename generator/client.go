package generator

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func (g *Generator) generateHeaderAndClient() {
	var f *protogen.GeneratedFile

	switch g.pathType {
	case "SOURCE_RELATIVE":
		f = g.plugin.NewGeneratedFile(fmt.Sprintf("%s/%s.db.go", g.protoPath, g.packageName), protogen.GoImportPath(g.importPath))
	default:
		f = g.plugin.NewGeneratedFile(fmt.Sprintf("%s/%s.db.go", g.importPath, g.packageName), protogen.GoImportPath(g.importPath))
	}

	f.P(`package `, g.packageName)
	f.P(``)
	f.P(`import "github.com/aws/aws-sdk-go/service/dynamodb"`)
	f.P(``)
	f.P(`type DB struct {`)
	f.P(`	Client *dynamodb.DynamoDB`)
	f.P(`}`)
	f.P(``)
	f.P(`func NewDB(ddb *dynamodb.DynamoDB) *DB {`)
	f.P(`	return &DB{`)
	f.P(`		Client: ddb,`)
	f.P(`	}`)
	f.P(`}`)
}

package generator

import (
	"fmt"
	"os"

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
}

var DynamoGoMap = map[dynamopb.KeyDefinitionAttributeType]string{
	dynamopb.KeyDefinition_STRING: "string",
	dynamopb.KeyDefinition_BINARY: "[]byte",
	dynamopb.KeyDefinition_NUMBER: "int",
}

// NewGenerator is a wrapper function to generate our dynamo code
func NewGenerator(plugin *protogen.Plugin, pathType string) (*Generator, error) {

	gen := &Generator{}

	gen.pathType = pathType

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

	return gen, nil
}

func (g *Generator) Generate() {

	g.generateHeaderAndClient()

	for _, pfile := range g.plugin.Files {

		if !pfile.Generate {
			continue
		}

		if len(pfile.Messages) == 0 {
			continue
		}

		for _, m := range g.getStorableMessages(pfile) {

			filename := pfile.GeneratedFilenamePrefix + ".db.go"

			f := g.plugin.NewGeneratedFile(filename, pfile.GoImportPath)

			g.generateMessageHeader(f, m)

			g.generateMessageCreate(f, m)

			g.generateMessageGet(f, m)

			g.generateMessageQuery(f, m)

		}

	}

	stdout := g.plugin.Response()

	out, err := proto.Marshal(stdout)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, string(out))

}

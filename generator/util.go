package generator

import (
	"errors"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/matt-potter/protoc-gen-go-dynamodb/dynamopb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

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

func getTimestampFields(m *protogen.Message) []*protogen.Field {

	out := []*protogen.Field{}

	for _, field := range m.Fields {
		if field.Message != nil {
			if field.Message.Desc.FullName() == "google.protobuf.Timestamp" {
				out = append(out, field)
			}

		}
	}

	return out
}

func (g *Generator) getStorableMessages(f *protogen.File) []*protogen.Message {

	storable := []*protogen.Message{}

	for _, msg := range f.Messages {
		if proto.HasExtension(msg.Desc.Options(), dynamopb.E_Config) {
			storable = append(storable, msg)
		} else {
			continue
		}
	}

	return storable
}

func sanitiseString(s string) string {
	re := regexp.MustCompile(`\W`)
	split := re.Split(s, -1)

	t := []string{}

	for i, str := range split {
		if i == 0 {
			t = append(t, strings.ToLower(str))
		} else {
			t = append(t, strings.Title(str))
		}
	}

	return strings.Join(t, "")
}

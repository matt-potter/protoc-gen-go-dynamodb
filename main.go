package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/matt-potter/protoc-gen-go-dynamodb/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	var SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	input, _ := ioutil.ReadAll(os.Stdin)

	var req pluginpb.CodeGeneratorRequest

	proto.Unmarshal(input, &req)

	opts := protogen.Options{}

	plugin, err := opts.New(&req)

	plugin.SupportedFeatures = SupportedFeatures

	if err != nil {
		panic(err)
	}

	var pathType string

	for _, param := range strings.Split(req.GetParameter(), ",") {

		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = param[i+1:]
			param = param[0:i]
		}
		switch param {
		case "paths":
			switch value {
			case "import":
				pathType = "IMPORT"
			case "source_relative":
				pathType = "SOURCE_RELATIVE"
			default:
				panic(fmt.Sprintf(`unknown path type %q: want "import" or "source_relative"`, value))
			}
		}
	}

	g, err := generator.NewGenerator(plugin, pathType)

	if err != nil {
		panic(err)
	}

	g.Generate()

}

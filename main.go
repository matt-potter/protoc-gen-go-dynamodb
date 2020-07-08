package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	// Protoc passes a slice of File structs for us to process
	for _, file := range plugin.Files {

		if len(file.Messages) == 0 {
			continue
		}

		var buf bytes.Buffer

		// 2. Write the package name
		buf.Write([]byte("package ddb\n\n"))
		pkg := fmt.Sprintf("import (\"%s\")", file.GoPackageName)
		buf.Write([]byte(pkg))

		// 3. For each message add our Foo() method
		for _, msg := range file.Messages {

			if strings.HasSuffix(string(msg.Desc.Name()), "Response") || strings.HasSuffix(string(msg.Desc.Name()), "Request") {
				continue
			}

			buf.Write([]byte(fmt.Sprintf(`
			func (x %s) Foo() string {
			   return "bar"
			}`, msg.Desc.Name())))

		}

		// 4. Specify the output filename, in this case test.foo.go
		filename := file.GeneratedFilenamePrefix + ".foo.go"
		file := plugin.NewGeneratedFile(filename, ".")

		// 5. Pass the data from our buffer to the plugin file struct
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

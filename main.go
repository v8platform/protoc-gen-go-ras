package main

import (
	"github.com/v8platform/protoc-gen-go-ras/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	protogen.Options{}.
		Run(func(plugin *protogen.Plugin) error {
			plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

			gen := generator.NewGenerator(plugin)
			gen.Generate()
			gen.GenerateOnce()

			return nil
		})
}

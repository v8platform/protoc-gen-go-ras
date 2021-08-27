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
			for _, f := range plugin.Files {
				if f.Generate && shouldProcess(f) {
					gen.GenerateFile(plugin, f)
				}
			}
			return nil
		})
}

func shouldProcess(file *protogen.File) bool {
	ignoredFiles := []string{"graphql/graphql.proto", "graphql.proto", "google/protobuf/descriptor.proto", "google/protobuf/wrappers.proto", "google/protobuf/timestamp.proto", "github.com/kitt-technology/protoc-gen-graphql/graphql/graphql.proto"}
	for _, ignored := range ignoredFiles {
		if *file.Proto.Name == ignored {
			return false
		}
	}
	// if proto.HasExtension(file.Proto.Options, graphql.E_Disabled) {
	// 	return !proto.GetExtension(file.Proto.Options, graphql.E_Disabled).(bool)
	// }
	return true
}

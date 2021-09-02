FROM scratch

LABEL "build.buf.plugins.runtime_library_versions.1.name"="google.golang.org/protobuf"
LABEL "build.buf.plugins.runtime_library_versions.1.version"="v1.27.1"

LABEL "build.buf.plugins.runtime_library_versions.0.name"="github.com/spf13/cast"
LABEL "build.buf.plugins.runtime_library_versions.0.version"="v1.4.1"

LABEL "build.buf.plugins.runtime_library_versions.2.name"="github.com/v8platform/encoder"
LABEL "build.buf.plugins.runtime_library_versions.2.version"="v0.0.1"

COPY protoc-gen-go-ras /
ENTRYPOINT ["/protoc-gen-go-ras"]
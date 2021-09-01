module github.com/v8platform/protoc-gen-go-ras

go 1.17

require (
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/spf13/cast v1.4.1
	github.com/v8platform/encoder v0.0.0-20210830084048-75ecfcb84e16
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/v8platform/encoder v0.0.0-20210830084048-75ecfcb84e16 => ../encoder

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

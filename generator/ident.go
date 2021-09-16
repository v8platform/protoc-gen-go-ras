package generator

import "google.golang.org/protobuf/compiler/protogen"

const (
	reflectPackage     = protogen.GoImportPath("reflect")
	fmtPackage         = protogen.GoImportPath("fmt")
	bytesPackage       = protogen.GoImportPath("bytes")
	stringsPackage     = protogen.GoImportPath("strings")
	syncPackage        = protogen.GoImportPath("sync")
	timePackage        = protogen.GoImportPath("time")
	regexpPackage      = protogen.GoImportPath("regexp")
	ctxPackage         = protogen.GoImportPath("context")
	anypbPackage       = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")
	protoPackage       = protogen.GoImportPath("google.golang.org/protobuf/proto")
	emptypbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/emptypb")
	deprecationComment = "// Deprecated: Do not use."
)

var (
	fmtSprintf     = protogen.GoIdent{GoName: "Sprintf", GoImportPath: fmtPackage}
	fmtErrorf      = protogen.GoIdent{GoName: "Errorf", GoImportPath: fmtPackage}
	bytesBuffer    = protogen.GoIdent{GoName: "Buffer", GoImportPath: bytesPackage}
	bytesNewBuffer = protogen.GoIdent{GoName: "NewBuffer", GoImportPath: bytesPackage}
)

const (
	encoderPackage = protogen.GoImportPath("github.com/v8platform/encoder/ras/codec256")
	castPackage    = protogen.GoImportPath("github.com/spf13/cast")
	ioPackage      = protogen.GoImportPath("io")
)

var ioReader = protogen.GoIdent{GoName: "Reader", GoImportPath: ioPackage}

var (
	formatBool     = protogen.GoIdent{GoName: "FormatBool", GoImportPath: encoderPackage}
	formatString   = protogen.GoIdent{GoName: "FormatString", GoImportPath: encoderPackage}
	formatInt      = protogen.GoIdent{GoName: "FormatInt", GoImportPath: encoderPackage}
	formatLong     = protogen.GoIdent{GoName: "FormatLong", GoImportPath: encoderPackage}
	formatFloat    = protogen.GoIdent{GoName: "FormatFloat", GoImportPath: encoderPackage}
	formatDouble   = protogen.GoIdent{GoName: "FormatDouble", GoImportPath: encoderPackage}
	formatTime     = protogen.GoIdent{GoName: "FormatTime", GoImportPath: encoderPackage}
	formatBytes    = protogen.GoIdent{GoName: "FormatBytes", GoImportPath: encoderPackage}
	formatSize     = protogen.GoIdent{GoName: "FormatSize", GoImportPath: encoderPackage}
	formatShort    = protogen.GoIdent{GoName: "FormatShort", GoImportPath: encoderPackage}
	formatNullable = protogen.GoIdent{GoName: "FormatNullable", GoImportPath: encoderPackage}
	formatUUID     = protogen.GoIdent{GoName: "FormatUuid", GoImportPath: encoderPackage}
	formatType     = protogen.GoIdent{GoName: "FormatType", GoImportPath: encoderPackage}
	formatByte     = protogen.GoIdent{GoName: "FormatByte", GoImportPath: encoderPackage}

	parseBool     = protogen.GoIdent{GoName: "ParseBool", GoImportPath: encoderPackage}
	parseString   = protogen.GoIdent{GoName: "ParseString", GoImportPath: encoderPackage}
	parseInt      = protogen.GoIdent{GoName: "ParseInt", GoImportPath: encoderPackage}
	parseLong     = protogen.GoIdent{GoName: "ParseLong", GoImportPath: encoderPackage}
	parseFloat    = protogen.GoIdent{GoName: "ParseFloat", GoImportPath: encoderPackage}
	parseDouble   = protogen.GoIdent{GoName: "ParseDouble", GoImportPath: encoderPackage}
	parseTime     = protogen.GoIdent{GoName: "ParseTime", GoImportPath: encoderPackage}
	parseSize     = protogen.GoIdent{GoName: "ParseSize", GoImportPath: encoderPackage}
	parseNullable = protogen.GoIdent{GoName: "ParseNullable", GoImportPath: encoderPackage}
	parseByte     = protogen.GoIdent{GoName: "ParseByte", GoImportPath: encoderPackage}
	parseBytes    = protogen.GoIdent{GoName: "ParseBytes", GoImportPath: encoderPackage}
	parseShort    = protogen.GoIdent{GoName: "ParseShort", GoImportPath: encoderPackage}
	parseUuid     = protogen.GoIdent{GoName: "ParseUUID", GoImportPath: encoderPackage}
	parseType     = protogen.GoIdent{GoName: "ParseType", GoImportPath: encoderPackage}

	codecVersion = protogen.GoIdent{GoName: "Version", GoImportPath: encoderPackage}

	encoders = map[string]protogen.GoIdent{
		"bool":     formatBool,
		"size":     formatSize,
		"nullable": formatNullable,
		"byte":     formatByte,
		"time":     formatTime,
		"bytes":    formatBytes,
		"short":    formatShort,
		"float":    formatFloat,
		"double":   formatDouble,
		"string":   formatString,
		"uuid":     formatUUID,
		"int":      formatInt,
		"long":     formatLong,
		"type":     formatType,
	}

	decoders = map[string]protogen.GoIdent{
		"bool":     parseBool,
		"size":     parseSize,
		"nullable": parseNullable,
		"byte":     parseByte,
		"time":     parseTime,
		"bytes":    parseBytes,
		"short":    parseShort,
		"float":    parseFloat,
		"double":   parseDouble,
		"string":   parseString,
		"uuid":     parseUuid,
		"int":      parseInt,
		"long":     parseLong,
		"type":     parseType,
	}
)

package generator

import "google.golang.org/protobuf/compiler/protogen"

const (
	reflectPackage = protogen.GoImportPath("reflect")
	fmtPackage     = protogen.GoImportPath("fmt")
	sortPackage    = protogen.GoImportPath("sort")
	stringsPackage = protogen.GoImportPath("strings")
	syncPackage    = protogen.GoImportPath("sync")
	timePackage    = protogen.GoImportPath("time")
	utf8Package    = protogen.GoImportPath("unicode/utf8")
)

var (
	fmtSprintf = protogen.GoIdent{GoName: "Sprintf", GoImportPath: fmtPackage}
	fmtErrorf  = protogen.GoIdent{GoName: "Errorf", GoImportPath: fmtPackage}
)

const (
	encoderPackage = protogen.GoImportPath("github.com/v8platform/encoder/ras")
	ioPackage      = protogen.GoImportPath("io")
)

var (
	formatBool    = protogen.GoIdent{GoName: "FormatBool", GoImportPath: encoderPackage}
	formatString  = protogen.GoIdent{GoName: "FormatString", GoImportPath: encoderPackage}
	formatInt32   = protogen.GoIdent{GoName: "FormatInt32", GoImportPath: encoderPackage}
	formatInt64   = protogen.GoIdent{GoName: "FormatInt64", GoImportPath: encoderPackage}
	formatFloat32 = protogen.GoIdent{GoName: "FormatFloat32", GoImportPath: encoderPackage}
	formatFloat64 = protogen.GoIdent{GoName: "FormatFloat64", GoImportPath: encoderPackage}
	formatTime    = protogen.GoIdent{GoName: "FormatTime", GoImportPath: encoderPackage}
	formatBytes   = protogen.GoIdent{GoName: "FormatBytes", GoImportPath: encoderPackage}
	formatSize    = protogen.GoIdent{GoName: "FormatSize", GoImportPath: encoderPackage}
	formatShort   = protogen.GoIdent{GoName: "FormatShort", GoImportPath: encoderPackage}

	parseBool     = protogen.GoIdent{GoName: "ParseBool", GoImportPath: encoderPackage}
	parseString   = protogen.GoIdent{GoName: "ParseString", GoImportPath: encoderPackage}
	parseInt      = protogen.GoIdent{GoName: "ParseInt", GoImportPath: encoderPackage}
	parseLong     = protogen.GoIdent{GoName: "ParseLong", GoImportPath: encoderPackage}
	parseFloat    = protogen.GoIdent{GoName: "ParseFloat", GoImportPath: encoderPackage}
	parseDouble   = protogen.GoIdent{GoName: "ParseDouble", GoImportPath: encoderPackage}
	parseTime     = protogen.GoIdent{GoName: "ParseTime", GoImportPath: encoderPackage}
	parseSize     = protogen.GoIdent{GoName: "ParseSize", GoImportPath: encoderPackage}
	parseNullSize = protogen.GoIdent{GoName: "ParseNullable", GoImportPath: encoderPackage}
	parseByte     = protogen.GoIdent{GoName: "ParseByte", GoImportPath: encoderPackage}
	parseBytes    = protogen.GoIdent{GoName: "ParseBytes", GoImportPath: encoderPackage}
	parseShort    = protogen.GoIdent{GoName: "ParseShort", GoImportPath: encoderPackage}
	parseUuid     = protogen.GoIdent{GoName: "ParseUuid", GoImportPath: encoderPackage}
	parseType     = protogen.GoIdent{GoName: "ParseType", GoImportPath: encoderPackage}

	encoders = map[string]protogen.GoIdent{
		"bool":     parseBool,
		"size":     parseSize,
		"nullable": parseNullSize,
		"byte":     parseByte,
		"time":     parseTime,
		"bytes":    parseBytes,
		"short":    parseShort,
		"float32":  parseFloat32,
		"float64":  parseFloat64,
		"string":   parseString,
		"uuid":     parseUuid,
		"int":      parseType,
		"int64":    parseInt64,
	}

	decoders = map[string]protogen.GoIdent{
		"bool":     parseBool,
		"size":     parseSize,
		"nullable": parseNullSize,
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

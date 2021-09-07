# protoc-gen-go-ras

[![Release](https://img.shields.io/github/release/v8platform/protoc-gen-go-ras.svg?style=for-the-badge)](https://github.com/v8platform/protoc-gen-go-ras/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Build status](https://img.shields.io/github/workflow/status/v8platform/protoc-gen-go-ras/goreleaser?style=for-the-badge)](https://github.com/v8platform/protoc-gen-go-ras/actions?workflow=goreleaser)
[![Codecov branch](https://img.shields.io/codecov/c/github/v8platform/protoc-gen-go-ras/master.svg?style=for-the-badge)](https://codecov.io/gh/v8platform/protoc-gen-go-ras)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/v8platform/protoc-gen-go-ras)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=for-the-badge)](https://saythanks.io/to/khorevaa)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## Описание

Это плагин для `protoc` и требует его установки

Обеспечивает генерация методов и различных помощников для RAS

Создает следующие методы для `proto` файлов

```go
func (m *message) Parse(r io.Reader, veriosn int32) err {
	// Сгенерированный код
}

func (m *message) Formatter(w io.Writer, veriosn int32) err {
    // Сгенерированный код
}
```


## Как установить

* Установить из [`releases`](https://github.com/v8platform/protoc-gen-go-ras/releases/)
* Использовать готовый образ `docker`
    * `docker pull v8platform/protoc-gen-go-ras:latest`
    * `docker pull ghcr.io/v8platform/protoc-gen-go-ras:latest`
* Сборка для через `go get`
```shell
go get github.com/v8platform/protoc-gen-go-ras/...
go install github.com/v8platform/protoc-gen-go-ras
```

### Описание возможностей 

TODO


## License

Лицензия [`LICENSE`](LICENSE)
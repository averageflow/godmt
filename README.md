# GoDMT

[![Build](https://img.shields.io/github/workflow/status/averageflow/godmt/Test)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/averageflow/godmt)](https://goreportcard.com/report/github.com/averageflow/godmt)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/averageflow/godmt)](https://pkg.go.dev/github.com/averageflow/godmt)
[![Maintainability](https://api.codeclimate.com/v1/badges/8ee5c4680a29aef11331/maintainability)](https://codeclimate.com/github/averageflow/godmt/maintainability)
[![codecov](https://codecov.io/gh/averageflow/godmt/branch/master/graph/badge.svg?token=F4HW4K40T6)](https://codecov.io/gh/averageflow/godmt)
[![License](https://img.shields.io/github/license/averageflow/godmt.svg)](https://github.com/averageflow/godmt/blob/master/LICENSE.md)

GoDMT, the one and only Go Data Model Translator. The goal of this project is to provide a tool that can parse Go files
that include `var`, `const`, `map`, `struct` and `type` into an abstract syntax tree, aka AST.

<p align="center">
  <img width="250" height="150" src="web/DMT.png">
</p>

That AST will then be transformed into data models for several programming languages. Currently GoDMT can perform
translations to:

- TypeScript
- Swift (using Decodable structs)
- JSON
- PHP

Some small adjustments may need to be made to integrate the output into a project, but this should already save you a
lot of time and hassle, and will help you stay in sync with the Go version of your data models, in other languages.
Comments will be carried over 😉.

Currently, the supported operating systems are all of UNIX family:

- Linux
- BSD
- macOS

## Talk is cheap, show code

Feel free to browse some examples that I am happy to provide here:

- [Complex Structures](examples/ComplexStructures.md)
- [Constants](examples/Constants.md)
- [Maps](examples/Maps.md)
- [Pointers](examples/Pointers.md)
- [Slices](examples/Slices.md)
- [Structs](examples/Structs.md)

## Usage

```
go run main.go -dir={scanDirectory} -translation={language} -preserve -tree
```

- `scanDirectory` represents a string that is the relative path of the directory whose Go files you want to scan. The
  scan occurs in a recursive manner, so all files from all contained folders will be scanned.
- `language` represents the output mode. If the `-translation` flag is not specified it will default to JSON. Currently
  supported options are:
    - `ts` or `typescript` for TypeScript conversion
    - `swift` for Swift conversion
    - `json` for JSON conversion
    - `php` for PHP conversion
- `preserve` is an optional boolean flag which will make the output structs preserve the original names, instead of
  using the (`json:"tag"`).
- `tree` is an optional boolean that when present will prevent any file operations being performed, and instead will
  show you the full abstract syntax tree of your files in the standard output.

Example usage:

```
go run main.go -dir=../../tests/data/ -translation=ts
```

After a successful run, the program will output a `result` folder in the current working directory with subfolders for
respective scanned packages. Filenames will be respected and maintained, with only changes to the extension.

## Building

To build this application as a binary simply navigate to `cmd/godmt` and run `go build`.

## Tips

### Tags and conversion
GoDMT tries its best to obtain the translated field name of struct fields from their struct tag. We can only choose one
tag to get its name to translate, and this happens in the following order of priority:

- `json:"field"`
- `xml:"field"`
- `form:"field"`
- `uri:"field"`
- `db:"field"`
- `mapstructure:"field"`
- `header:"field"`

If none of the above tag values are matched, then the translated name will remain as the original. As an example, a
valid tag to translate to:

```go
type ExampleType struct {
    Name         string      `json:"myFieldName"`
}
```

will be converted to:

```ts
export interface ExampleType {
  myFieldName: string
}
```

but if we would have a non-valid or non-existent tag:

```go
type ExampleType struct {
    Name         string      `whatever:"myFieldName"`
    Age int 
}
```

then it would be converted to:

```ts
export interface ExampleType {
  Name: string
  Age: number
}
```

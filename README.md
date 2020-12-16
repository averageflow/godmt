# Go Schema Converter

The goal of this repository is to provide a tool that can parse Go files that include `var`, `const`, `map`, `struct` and `type` into an abstract syntax tree, aka AST.

That AST will then be transformed into data models for several programming languages. Some small adjustments may need to be made to integrate the output into a project, but this should already save you a lot of time and hassle, and will help you stay in sync with the Go version of your data models, in other languages.

Comments will be carried over :)

Currently this tool performs conversions to TypeScript and to a custom JSON output.


## Usage

```
go run main.go -dir={scanDirectory} -translation={language} -preserve
```

- `scanDirectory` represents a string that is the relative path of the directory whose Go files you want to scan. The scan occurs in a recursive manner, so all files from all contained folders will be scanned.
- `language` represents the output mode. If the `-translation` flag is not specified it will default to JSON. Currently supported options are:
    - `ts` or `typescript` for TypeScript conversion
- `preserve` is an optional boolean flag which will make the output structs preserve the original names, instead of using the (`json:"tag"`).

Example usage:

```
go run main.go -dir=../../tests/data/ -translation=ts
```

After a successful run, the program will output a `result.*` file in the current working directory.

## Building

To build this application as a binary simply navigate to `cmd/goschemaconverter` and run `go build main.go`.
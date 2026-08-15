# Best practices for writing Go code

## Go vet

`go vet` analyzes our code for possible issues that were not caught by the Go compiler. It is a good tool for finding logical errors.

```sh
go vet projectName.go
```

## Style check with golint

`golint is` another tool for styling the Go code according to conventions mentioned in Effective Go in the official documentation of the language.

```sh
# go install golang.org/x/lint/golint@latest

golint projectName.go
```

## Formatting with gofmt and goimports

There are plenty of formatting tools, but the most popular of them are `gofmt` and `goimports`. They improve the readability of the code without affecting the execution.

```sh
gofmt projectName.go
```

`goimports` does the same formatting with one exception. It adds missed imports and removes unreferenced ones. For example, when we write code, there can be cases when we have unreferenced packages or miss some packages that we forgot to import.

```sh
# go install golang.org/x/tools/cmd/goimports@latest

goimports projectName.go
```

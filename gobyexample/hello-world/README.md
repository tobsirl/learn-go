# Hello World Example

This is a simple Go example showing a `hello` function and a test.

## Files

- `hello.go` – contains `hello()` and `main()`
- `hello_test.go` – tests the `hello()` function
- `go.mod` – module definition

## Run the program

```sh
cd gobyexample/hello-world
go run .
```

## Run tests

```sh
cd gobyexample/hello-world
go test ./...
```

## Why the earlier error happened

You saw:

```
go: cannot find main module, but found .git/config in /home/tobsirl/learn-go
```

Go 1.16+ requires a module (a `go.mod` file) for most operations outside GOPATH. Since there was no `go.mod` in the directory you were running tests from, Go looked upward, found the git repo root, and suggested initializing a module. Adding the `go.mod` in the example directory defined the module path so `go test` worked.

## Module path choice

The module path was set to `github.com/tobsirl/learn-go/gobyexample/hello-world`. If you plan to have many examples, you could instead create a single module at the repo root (`github.com/tobsirl/learn-go`) and place packages beneath it—then you’d run `go test ./...` from the root. For isolated examples, per-folder modules are fine.

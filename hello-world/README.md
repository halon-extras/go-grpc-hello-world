# hello-world

A basic hello world plugin.

## Build

```
go build -buildmode c-shared -o hello-world.so hello-world.go
```

## Test

```
hsh --config hsh.yaml --plugin hello-world test.hsl
```

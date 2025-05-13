## Golang

## Go mod

```sh
go mod init demo
```

## to run 

```sh
go run main.go
```

## to build debug build

```sh
go build -o hello main.go
```


## to build release build

```sh
 go build -ldflags="-s -w" -o small-hello main.go
```

## go install

```sh
GOBIN=/Users/jiten/workspace/training/sanas-golang-training/bin honnef.co/go/tools/cmd/staticcheck@latest         
```


## distribution list 

```sh
go tool dist list 
```

## cross compilation and build

```sh
GOOS=windows GOARCH=amd64 go build -o build/hello_windows_amd64.exe mai
n.go
```

## Compilation


```

```

## Keywords (3 of 25)

func,import,package

## Builtin functions(2 of 18)
print, println


## For Linking Issue

sudo GODEBUG=installgoroot=all go install std


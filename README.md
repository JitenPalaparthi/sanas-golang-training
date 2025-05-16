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

## Keywords (25 of 25)

break,chan,continue,case,const,defer,default,if,else,fallthrough,for,func,go,goto,import,interface,map,package,range,return,type,select,switch,struct,var 

## Builtin functions(18 of 18)
append,cap,clear,close,copy,complex,delete, imag,len, make, min, max,new, panic,print, println,real,recover

## For Linking Issue

sudo GODEBUG=installgoroot=all go install std


https://docs.google.com/presentation/d/1WVvsbvgHKBrNrKtnT4XWRfrsfkNlbw5u6L9O1DeVBn0/edit?usp=sharing
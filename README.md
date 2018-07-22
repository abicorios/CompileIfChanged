# compileIfChanged
This soft compiles if source has new changes.
# compilation
`go build compileIfChanged.go`
# usage
`compileIfChanged.exe program.go program.exe go build program.go`

where

`compileIfChanged.exe` is our tool,

`program.go` is some sourcefile,

`program.exe` is binfile compiled from program.go,

`go build program.go` is command to compile program.go.
# result
Our tool run a compile command if a sourcefile is updated after last compilation. If sourcefile is older than binfile, then our tool do nothing. It is useful for creating the optimised bat-scripts, which recompile and run developed software. For example you can made `start.bat` file with next content:

```
compileIfChanged.exe program.go program.exe go build program.go
program.exe
```

It can be useful if you use custom command line tools in development. In Golang you can use just `go run program.go`. But many other languages have not such possibility. It can be next:

```
compileIfChanged.exe program.c program.exe gcc program.c -o program.exe
program.exe
```

If you do not have gcc in Windows, you can [install Chocolatey](https://chocolatey.org/install), than run:
```
choco install mingw
refreshenv
```

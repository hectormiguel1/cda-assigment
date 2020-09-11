# cda-assigment
Computer Memory Simulator written in Go

To run program, Set up go enviroment (windows, mac os, linux) 

https://golang.org/doc/install 

Clone repo 

> git clone github.com/hectormiguel1/cda-assigment

Enter cloned directory

> cd cda-assigment 

Make sure to modify the runner.go to match your cache configuration, as well as your penaltty values. 

```go
 const (
	OperatingMode = "Write Back!"
	WriteBack = true
	OffsetLength   = 2
	IndexLength    = 3
	NumOfDataCells = 4
	NumOfIndex 	   = 8
	InCachePen     = 4
	InMemPen       = 400
)
```
This will build the program and run it, using stdin from console

> go run runner.go 

Will use test as stdin, output to console 

> go run runner.go < test 

This generates runner as an executable

> go build runner.go

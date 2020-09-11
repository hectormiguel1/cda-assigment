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
	OperatingMode = "Write Back!" //this might change in the future.
	WriteBack = true //true for using dirty bits, false for write through 
	OffsetLength   = 2 // how many bits to use for offset
	IndexLength    = 3 // how many bits to use for index 
	NumOfDataCells = 4 // number of data cells per index
	NumOfIndex 	   = 8 // number of indices present in the cache 
	InCachePen     = 4 //Penalty when Hit 
	InMemPen       = 400 //Penalty when miss 
)
```
This will build the program and run it, using stdin from console

> go run runner.go 

Will use test as stdin, output to console 

> go run runner.go < test 

This generates runner as an executable

> go build runner.go

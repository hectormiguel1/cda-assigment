package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
/* Modify this value depending on the cache configuration */
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
type Data struct {
	MemoryAddress int
}

type MemorySegment struct {
	Dirty bool
	Valid   bool
	DataTag string
	Data [NumOfDataCells]Data
}

func (m MemorySegment) String() string {
	stringForm := fmt.Sprintf("Valid: %v \t Dirty: %#v \t Tag: %s \t",
		m.Valid, m.Dirty,m.DataTag)

	for data := range m.Data {
		stringForm += fmt.Sprintf("Data %b {M[%d]},", data, m.Data[data].MemoryAddress)
	}
	return stringForm+"\n"
}
const (
	ReadOperation  = "READ"
	WriteOperation = "WRITE"
	EXIT           = "EXIT"
	OperationIndex = 0
	DataIndex      = 1
	BinaryBase = 2
	IntBase = 32
)
func main() {

	penaltyTally := 0

	//Empty memory segment (Filled with garbage, valid initialized to 0)
	Cache := make([]MemorySegment, NumOfIndex)

	//Initialize reader, to read from console/stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Program Running in %s! \n", OperatingMode)
	for {
		currentPen := 0
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			continue
		}
		splitString := strings.Fields(text)
		operation := strings.ToUpper(splitString[OperationIndex])
		if operation == EXIT {
			break
		} else if operation == ReadOperation || operation == WriteOperation {
			NumberRead, _ := strconv.Atoi(splitString[DataIndex])
			data := fmt.Sprintf("%08b", NumberRead)
			fmt.Printf("Proccesing Operation: %s, Memory: %d. %d as binary: %s \n", operation, NumberRead, NumberRead, data)
			data, offsetString := data[:len(data)-OffsetLength], data[len(data)-OffsetLength:]
			data, indexString := data[:len(data)-IndexLength], data[len(data)-IndexLength:]
			tagString := data

			fmt.Printf("Offset: %s, Index: %s, tag: %s \n", offsetString, indexString, tagString)
			index, _ := strconv.ParseInt(indexString, BinaryBase, IntBase)
			offset, err := strconv.ParseInt(offsetString, BinaryBase, IntBase)
			if err != nil {
				offset = 0
			}
			hit := FindInCache(Cache, int(index), tagString)
			if hit {
				fmt.Printf("Hit for %d in Cache! \n", NumberRead)
				currentPen += InCachePen
			} else {
				fmt.Printf("Miss for %d in Cache! \n", NumberRead)
				currentPen += InCachePen
				InsertIntoCache(Cache, int(index), int(offset), NumberRead, tagString)
				currentPen += InMemPen
			}

			if operation == WriteOperation && WriteBack {
				Cache[index].Dirty = true
			} else if operation == WriteOperation && !hit {
				currentPen += InMemPen
			}
			penaltyTally += currentPen
			fmt.Printf("Operation Penalty: %d \n", currentPen)
			PrintCache(Cache, penaltyTally)

		}
	}
}

func PrintCache(Cache []MemorySegment, cyclePen int) {
	fmt.Printf("Current Penalty Tally: %d \n", cyclePen)
	indexSizeString := fmt.Sprintf("Index: %%0%db \t",IndexLength)

	for memory := range Cache {
		tmpString := fmt.Sprintf(indexSizeString, memory) + Cache[memory].String()
		fmt.Printf(tmpString)
	}
}

func InsertIntoCache(cache []MemorySegment, index int, offset int, memAddr int, tagString string) {
	cache[index].DataTag = tagString
	startMemoryCopy := memAddr - offset
	for offsetIndex := 0; offsetIndex < NumOfDataCells; offsetIndex++ {
		cache[index].Data[offsetIndex].MemoryAddress = int(startMemoryCopy)
		startMemoryCopy++
	}
	cache[index].Valid = true
}

func FindInCache(cache []MemorySegment, index int, dataTag string) bool {
	if !cache[index].Valid {
		return false
	} else {
		if cache[index].DataTag == dataTag {
			return true
		}  else {
			return false
		}
	}
}
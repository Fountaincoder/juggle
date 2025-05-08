package main

import (
	"encoding/binary"
    "fmt"
    "os"
    "io"
    "strings"
    "path/filepath"
)

type Tuple struct {
    Delta  uint8
    Count  uint8
}

func main() {

 // Check if a filename was provided
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <filename>")
        return
    }

    // Read the filename from command line arguments
    filename := os.Args[1]
    
	file3, err := os.Open(filename)
	    if err != nil {
	        fmt.Println("Error opening file:", err)
	                return
	    }
	defer file3.Close()
	    
	var tuples []Tuple

    // Read the binary file and decode the tuples
    for {
        var tuple Tuple
        err := binary.Read(file3, binary.LittleEndian, &tuple)
        if err != nil {
            if err == io.EOF {
                break // End of file reached
            }
            fmt.Println("Error reading from file:", err)
            return
        }
        tuples = append(tuples, tuple) // Add the decoded tuple to the slice
    }  
		
	var deltas []uint8
        
	for _, tuple := range tuples {
        deltas = append(deltas, tuple.Delta)
    }
    var out_data []uint16	    

    keys2 := deltaDecode(deltas)
	fmt.Println(len(keys2))
    for i,key := range keys2 {
    	count := tuples[i].Count
    	for j :=0 ; j < int(count) ; j++{
    		out_data = append(out_data,key)
    	}
    } 
    // Create or open a file
    	newFilename := strings.TrimSuffix(filename, filepath.Ext(filename)) 
        file4, err := os.Create(newFilename)
        if err != nil {
            fmt.Println("Error creating file:", err)
            return
        }
        defer file4.Close()
    
        // Write the array to the file
        for _, value := range out_data {
            _, err := fmt.Fprintf(file4, "%d\n", value)
            if err != nil {
                fmt.Println("Error writing to file:", err)
                return
            }
    	}
    	err  = os.Remove(filename)
    	if err != nil {
        	fmt.Println("Error deleting file:", err)
        return
    	}
}

 
func deltaDecode(encoded []uint8) []uint16 {
    if len(encoded) == 0 {
        return nil
    }

    decoded := make([]uint16, len(encoded))
    decoded[0] = uint16(encoded[0]) // First element remains the same
    for i := 1; i < len(encoded); i++ {
        decoded[i] =  uint16(encoded[i]) + decoded[i-1] // Add the delta to the previous value
    }
    return decoded
}

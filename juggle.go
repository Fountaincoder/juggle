package main

import (
	"encoding/binary"
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sort"

)
type Tuple struct {
    Delta  uint8
    Count  uint32
}

func main() {

 	if len(os.Args) < 2 {
        fmt.Println("Please provide a filename as an argument.")
        return
    }

    // Get the filename from the command-line arguments
    filename := os.Args[1]
 
    // Open the file
    file, err := os.Open(filename)
    
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    //build the histogram frequency table
	frequencyMap := make(map[uint16]uint32)
    // Use a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        num, err := strconv.Atoi(line) // Convert the line to an integer

        // output error message if fails
        if err != nil {
            fmt.Println("Error converting line to integer:", err)
            continue
        }
        //read into the frequency map
        frequencyMap[uint16(num)]++
   
    }
	//keys is eventually used to produce the delta encoding
    var keys []uint16
       
	// strip out the keys 
	for key, _  := range frequencyMap {
        keys = append(keys, key)
    }

    
	sort.Slice(keys, func(i, j int) bool {
	        return keys[i] < keys[j]
	})
		
    d_enc := deltaEncode(keys)
	newFilename := filename + ".jug"
	file2, err := os.Create(newFilename)
	    if err != nil {
	        fmt.Println("Error creating file:", err)
	        return
	    }
	defer file2.Close()
	i := 0     
	for _, delta:= range d_enc {
		for frequencyMap[uint16(i)] == 0 { i++}
	    err := binary.Write(file2, binary.LittleEndian, Tuple{Delta: delta, Count: frequencyMap[uint16(i)]})
		i++
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
     
	}
}

func deltaEncode(keys []uint16) []uint8 {
    if len(keys) == 0 {
        return nil
    }

    encoded := make([]uint8, len(keys))
    encoded[0] = uint8(keys[0]) // First key remains the same
    for i := 1; i < len(keys); i++ {
        encoded[i] = uint8(keys[i] - keys[i-1]) // Store the difference
    }
    return encoded
}

package main

import (
    "fmt"
    "math/rand"
    "time"
    "math"
    "os"
    "strconv")

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a number as an argument.")
        return
    }

    // Read the argument and convert it to an integer
    arg := os.Args[1] // The first argument (os.Args[0] is the program name)
    number, err := strconv.Atoi(arg)
    if err != nil {
        fmt.Println("Invalid number:", err)
        return
    }
    // Create a file for writing
    rand.Seed(time.Now().UnixNano())
	top := int(math.Pow(2,16))
    file, err := os.Create("integers.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Write integers to the file
    for i := 1; i <= number; i++ {
        _, err := fmt.Fprintf(file, "%d\n", rand.Intn(top)) // Write each integer followed by a newline
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }
    // Seed the random number generator
}


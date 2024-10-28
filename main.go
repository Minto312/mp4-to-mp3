package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go ./main <inputfile.mp4>")
        return
    }

    inputFile := os.Args[1]
    outputFile := inputFile[:len(inputFile)-4] + ".mp3"

    cmd := exec.Command("ffmpeg", "-i", inputFile, outputFile)
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error converting file:", err)
        return
    }

    fmt.Println("Conversion successful:", outputFile)
}
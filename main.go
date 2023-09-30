package main

import (
	"fmt"
	"os"
	"strings"
)

func variance(slice []int) float64 {
	var sum float64 = 0
	for _, v := range slice {
		sum += float64(v)
	}
	mean := sum / float64(len(slice))

	var vari float64 = 0
	for _, v := range slice {
		vari += (float64(v) - mean) * (float64(v) - mean)
	}
	vari /= float64(len(slice) - 1)
	return vari
}

func main() {
	args := os.Args
	var content []byte
	var err error
	if len(args) > 1 {
		file := args[1]
		content, err = os.ReadFile(file)
	} else {
		var path string
		fmt.Print("Enter Path: ")
		fmt.Scanln(&path)
		content, err = os.ReadFile(path)
	}
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(content))
	text := string(content)

	sentences := strings.Split(text, ".")
	sentences = sentences[:len(sentences)-1]
	var lens = make([]int, len(sentences))
	// fmt.Println(sentences)

	for i, sc := range sentences {
		lens[i] = len(strings.Split(sc, " "))
	}
	// fmt.Println(lens)

	VARIANCE := variance(lens)
	fmt.Println(VARIANCE)

	var confidence float32

	if VARIANCE > 100.0 {
		confidence = 100
	} else if VARIANCE < 20.0 {
		confidence = 0
	} else {
		confidence = (float32(VARIANCE) - 20.0) * 100 / 80
	}

	fmt.Printf("\n\nI am %v%% confident that this is written by Human.\n", confidence)
	if confidence < 35.0 {
		fmt.Println("Meaning this is most likely written by AI.")
	} else if confidence > 65.0 {
		fmt.Println("Meaning this is most likely written by Human.")
	} else {
		fmt.Println("Can't really be sure. Longer Piece of text needed for more confidence")
	}

}

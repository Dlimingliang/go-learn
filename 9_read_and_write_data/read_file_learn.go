package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//inputFile, inputErrot := os.Open("9_read_and_write_data")
	//if inputErrot != nil {
	//	fmt.Printf("An error occurred on opening the inputfile\n" +
	//		"Does the file exist?\n" +
	//		"Have you got access to it?\n")
	//	return
	//}
	//
	//defer inputFile.Close()
	//
	//inputReader := bufio.NewReader(inputFile)
	//for {
	//	inputString, readerError := inputReader.ReadString('\n')
	//	fmt.Printf("The input was: %s", inputString)
	//	if readerError == io.EOF {
	//		return
	//	}
	//}

	sourceFile := "9_read_and_write_data/a.txt"
	targetFile := "9_read_and_write_data/products_copy.txt"
	buf, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(targetFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
	copyFile("9_read_and_write_data/a_copy.txt", "9_read_and_write_data/a.txt")
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

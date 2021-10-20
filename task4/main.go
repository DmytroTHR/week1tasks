package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	instruction = `Program takes 2 or 3 parameters on start: 
		<file name> <substring to count>
		<file name> <substring to replace> <replacing substring>`
)

func main() {
	params := os.Args
	switch len(params) {
	case 3:
		res, ok := countSubstringsInFile(params[1], params[2])
		if ok {
			fmt.Printf("Substring <%s> found in file %s %d times\n", params[2], params[1], res)
		} else {
			fmt.Println("There were problems finding substring in file")
		}
	case 4:
		ok := replaceSubstringInFile(params[1], params[2], params[3])
		if ok {
			fmt.Printf("Substring <%s> replaced by <%s> in file %s\n", params[2], params[3], params[1])
		} else {
			fmt.Println("There were problems replacing substring in file")
		}
	default:
		printInstructions()
	}
}

func countSubstringsInFile(fileName, substring string) (int, bool) {
	file, ok := getFile(fileName)
	if !ok {
		return 0, false
	}

	fileData, ok := getStringFromFile(file)
	if !ok {
		return 0, false
	}

	reg := regexp.MustCompile(substring)
	matchingStrings := reg.FindAllString(fileData, -1)

	return len(matchingStrings), true
}

func getStringFromFile(file FileTask) (string, bool) {
	strFile, err := file.ReadToString()
	if err != nil {
		fmt.Printf("Error reading file: \n %s\n", err.Error())
		return "", false
	}

	return strFile, true
}

func getFile(fileName string) (FileTask, bool) {
	file, ok := MakeFile(fileName)
	if !ok {
		fmt.Printf("File %s doesn't exist.\n", fileName)
		return FileTask{}, false
	}

	return file, true
}

func replaceSubstringInFile(fileName, substrFrom, substrTo string) bool {
	file, ok := getFile(fileName)
	if !ok {
		return false
	}

	fileData, ok := getStringFromFile(file)
	if !ok {
		return false
	}

	newFileData := strings.ReplaceAll(fileData, substrFrom, substrTo)
	err := file.RewriteWithString(newFileData)

	return err == nil
}

func printInstructions() {
	fmt.Println(instruction)
}

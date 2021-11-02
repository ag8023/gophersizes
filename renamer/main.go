package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	count := 0
	var toRename []string
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Dir: ", file.Name())
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}
	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFilename) // to handle different oses
		fmt.Printf("mv %s => %s\n", origPath, newPath)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}

	}
	//origPath := fmt.Sprintf("%s/%s")
	//newPath := fmt.Sprintf("%s/%s")
}

//match returns the new file name, or an error if the
//file name didnt match our pattern
func match(filename string, total int) (string, error) {
	pieces := strings.Split(filename, ".")
	//fmt.Println(pieces)
	ext := pieces[len(pieces)-1]
	//fmt.Println(ext)
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	//fmt.Println(tmp)
	pieces = strings.Split(tmp, "_")
	//fmt.Println(pieces)
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", filename)
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}

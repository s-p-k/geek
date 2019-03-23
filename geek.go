package main

import "fmt"
import "io/ioutil"
import "log"
import "os"

func geekDirExists() {
	filename := os.Getenv("HOME") + "/.geek"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File does not exist: ", filename)
		fmt.Println("Creating the directory now: ", filename)

		os.Mkdir(filename, 0740)
	}

}

func listCheatsheets() {
	cheatSheetDir := os.Getenv("HOME") + "/.geek"

	files, err := ioutil.ReadDir(cheatSheetDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func main() {
	fmt.Println("Setting geek env variable...")
	os.Setenv("GEEKHOME", "/home/spk/.geek")
	fmt.Println("GEEKHOME dir is:", os.Getenv("GEEKHOME"))
	k := os.Getenv("HOME")
	fmt.Println("HOME env var is: ", k)
	geekDirExists()
	listCheatsheets()
}

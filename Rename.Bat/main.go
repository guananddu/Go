package main

// import packages
import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// def args
var defaultOutput = strings.Join([]string{
	".",
	string(os.PathSeparator),
	"output"}, "")

var resourceDir = flag.String("re", ".", "your resources's dir")
var targetDir = flag.String("ta", defaultOutput, "your target's dir")
var replaceType = flag.Bool("r", false, "rename type: replace")
var appendType = flag.Bool("a", false, "rename type: append")
var numType = flag.Bool("n", false, "rename type: num")

// defaults
var rType bool
var aType bool
var nType bool

// define type
type fileHandler func(filepath string) error

// read file
func getFileList(path string, callback fileHandler) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		callback(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func handler() {
	getFileList(*resourceDir, func(filepath string) error {
		fmt.Printf("Handleing...: %s\n", filepath)

		splitArr := strings.Split(filepath, string(os.PathSeparator))
		filepath = splitArr[len(splitArr)-1]

		fmt.Printf("Origin Name: %s\n", filepath)

		re := regexp.MustCompile("[0-9]{1,3}")
		num := re.FindString(filepath)
		oldnum := num

		prefix := ""
		lenn := len(num)
		if lenn == 1 {
			prefix = "00"
		} else if lenn == 2 {
			prefix = "0"
		} else {
		}

		num = strings.Join([]string{prefix, num}, "")

		oldfilepath := filepath
		filepath = strings.Replace(filepath, oldnum, num, -1)

		extNameArr := strings.Split(filepath, ".")
		extName := extNameArr[len(extNameArr)-1]

		// diff name
		if aType {
			filepath = strings.Join([]string{num, filepath}, "")
		}
		if nType {
			filepath = strings.Join([]string{num, ".", extName}, "")
		}

		fmt.Printf("New Name: %s\n", filepath)

		runCmd(
			strings.Join([]string{
				*resourceDir,
				string(os.PathSeparator),
				oldfilepath}, ""),
			strings.Join([]string{
				*targetDir,
				string(os.PathSeparator),
				filepath}, ""))

		return nil
	})
}

func runCmd(re string, ta string) {
	cmd := exec.Command("cmd", "/C", "copy", re, ta)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func judgeType() {
	// set value
	if *replaceType {
		rType = *replaceType
	}

	if *appendType {
		aType = *appendType
	}

	if *numType {
		nType = *numType
	}

	// all false
	if !rType && !aType && !nType {
		rType = true
	}

	// more than 1 true
	mt1t := 0
	if rType {
		mt1t++
	}
	if aType {
		mt1t++
	}
	if nType {
		mt1t++
	}
	if mt1t > 1 {
		rType = true
		aType = false
		nType = false
	}
}

func main() {
	flag.Parse()
	judgeType()

	fmt.Print("Welcome! The Raname Exe, Auth Guananddu.\n")
	fmt.Print("-------------------------------------\n")

	fmt.Printf("Resources Dir: %s\n", *resourceDir)
	fmt.Printf("Target Dir: %s\n", *targetDir)
	fmt.Printf("Rename replaceType: %v\n", rType)
	fmt.Printf("Rename appendType: %v\n", aType)
	fmt.Printf("Rename numType: %v\n", nType)

	fmt.Print("-------------------------------------\n")
	fmt.Print("Start...\n")

	handler()

	fmt.Print("All Done,Ok!\n")
}

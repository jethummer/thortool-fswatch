package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func CheckErr(errMasg error) {
	if errMasg != nil {
		fmt.Println(errMasg)
		panic(errMasg)
	}
}

func PrintSystemInfo() {
	fmt.Print("CPU: ")
	fmt.Println(runtime.NumCPU())
	fmt.Print("GOROUTINE: ")
	fmt.Println(runtime.NumGoroutine())
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

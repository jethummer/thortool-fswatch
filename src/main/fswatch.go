package main

import (
	"container/list"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var DIR_TO_WATCH = list.New()

var watcher *fsnotify.Watcher

var dir string
var kill bool

func FsWatch() {
	process()

	show_dir_to_watch()
	create_watcher()
}

// This func parses and validates cmdline args
func process() {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Error: no such path: %s", dir)
		os.Exit(1)
	}
	DIR_TO_WATCH.PushBack(dir)
}

func show_dir_to_watch() {
	fmt.Println(" ---------------------------------------------")
	fmt.Println(" |                 fswatch                   |")
	fmt.Println(" ---------------------------------------------")
	fmt.Print(" |   Path Watch:")
	for e := DIR_TO_WATCH.Front(); e != nil; e = e.Next() {
		fmt.Print(" / [" + e.Value.(string) + "]")
	}
	fmt.Println("")
	fmt.Println(" ---------------------------------------------")

}

func create_watcher() {

	switch runtime.GOOS {
	case "linux":
	case "darwin":
		if ProcessDetectMac("fswatch") {
			fmt.Println(" |   FSWatch is running.")
			fmt.Println(" ---------------------------------------------")
		}
		break
	case "windows":
		if ProcessDetectWin("fswatch") {
			fmt.Println(" |   FSWatch is running.")
			fmt.Println(" ---------------------------------------------")
		}
	}


	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	for e := DIR_TO_WATCH.Front(); e != nil; e = e.Next() {
		if err := filepath.Walk(e.Value.(string), watchDir); err != nil {
			fmt.Println("ERROR", err)
		}
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("[MODIFY] : [", event.Name, "]")
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("[CREATE] : [", event.Name, "]")
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("[REMOVE] : [", event.Name, "]")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	<-done
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, file_info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}

	if file_info.IsDir() {
		//fmt.Printf("--- %s\n", file_info.Name())
		return watcher.Add(path)
	}

	return nil
}

func ProcessDetectWin(appName string) bool {
	if !kill {
		return true
	}
	cmd := exec.Command("wmic", "process", "get", "name,processId,CreationDate")
	output, _ := cmd.Output()

	pros := strings.Split(string(output), "\n")

	lasttime := ""
	lastpid := ""
	for _, v := range pros {
		if strings.Contains(v, appName) {
			//fmt.Println(v)
			seg := strings.Fields(v)
			//fmt.Println(len(seg))
			pid := seg[2]
			time := seg[0]
			name := seg[1]
			if strings.Contains(name, appName) {
				fmt.Println(pid,time,name)
				if lastpid == "" {
					lastpid = pid
					lasttime = time
				} else {
					if time > lasttime {
						cmd = exec.Command("taskkill","-T","-F","-pid",lastpid)
						output, _ = cmd.Output()
						lasttime = time
						lastpid = pid
					}
				}
			}
		}
	}

	return true
}

func ProcessDetectMac(appName string) bool {

	if !kill {
		return true
	}

	cmd := exec.Command("ps", "-a", "-c")
	output, _ := cmd.Output()

	pros := strings.Split(string(output), "\n")

	for _, v := range pros {
		if strings.Contains(v, appName) {
			//fmt.Println(v)
			seg := strings.Fields(v)
			//fmt.Println(len(seg))
			pid := seg[0]
			time := seg[2]
			name := seg[3]
			if appName == name && time > "0:00.01" {
				cmd = exec.Command("kill",pid)
				output, _ = cmd.Output()
				fmt.Println("out : ",string(output))
			}
		}
	}

	return true
}

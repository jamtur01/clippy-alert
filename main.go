package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	gosxnotifier "github.com/deckarep/gosx-notifier"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("myBox", "./resources")
	clippy, err := box.Find("clippyicon.png")
	if err != nil {
		log.Fatal(err)
	}

	cpath := filepath.Join(os.TempDir(), "clippyicon.png")
	err = ioutil.WriteFile(cpath, clippy, 0777)
	if err != nil {
		log.Fatal(err)
	}

	text := os.Args[1]

	alert := gosxnotifier.NewNotification(text)
	alert.Title = "Clippy"
	alert.Group = "com.unique.clippy.identifier"
	alert.AppIcon = cpath

	err = alert.Push()
	if err != nil {
		log.Println("Uh oh! Alert no good")
	}
}

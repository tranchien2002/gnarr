package main

import (
	"fmt"
	"gnarr/picker"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	arg := os.Args[1:]

	err := os.MkdirAll("output", os.FileMode(0777))
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetPrefix("gnarrr: ")

	for i := range arg {
		t := time.Now()
		log.Println(arg[i])
		legis, err := ioutil.ReadFile(arg[i])
		if err != nil {
			panic(err)
		}

		jsonout, name := picker.ToJSON(legis)
		if name == " uid" {
			name = "unidentify" + time.Since(t).String()
		}

		err = ioutil.WriteFile("output/"+name+".json", jsonout, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println(name + " exported.")

		elapsed := time.Since(t)
		fmt.Println(elapsed)
		fmt.Println()
	}

}

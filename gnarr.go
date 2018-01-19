package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gnarr/picker"
)

func main() {

	//thiết lập đầu ra cho package "log"
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	}
	
	log.SetOutput(f)
	log.SetPrefix("gnarrr: ")

	//nhận danh sách đối số
	arg := os.Args[1:]

	//tạo thư mục output
	err = os.MkdirAll("output", os.FileMode(0777))
	if err != nil {
		log.Println(err)
	}

	len := len(arg)
	//xử lý từng đối số
	for i := 0; i < len; i++ {
		t := time.Now()
		log.Println(arg[i])

		//đọc dữ liệu từ file
		legis, err := ioutil.ReadFile(arg[i])
		if err != nil {
			log.Println(err)
			continue
		}

		//chuyển dữ liệu sang dạng JSON
		jsonout, name := picker.ToJSON(legis)
		if name == " uid" {
			name = "unidentify" + time.Since(t).String()
		}

		//ghi dữ liệu JSON ra tệp tương ứng
		err = ioutil.WriteFile("output/"+name+".json", jsonout, 0644)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(name + " exported.")

		elapsed := time.Since(t)
		fmt.Println(elapsed)
		fmt.Println()
	}

}

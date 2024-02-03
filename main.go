package main

import (
	"Spider/random"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var code string

func read_website() {
	var website string
	fmt.Scanf("%s", &website)
	resp, err := http.Get(website)
	if err != nil {
		fmt.Println("", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("抱歉，读取时出了一点小错误！", err)
		return
	}
	code = string(body)
}

func save_code() {
	random.Randint()
	filePath := "code/" + random.Number + ".txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("抱歉，保存时出现了一些小错误", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(code)
	write.Flush()
}
func main() {
	read_website()
	save_code()
}

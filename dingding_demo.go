package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Start test the ding talk message sending")
	msg := "con!"
	sendDingMsg(msg)
	fmt.Println("End test the ding talk message sending")

}

func sendDingMsg(msg string) {
	webhook := `https://oapi.dingtalk.com/robot/send?access_token=ec4917bcb35f15f45ca97219ccb73fb69f5e965b95fb3e71b43ac37684f630bd`
	content := `{
		"msgtype" : "text",
		"text": {"contet": "'+ msg +'"}
	}`

	// create a request
	req, err := http.NewRequest("POST", webhook, strings.NewReader(content))
	if err != nil {
		fmt.Println("create error")

	}

	client := &http.Client{}

	//Set the request header
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	//Send the request
	resp, err := client.Do(req)

	//Close the request
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("send error")

	}

}

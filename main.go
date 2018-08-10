package main

import (
	"TelegramBot/TeleBotApi"
	"fmt"
)

// func main() {

// 	var teleurl = "https://api.telegram.org/bot691471816:AAEYREfj5h4AwXyGNDKKSAIStkoiAdrux-A/getMe"

// 	proxy := func(_ *http.Request) (*url.URL, error) {
// 		return url.Parse("http://127.0.0.1:1087")
// 	}

// 	client := &http.Client{
// 		Transport: &http.Transport{
// 			Dial: func(netw, addr string) (net.Conn, error) {
// 				conn, err := net.DialTimeout(netw, addr, time.Second*60)
// 				if err != nil {
// 					return nil, err
// 				}
// 				conn.SetDeadline(time.Now().Add(time.Second * 60))
// 				return conn, nil
// 			},
// 			ResponseHeaderTimeout: time.Second * 60,
// 			Proxy: proxy,
// 		},
// 	}

// 	resp, err := client.Get(teleurl)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	defer resp.Body.Close()

// 	fmt.Println(resp)

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	fmt.Println(string(body))
// }

func main() {
	user := TeleBotApi.NewBotApi("https://api.telegram.org/bot691471816:AAEYREfj5h4AwXyGNDKKSAIStkoiAdrux-A/getMe")
	fmt.Print(user.FirstName)
}

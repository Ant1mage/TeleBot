package TeleBotApi

import (
	. "TelegramBot/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

var proxy = func(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://127.0.0.1:1087")
}

var client = &http.Client{
	Transport: &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(netw, addr, time.Second*60)
			if err != nil {
				return nil, err
			}
			conn.SetDeadline(time.Now().Add(time.Second * 60))
			return conn, nil
		},
		ResponseHeaderTimeout: time.Second * 60,
		Proxy: proxy,
	},
}

func NewBotApi(token string) *User {

	resp, err := client.Get(token)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Print(body)

	if err != nil {
		log.Panic(err)
	}

	user := &User{}

	json.Unmarshal(body, user)

	return user

}

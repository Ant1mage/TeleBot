package teleBot

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	httpaddr = "https://api.telegram.org/bot"
	token    = "691471816:AAEYREfj5h4AwXyGNDKKSAIStkoiAdrux-A/"
)

var authorized = false

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

func NewBotApi() (User, error) {

	resp, err := MakeRequest("getme", nil)
	if err != nil {
		return User{}, err
	}
	var user User
	json.Unmarshal(resp.Result, &user)
	return user, nil

}

func MakeRequest(endpoint string, params url.Values) (APIResponse, error) {
	method := fmt.Sprintf(httpaddr + token + endpoint)
	resp, err := client.PostForm(method, params)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	var apiResp APIResponse

	bytes, err := DecodeAPIResponse(resp.Body, &apiResp)

	if err != nil {
		return apiResp, err
	}
	log.Printf("%s resp: %s", endpoint, bytes)

	if !apiResp.Ok {
		parameters := ResponseParameters{}
		if apiResp.Parameters != nil {
			parameters = *apiResp.Parameters
		}
		return apiResp, Error{apiResp.Description, parameters}
	}

	return apiResp, nil
}

func DecodeAPIResponse(responseBody io.Reader, resp *APIResponse) (_ []byte, err error) {

	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	return data, nil
}

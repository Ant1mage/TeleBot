package teleBot

import (
	"bytes"
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

func MakeTuringResult(input string, messageType int) ([]TuringResults, error) {
	res, err := MakeTuringRequest(input, messageType)
	if err != nil {
		return make([]TuringResults, 0), err
	}
	// FIXME: check status code
	return res.Results, nil
}

func MakeTuringRequest(input string, messageType int) (*TuringResponse, error) {

	var params TuringParams
	var apiResp TuringResponse

	params.UserInfo.ApiKey = turingApiKey
	params.UserInfo.UserId = turingUserId
	if messageType == 0 {
		params.Perception.InputText.Text = input
	} else if messageType == 1 {
		params.Perception.InputImage.Url = input
	}

	paramsJson, err := json.Marshal(params)
	if err != nil {
		return &apiResp, err
	}
	fmt.Println(string(paramsJson))

	request, err := http.NewRequest("POST", turingaddr, bytes.NewBuffer(paramsJson))
	if err != nil {
		return &apiResp, err
	}
	request.Header.Set("Content-Type", "application/json")

	var client = http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &apiResp, err
	}

	err = json.Unmarshal(data, &apiResp)
	if err != nil {
		return &apiResp, err
	}

	return &apiResp, err

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

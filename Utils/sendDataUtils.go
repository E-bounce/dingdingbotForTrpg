package Utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var hookUrl = Userconfig["webhook"]

func GetSendUrl()string{
	timestamp := GetTimestamp()
	sign := GetSignatureSubmit(timestamp)
	return hookUrl+"&timestamp="+strconv.Itoa(timestamp)+"&sign="+sign
}

func Send(jsonData *SendData)error{
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(jsonData)
	req,err := http.NewRequest("POST",GetSendUrl(),body)
	if err!= nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	return nil
}
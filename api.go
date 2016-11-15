package main

import (
      "net/http"
      "errors"
      "fmt"
      "log"
  
      "./signature/"
  
      // "golang.org/x/net/context/ctxhttp"
)

const (
      x_useragent = "Instagram 9.2.0 Android (18/4.3; 320dpi; 720x1280; LG; D802; armani; qcom; en_US)"
)

type InstagramClient struct {
  user string
  pass string
  isLogged bool
  http *http.Client
  useragent string
}

type ReqResponse struct {
  status int
  data string
}

func NewInstaClient(user string, pass string) *InstagramClient {
  client := &http.Client
  return &InstagramClient{user: user, pass: pass, isLogged: false, http: client, useragent: x_useragent}
}

func (i *InstagramClient) SendRequest(ispost bool, endpoint string, data string) (*ReqResponse, error) {
  if(i.isLogged == false){
    err = errors.New("Not Logged In")
    return nil, err
  }
  else {
    if(ispost == True){
      req, err := http.NewRequest("POST", endpoint, datajson)
      if(err != nil){
        return nil, err
      }
      req.Header.Add("Connection", "close")
      req.Header.Add("Accept", "*/*")
      req.Header.Add("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
      req.Header.Add("Cookie2", "$Version=1")
      req.Header.Add("Accept-Language", "en-US")
      req.Header.Add("User-Agent", classified)
      resp, err := i.http.Do(req)
      if(err != nil){
        return nil, err
      }
      if(resp.StatusCode >= 400){
        str := fmt.Sprintf("Returned %s", resp.Status)
        err := errors.New(str)
        return nil, err
      }
      return &ReqResponse{status: resp.StatusCode, data: "none"}
    }
    else{
      req, err := http.NewRequest("GET", endpoint, datajson)
      if(err != nil){
        return nil, err
      }
      req.Header.Add("Connection", "close")
      req.Header.Add("Accept", "*/*")
      req.Header.Add("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
      req.Header.Add("Cookie2", "$Version=1")
      req.Header.Add("Accept-Language", "en-US")
      req.Header.Add("User-Agent", classified)
      resp, err := i.http.Do(req)
      if(err != nil){
        return nil, err
      }
      if(resp.StatusCode >= 400){
        str := fmt.Sprintf("Returned %s", resp.Status)
        err := errors.New(str)
        return nil, err
      }      
      return &ReqResponse{status: resp.StatusCode, data: "none"}
    }
  }
}

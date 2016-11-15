package igreq

import (
    "net/http"
    "strconv"
    "log"
    "io"
)

type Agent struct {
  uagent string
}

func NewAgent(uagent string) *Agent {
  return &Agent{uagent}
}

func (a *Agent) SetUagent(str string) {
  a.uagent = str
}

func (a Agent) GetUagent() string {
  return a.uagent
}

func (a Agent) SendRequest(endpoint string, body io.Reader) (bool, *http.Response) {
  client := &http.Client{}
  if body == nil {
    r, _ := http.NewRequest("GET", endpoint, nil) 
    r.Header.Add("Connection", "close")
    r.Header.Add("Accept", "*/*")
    r.Header.Add("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
    r.Header.Add("Cookie2", "$Version=1")
    r.Header.Add("Accept-Language", "en-US")
    r.Header.Add("User-Agent", a.uagent)
    res, err := client.Do(r)
    if err != nil {
      log.Fatal(err)
      return false, nil
    }
    if res.StatusCode != 200 {
      log.Fatal("Unexpected Status Code: " + strconv.Itoa(res.StatusCode))
      return false, nil
    }
    return true, res
  } else {
    r, _ := http.NewRequest("POST", endpoint, body)
    r.Header.Add("Connection", "close")
    r.Header.Add("Accept", "*/*")
    r.Header.Add("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
    r.Header.Add("Cookie2", "$Version=1")
    r.Header.Add("Accept-Language", "en-US")
    r.Header.Add("User-Agent", a.uagent)
    res, err := client.Do(r)
    if err != nil {
      log.Fatal(err)
      return false, nil
    }
    if res.StatusCode != 200 {
      log.Fatal("Unexpected Status Code: " + strconv.Itoa(res.StatusCode))
      return false, nil
    }
    return true, res
  }
}


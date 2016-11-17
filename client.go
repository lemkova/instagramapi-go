package igapi

import (
      "./signature"
      "./igreq"
       
      u "net/url"
      "strings"
      "fmt"
      "log"
)

const ( 
    uagent = "Instagram 9.2.0 Android (18/4.3; 320dpi; 720x1280; Xiaomi; HM 1SW; armani; qcom; en_US)"
    challange = "si/fetch_headers/?challenge_type=signup&guid="
)

type InstagramClient struct {
  user string
  pass string
  csrftoken string
  isLogged bool
  deviceid string
  agent *igreq.Agent
  uuid string
}

type Stage2 struct {
  
}

func NewClient(user string, pass string) *InstagramClient {
  and_id := signature.GenerateDeviceID(user+pass)
  agent := igreq.NewAgent(uagent)
  uuid := signature.GenerateUUID(true)
  return &InstagramClient{user, pass, "", false, and_id, agent, uuid}
}

func (i *InstagramClient) Login() bool {
  if i.isLogged != true {
    constant := &signature.Constants{}
    call := fmt.Sprintf("%s%s%s", constant.GetApiEndpoint(), challange, signature.GenerateUUID(false))
    err, res := i.agent.SendRequest(call, nil)
    if err != nil {
      log.Fatal(err)
      return false
    }
    cookies := res.Cookies()[0].String()
    tok := strings.Split(cookies, ";")
    csrf := strings.Replace(tok[0], "csrftoken=", "", -1)
    i.csrftoken = csrf
    //fmt.Println(csrf) //DEBUG
    //Stage 2
    data := u.Values{}
    data.Set("data", 
    return true
  }
  return false
}
package main

import (
      "fmt"
      "./igreq"
      "./signature"
      "log"
      "io"
      "os"
)

func main(){
  constant := &signature.Constants{}
  endpoint := fmt.Sprintf("%s", constant.GetApiEndpoint())
  shelter := igreq.NewAgent("Instagram 9.2.0 Android (18/4.3; 320dpi; 720x1280; Xiaomi; HM 1SW; armani; qcom; en_US)")
  call := fmt.Sprintf("si/fetch_headers/?challenge_type=signup&guid=%s", signature.GenerateUUID(false))
  fmt.Println(endpoint+call+"\n")
  _, res := shelter.SendRequest(endpoint+call, nil)
  fmt.Printf("%d", res.StatusCode)
  defer res.Body.Close()
  _, err := io.Copy(os.Stdout, res.Body)
  if err != nil {
     log.Fatal(err)
  }
  lol := res.Cookies()
  fmt.Println("\n", lol)
}


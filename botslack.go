package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/alexamantea/go/botusers"
  
)

const BotecoChannel="C4G5WV3PF"

//func main() {
//  var SlackToken string = ""
//   var url string = fmt.Sprintf("https://slack.com/api/chat.postMessage=text", SlackToken, FutebolChannelId)
//   log.Println(url)
//  resp, err := http.Post(url, "application/json", nil)
//	log.Println("Oi")
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  var result map[string]interface{}
//  fmt.Println("Making the request")
//  json.NewDecoder(resp.Body).Decode(&result)
//  fmt.Println("Making the request2")
//  
//  fmt.Println(result)
//	var removedChannel bool = result["ok"].(bool)
//	if result {
//		log.Println("Saiu do #futebol")
//	}
//  
//}
func main() {

  UsuarioPostar, ImageURL := botusers.GetBotUserInfo(2)
  TextoPostar := os.Args[2]
  
  var SlackToken string = "xoxp-152150085732-152200984037-504478997026-6a12117f6301f8159a931cbed2e6fa42"
     var url string = fmt.Sprintf("https://slack.com/api/chat.postMessage?token=%s&channel=%s&text=%s&username=%s&icon_url=%s",
      SlackToken, BotecoChannel, TextoPostar, UsuarioPostar, ImageURL)
   log.Println(url)
  resp, err := http.Post(url, "application/json", nil)
	 if err != nil {
    log.Fatal(err)
  }
  
    var result map[string]interface{}
	fmt.Println("Making the request")
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
  
	//var postedInChannel bool = result["ok"].(bool)
	//if postedInChannel {
		//log.Println("postei no #boteco")
//	}
  
}
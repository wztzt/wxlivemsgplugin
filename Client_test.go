package wxlivemsgplugin_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/wztzt/wxlivemsgplugin"
)

func TestClient(t *testing.T) {
	cl := wxlivemsgplugin.Client{
		Cookies: "xx",
	}
	response := cl.Check_live_status()
	fmt.Println(response["errCode"])
	if response["errCode"].(float64) != 0 {
		return
	}
	data := response["data"].(map[string]interface{})
	fmt.Println(data["liveId"], data["liveObjectId"])
	fmt.Println(response)
	response = cl.Join_Live(data["liveId"].(string), data["liveObjectId"].(string))
	liveCookies := response["data"].(map[string]interface{})["liveCookies"].(string)
	for {
		response := cl.NewMsg(data["liveId"].(string), data["liveObjectId"].(string), liveCookies)
		liveCookies = response["data"].(map[string]interface{})["liveCookies"].(string)
		msgList := response["data"].(map[string]interface{})["msgList"].([]interface{})
		for _, msg := range msgList {
			fmt.Printf("%v:%v\n", msg.(map[string]interface{})["nickname"], msg.(map[string]interface{})["content"])
		}
		time.Sleep(1 * time.Second)
	}

}

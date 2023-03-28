package wxlivemsgplugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	check_live_status_url = "https://channels.weixin.qq.com/cgi-bin/mmfinderassistant-bin/live/check_live_status"
	join_live_url         = "https://channels.weixin.qq.com/cgi-bin/mmfinderassistant-bin/live/join_live"
	new_msg_url           = "https://channels.weixin.qq.com/cgi-bin/mmfinderassistant-bin/live/msg"
)

type Client struct {
	Cookies string
}

func (c *Client) Check_live_status() map[string]interface{} {
	req := CheckLiveStatusRequest{}
	data, err := json.Marshal(&req)
	if err != nil {
	}
	request, err := http.NewRequest("POST", check_live_status_url, bytes.NewReader(data))
	request.Header.Add("Cookie", c.Cookies)
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	buf := make([]byte, 2048)
	len, err := response.Body.Read(buf)
	liveStatusResponse := make(map[string]interface{})
	json.Unmarshal(buf[:len], &liveStatusResponse)
	return liveStatusResponse
}

func (c *Client) Join_Live(liveId, objectId string) map[string]interface{} {
	req := JoinLiveRequest{
		LiveId:         liveId,
		ObjectId:       objectId,
		FinderUserName: "v2_060000231003b20faec8c6e58919c4d7c70de832b077998126e081a5a1083e3f37ad167f2b34@finder",
		Scene:          7,
		ReqScene:       7,
		TimeStamp:      fmt.Sprintf("%v", time.Now().UnixMilli()),
		LogFinderId:    "v2_060000231003b20faec8c6e58919c4d7c70de832b077998126e081a5a1083e3f37ad167f2b34@finder",
	}
	data, err := json.Marshal(&req)
	if err != nil {
	}
	request, _ := http.NewRequest("POST", join_live_url, bytes.NewReader(data))
	request.Header.Add("Cookie", c.Cookies)
	request.Header.Add("Content-type", "application/json")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	buf, err := ioutil.ReadAll(response.Body)
	joinLiveResponse := make(map[string]interface{})
	json.Unmarshal(buf, &joinLiveResponse)
	return joinLiveResponse
}

func (c *Client) NewMsg(liveId, objectId, liveCookies string) map[string]interface{} {
	req := MsgRequest{
		LiveId:         liveId,
		ObjectId:       objectId,
		LiveCookies:    liveCookies,
		FinderUserName: "v2_060000231003b20faec8c6e58919c4d7c70de832b077998126e081a5a1083e3f37ad167f2b34@finder",
		Scene:          7,
		ReqScene:       7,
		TimeStamp:      fmt.Sprintf("%v", time.Now().UnixMilli()),
		LogFinderId:    "v2_060000231003b20faec8c6e58919c4d7c70de832b077998126e081a5a1083e3f37ad167f2b34@finder",
	}

	data, err := json.Marshal(&req)
	if err != nil {
	}
	request, _ := http.NewRequest("POST", new_msg_url, bytes.NewReader(data))
	request.Header.Add("Cookie", c.Cookies)
	request.Header.Add("Content-type", "application/json")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	buf, err := ioutil.ReadAll(response.Body)
	msgResponse := make(map[string]interface{})
	json.Unmarshal(buf, &msgResponse)
	return msgResponse
}

package wxlivemsgplugin

type CheckLiveStatusRequest struct {
	TimeStamp       string
	LogFinderUin    string
	LogFinderId     string
	RawKeyBuff      string
	PluginSeesionId string
	Scene           int32
	ReqScene        int32
}

type JoinLiveRequest struct {
	FinderUserName  string `json:"finderUsername"`
	LiveId          string `json:"liveId"`
	ObjectId        string `json:"objectId"`
	PluginSeesionId []byte `json:"pluginSessionId"`
	RawKeyBuff      []byte `json:"rawKeyBuff"`
	ReqScene        int32  `json:"reqScene"`
	Scene           int32  `json:"scene"`
	TimeStamp       string `json:"timestamp"`
	LogFinderId     string `json:"_log_finder_id"`
	LogFinderUin    string `json:"_log_finder_uin"`
}

type MsgRequest struct {
	FinderUserName  string `json:"finderUsername"`
	LiveCookies     string `json:"liveCookies"`
	LiveId          string `json:"liveId"`
	ObjectId        string `json:"objectId"`
	PluginSeesionId []byte `json:"pluginSessionId"`
	RawKeyBuff      []byte `json:"rawKeyBuff"`
	ReqScene        int32  `json:"reqScene"`
	Scene           int32  `json:"scene"`
	TimeStamp       string `json:"timestamp"`
	LogFinderId     string `json:"_log_finder_id"`
	LogFinderUin    string `json:"_log_finder_uin"`
}

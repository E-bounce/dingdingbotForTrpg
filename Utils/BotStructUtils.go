package Utils

type ReceiveData struct {
	Msgtype           string            `json:"msgtype"`
	Text              map[string]string `json:"text"`
	MsgId             string            `json:"msgId"`
	CreateAt          int               `json:"createAt"`
	ConversationType  string            `json:"conversationType"`
	ConversationId    string            `json:"conversationId"`
	ConversationTitle string            `json:"conversationTitle"`
	SenderId          string            `json:"senderId"`
	SenderNick        string            `json:"senderNick"`
	SenderCorpId      string            `json:"senderCorpId"`
	SenderStaffId     string            `json:"senderStaffId"`
	ChatbotUserId     string            `json:"chatbotUserId"`
	AtUsers           []map[string]string `json:"atUsers"`
}

type SendData struct {
	Msgtype string `json:"msgtype"`
	Text map[string]string `json:"text"`
}

var Userconfig = map[string]string{
	"AppKey": "xxxxx", //这里填钉钉开发者后台给出的APPsecret
	"secretkey": "xxx", // 这里填机器人界面下的密钥SE开头的
	"webhook": "xxx", //这里填机器人的webhook地址
}
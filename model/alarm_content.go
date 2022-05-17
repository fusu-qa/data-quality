package model

type AlarmContent struct {
	RuleName    string `json:"ruleName"`
	Desc        string `json:"desc"`
	Ip          string `json:"ip"`
	DbName      string `json:"dbName"`
	TbName      string `json:"tbName"`
	ColName     string `json:"colName"`
	ExpectValue string `json:"expectValue"`
	ActualValue string `json:"actualValue"`
}

func NewAlarmContent(ruleName string, desc string, ip string, dbName string, tbName string, colName string, expectValue string, actualValue string) *AlarmContent {
	return &AlarmContent{RuleName: ruleName, Desc: desc, Ip: ip, DbName: dbName, TbName: tbName, ColName: colName, ExpectValue: expectValue, ActualValue: actualValue}
}

type FeisuMsgBody struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}

func NewFeisuMsgBody(msgType string, content Content) *FeisuMsgBody {
	return &FeisuMsgBody{MsgType: msgType, Content: content}
}

type Content struct {
	Text string `json:"text"`
}

func NewContent(text string) *Content {
	return &Content{Text: text}
}

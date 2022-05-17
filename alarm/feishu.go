package alarm

import (
	"data-quality/model"
	"data-quality/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func ToFeisu(content *model.AlarmContent) {
	url := "https://open.feishu.cn/open-apis/bot/v2/hook/22ecb732-4fc3-4516-8856-c11126a"
	contentType := "application/json"
	builder := strings.Builder{}
	builder.WriteString("规则名：" + content.RuleName + "\n")
	builder.WriteString("规则描述：" + content.Desc + "\n")
	builder.WriteString("库IP：" + content.Ip + "\n")
	builder.WriteString("库名：" + content.DbName + "\n")
	builder.WriteString("表名：" + content.TbName + "\n")
	builder.WriteString("字段：" + content.ColName + "\n")
	builder.WriteString("期望值：" + content.ExpectValue + "\n")
	builder.WriteString("实际值：" + content.ActualValue + "\n")
	msgBody := model.NewFeisuMsgBody("text", *model.NewContent(builder.String()))
	bytes, _ := json.Marshal(msgBody)
	str := util.Byte2Str(bytes)
	resp, err := http.Post(url, contentType, strings.NewReader(str))
	if err != nil {
		fmt.Printf("post failed, err:%v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("close failed, err:%v", err)
		}
	}(resp.Body)
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

}

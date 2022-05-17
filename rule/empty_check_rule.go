package rule

import (
	"data-quality/alarm"
	"data-quality/db"
	"data-quality/model"
	"fmt"
	"strconv"
)

type EmptyValueChecker struct {
}

func NewEmptyValueChecker() *EmptyValueChecker {
	return &EmptyValueChecker{}
}

func (checker EmptyValueChecker) Check() {
	ip := "10.70.104.17"
	port := "9000"
	dbName := "lcc"
	username := "dsfd"
	password := "dfdfd"
	dsn := fmt.Sprintf("tcp://%s:%s?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20", ip, port, dbName, username, password)
	ch, _ := db.InitCH(dsn)
	var count int64
	userNodeMetric := model.UserNodeMetric{}
	ch.Debug().Model(&userNodeMetric).
		Where("cluster_name=?", "lcc").
		Where("node_label=?", "").Count(&count)
	if count > 0 {
		actualValue := strconv.Itoa(int(count))
		expectValue := "0"
		content := model.NewAlarmContent(checker.Name(), checker.Desc(), ip, dbName, userNodeMetric.TableName(), "node_label", expectValue, actualValue)
		alarm.ToFeisu(content)
	}
}

func (checker EmptyValueChecker) Name() string {
	return "EmptyValueCheck"
}

func (checker EmptyValueChecker) Desc() string {
	return "空值检查"
}

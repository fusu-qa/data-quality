package model

import "time"

type UserNodeMetric struct {
	ClusterID    int64     `json:"clusterID"`
	ClusterName  string    `json:"clusterName"`
	NodeName     string    `json:"nodeName"`
	NodeLabel    string    `json:"nodeLabel"`
	NodeIp       string    `json:"nodeIp"`
	CoreTotal    int64     `json:"coreTotal"`
	CoreAvi      int64     `json:"coreAvi"`
	CoreAlt      int64     `json:"coreAlt"`
	CpuUsage     float64   `json:"cpuUsage"`
	MemTotal     int64     `json:"memTotal"`
	MemAvi       int64     `json:"memAvi"`
	MemAlt       int64     `json:"memAlt"`
	MemUsageByte int64     `json:"memUsageByte"`
	MemUsage     float64   `json:"memUsage"`
	DiskTotal    int64     `json:"diskTotal"`
	DiskAvi      int64     `json:"diskAvi"`
	DiskAlt      int64     `json:"diskAlt"`
	DiskUsage    float64   `json:"diskUsage"`
	RunningNum   int32     `json:"runningNum"`
	CollectedAt  int64     `json:"collectedAt"`
	Partition    string    `json:"partition"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (UserNodeMetric) TableName() string {
	return "user_node_metric"
}

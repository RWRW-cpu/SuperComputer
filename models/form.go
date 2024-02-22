package models

type IteaqLengTongDaoYanGan struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

type IteaqAirConditioner1s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

type IteaqAirConditioner2s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

type IteaqAirConditioner3s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqAirConditioner4s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqAirConditioner5s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqAirConditioner6s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

type IteaqJiDianQiZhuZhuang struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

type IteaqMenJinJiFang struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

type IteaqMenJinWeiMoKuai struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

type IteaqOtherDianChi struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqPeiDianJingMi struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqPeiDianPd1s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqPeiDianUps1s struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqTemp struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}
type IteaqUps struct {
	Id           int    `json:"id"`
	SignalName   string `json:"signal_name"`
	Value        string `json:"value"`
	SamplingTime int64  `json:"sampling_time"`
}

//硬盘
type GatewayDiskUsage struct {
	Id       int    `json:"id"`
	Time     int64  `json:"time"`
	CpuUsage string `json:"cpu_usage"`
}

type GatewayMemoryUsage struct {
	Id          int    `json:"id"`
	Time        int64  `json:"time"`
	MemoryUsage string `json:"memory_usage"`
}

type GatewayCpuUsage struct {
	Id        int    `json:"id"`
	Time      int64  `json:"time"`
	DiskUsage string `json:"disk_usage"`
}

type GatewaySession struct {
	Id           int    `json:"id"`
	Time         int64  `json:"time"`
	SessionCount string `json:"session_count"`
	SessionNew   string `json:"session_new"`
}
package svc

import (
	"os"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/zeromicro/go-zero/core/logx"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		logx.Errorf("初始化 sentinel 异常: %v", err)
		os.Exit(1)
	}

	//配置限流规则
	//这种配置应该从nacos中读取
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "test",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, // 拒绝请求
			Threshold:              3,           // 表示6秒内最多允许3个请求
			StatIntervalInMs:       6000,
		},
	})

	if err != nil {
		logx.Errorf("sentinel 加载规则失败: %v", err)
		os.Exit(1)
	}
}

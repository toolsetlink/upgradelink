package crontab

type TaskConfig struct {
	Name      string `json:"name"`      //任务名称
	Scheduler string `json:"scheduler"` //@every表达式中的单位可以是s（秒）、m（分钟）、h（小时）等
	Enabled   bool   // 是否启用
}

func DefaultTask() []TaskConfig {
	return []TaskConfig{
		{
			Name:      "SayHello",
			Scheduler: "@every 30m", // 每30分钟执行一次
			Enabled:   false,
		},
		{
			Name:      "ApkPatch",
			Scheduler: "@every 30m", // 每30分钟执行一次
			Enabled:   true,
		},
	}
}

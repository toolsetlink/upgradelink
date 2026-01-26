package crontab

import (
	"context"
	"fmt"
	"log"
	"upgradelink-admin/server/api/internal/crontab/task/apk_patch"
	"upgradelink-admin/server/api/internal/crontab/task/hello"
	"upgradelink-admin/server/api/internal/svc"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron   *cron.Cron
	svcCtx *svc.ServiceContext
}

// NewScheduler creates a new scheduler instance.
func NewScheduler(svcCtx *svc.ServiceContext) *Scheduler {
	return &Scheduler{
		cron:   cron.New(), // 支持秒级精度 （可选）
		svcCtx: svcCtx,
	}
}

// RegisterTask registers a api with the scheduler.
func (s *Scheduler) RegisterTask(name, schedule string, task func(ctx context.Context, svcCtx *svc.ServiceContext)) {
	_, err := s.cron.AddFunc(schedule, func() {
		task(context.Background(), s.svcCtx)
	})
	if err != nil {
		log.Fatalf("Failed to register task %s: %v", name, err)
	}
	log.Printf("Registered task %s with schedule %s", name, schedule)
}

func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("Scheduler started")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Println("Scheduler stopped")
}

// InitScheduler 初始化定时任务调度器
func InitScheduler(svcCtx *svc.ServiceContext) (*Scheduler, error) {
	log.Println("Initializing scheduler...")
	scheduler := NewScheduler(svcCtx)
	if scheduler == nil {
		return nil, fmt.Errorf("failed to create scheduler")
	}
	taskConfigs := DefaultTask()
	if taskConfigs == nil {
		return nil, fmt.Errorf("failed to load api configurations")
	}
	for _, taskConfig := range taskConfigs {
		if !taskConfig.Enabled {
			log.Printf("Task %s is disabled, skipping...", taskConfig.Name)
			continue
		}

		log.Printf("Registering api %s...", taskConfig.Name)

		switch taskConfig.Name {
		case "SayHello":
			scheduler.RegisterTask(taskConfig.Name, taskConfig.Scheduler, hello.SayHello)
		case "ApkPatch":
			scheduler.RegisterTask(taskConfig.Name, taskConfig.Scheduler, apk_patch.ApKPatch)
		case "xxxx":
			// 注册其他定时任务
		default:
			log.Printf("Unknown task: %s", taskConfig.Name)
		}
	}
	return scheduler, nil
}

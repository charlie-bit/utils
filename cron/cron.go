package cron

import (
	"fmt"
	"log"
	"time"
	"github.com/charlie-bit/utils/config"

	"github.com/robfig/cron/v3"
)

type _Jobs interface {
	Run()
	SetLockExpire(duration time.Duration)
	Unlock() error
}

var jobs = map[string]_Jobs{
	"hello_world": &HelloWorld{},
}

func NewTasks(tasks []config.Task) (*cron.Cron, error) {
	// If a previous invocation is still running. It logs skips to the given logger at Info level.
	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
	for _, task := range tasks {
		job, ok := jobs[task.Name]
		if ok {
			// set expiration
			expiration, err := time.ParseDuration(task.Expiration)
			if err != nil {
				return nil, fmt.Errorf("init task ParseDuration expiration failed  {err: %v}", err.Error())
			}
			job.SetLockExpire(expiration)

			// add job
			_, err = c.AddJob(task.Timer, job)
			if err != nil {
				return nil, fmt.Errorf("init task AddJob failed  {err: %v}", err.Error())
			}
			// Tasks that need to be performed first
			if task.FirstRun {
				if err := job.Unlock(); err == nil {
					go job.Run()
				}
			}
		}
	}

	c.Start()
	log.Println("init tasks success...")
	return c, nil
}

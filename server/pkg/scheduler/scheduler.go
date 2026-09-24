package scheduler

import (
	"log"
	"time"
)

type Job struct {
	Name     string
	Callback func() error
	Interval time.Duration
	Stop     chan bool
}

type Scheduler struct {
	jobs map[string]Job
}

func (instance *Scheduler) AddJob(name string, callback func() error, interval time.Duration) {
	stop := instance.launch(callback, interval)
	instance.jobs[name] = Job{
		Name:     name,
		Callback: callback,
		Interval: interval,
		Stop:     stop,
	}
}

func (instance *Scheduler) launch(callback func() error, interval time.Duration) chan bool {
	ticker := time.NewTicker(interval)
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:

				err := callback()
				if err != nil {
					log.Println(err)
				}
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}

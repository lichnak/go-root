package main

import (
	"github.com/rk/go-cron"
	"time"
)

func task1(t time.Time) {
	println("task1:", t.String())
}

func task2(t time.Time) {
	println("task2:", t.String())
}

func task3(t time.Time) {
	println("task3:", t.String())
}

func main() {
	c := make(chan bool)

	cron.NewDailyJob(cron.ANY, cron.ANY, 0, task1)
	cron.NewDailyJob(cron.ANY, cron.ANY, 10, task2)
	cron.NewDailyJob(cron.ANY, cron.ANY, cron.ANY, task3)

	<-c
}

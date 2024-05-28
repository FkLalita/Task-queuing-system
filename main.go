package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/FkLalita/task-queuing/task"

	"github.com/FkLalita/task-queuing/worker"
)

func main() {
	ctx := context.Background()
	fmt.Println("Go Redis")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	tasks := &task.Task{
		ID:   "33",
		Type: "send-email",
		Data: map[string]interface{}{"to": "user@example.com", "subject": "Important message"},
	}
	task.SubmitTask(ctx, *rdb, tasks)

	//just trying to check if it works, do not do it this way
	defer worker.ProcessTasks(ctx, *rdb)

}

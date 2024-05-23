package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Task struct {
	ID   string
	Type string
	Data map[string]interface{}
}

func SubmitTask(ctx context.Context, rdb redis.Client, task *Task) {
	taskByte, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(taskByte)

	_, err = rdb.LPush(ctx, "task", taskByte).Result()
	if err != nil {
		fmt.Println(err)
	}

}

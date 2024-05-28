// Package task ...
package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Task struct
type Task struct {
	ID   string                 `json:"id"`
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// SubmitTask take the task struct and Marshal to json and Push to reddit
func SubmitTask(ctx context.Context, rdb redis.Client, task *Task) {
	taskByte, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(taskByte), "annn")

	_, err = rdb.LPush(ctx, "tasks", taskByte).Result()
	if err != nil {
		fmt.Println(err)
	}

}

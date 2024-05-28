package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FkLalita/task-queuing/task"
	"github.com/redis/go-redis/v9"
)

// worker function that continously poll the redis server for tasks
// it Unmarshal the tasks if found

func ProcessTasks(ctx context.Context, client redis.Client) {
	for {
		fmt.Println("checking///")
		taskBytes, err := client.BRPop(ctx, 1000000000, "tasks").Result()
		if err != nil {
			if err == redis.Nil { // wait for task to pop up
				fmt.Println("no task yet")
				continue
			}
			fmt.Println(err)

		}
		var task task.Task
		err = json.Unmarshal([]byte(taskBytes[1]), &task) // Ignore first element
		if err != nil {
			fmt.Println(err)
			break
		}

		//process task herer
		fmt.Println("Task--  ", task)
	}
}

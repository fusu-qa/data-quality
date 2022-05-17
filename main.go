package main

import (
	"data-quality/rule"
	"data-quality/task"
)

func main() {
	task.Run(rule.NewEmptyValueChecker().Check)
}

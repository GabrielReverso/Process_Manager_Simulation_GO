package prioritylist

import (
	taskcreate "module/taskCreate"
	"sort"
)

var taskChanel = make(chan []taskcreate.TaskStruct)

type Task []taskcreate.TaskStruct

func (t Task) Len() int {
	return len(t)
}

func (t Task) Less(i, j int) bool {
	return t[i].Priority < t[j].Priority
}

func (t Task) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func Prioritylist(tasks []taskcreate.TaskStruct) {

	count := 0
	tempo := 0

	for {
		allCompleted := true
		for _, task := range tasks {
			if task.QuantumI != 0 {
				allCompleted = false
			}

			for quantuns := 0; quantuns < 10; quantuns++ {
				if task.QuantumI == 0 {
					break
				}
				task.QuantumI--
				tempo++
			}

		}

		count++

		if count == 3 {
			count = 0
			go RoutineProcessVector(taskChanel, tasks)

			tasks = <-taskChanel
		}

		if allCompleted {
			break
		}

	}

}

func RoutineProcessVector(channel chan []taskcreate.TaskStruct, TaskI []taskcreate.TaskStruct) {

	t := Task(TaskI)

	sort.Sort(t)

	TaskI = []taskcreate.TaskStruct(t)

	channel <- TaskI

}

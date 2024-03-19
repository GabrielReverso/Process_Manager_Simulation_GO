package prioritylist

import (
	//"fmt"
	taskcreate "module/taskCreate"
)

type Task []taskcreate.TaskStruct

func (t Task) len() int {
	return len(t)
}

func (t Task) Less(i, j int) bool {
	return t[i].Priority < t[j].Priority
}

func (t Task) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func Prioritylist() {

}

func RoutineProcessVector() {

	//Canal de comunicação
	//Datachanel := make(chan []taskcreate.TaskStruct)

	go func() {

		Tasks := taskcreate.TaskVetorCreator(20)

		for {
			TamanoVetor := len(Tasks)

			for i := 0; i < TamanoVetor; i++ {

			}
		}

	}()

}

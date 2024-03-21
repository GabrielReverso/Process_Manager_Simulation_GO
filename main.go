package main

import (
	"fmt"
	roundrobin "module/roundRobin"
	taskcreate "module/taskCreate"
	priorityList "module/priorityList"
)

func main() {

	qtdTask := 1
	fmt.Println("Digite o numero de processos (1 - 50): ")
	fmt.Scanf("%d", &qtdTask)

	var tasks []taskcreate.TaskStruct

	if qtdTask < 1 || qtdTask > 50 {
		fmt.Println("Numero invalido!")
		return
	} else {
		tasks = taskcreate.TaskVetorCreator(qtdTask)
	}

	fmt.Println()

	managerOption := 1

	fmt.Println("Digite o numero do algoritimo de teste: ")
	fmt.Println("1 -> Round Robin")
	fmt.Scanf("%d", &managerOption)

	fmt.Print("\033[H\033[2J")

	switch managerOption {
	case 1:
		{
			roundrobin.RoundRobin(tasks)
		}
	case 2:
		{
			priorityList.Prioritylist(tasks)
		}
	default:
		{
			roundrobin.RoundRobin(tasks)
		}
	}
}

package main

import (
	"fmt"
	firstInfirstOut "module/firstInFirstOut"
	"module/priorityFirst"
	priorityList "module/priorityList"
	roundrobin "module/roundRobin"
	"module/roundRobinPriority"
	"module/shortestFirst"
	taskcreate "module/taskCreate"
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
	fmt.Println("2 -> Round Robin - Priority")
	fmt.Println("3 -> Shortest First")
	fmt.Println("4 -> Priority First")
	fmt.Println("5 -> First In First Out")
	fmt.Println("5 -> Priority List")
	fmt.Println("")
	fmt.Scanf("%d", &managerOption)

	fmt.Print("\033[H\033[2J")

	switch managerOption {
	case 1:
		{
			roundrobin.RoundRobin(tasks)
		}
	case 2:
		{
			roundRobinPriority.RoundRobinPriority(tasks)
		}
	case 3:
		{
			shortestFirst.ShortestFirst(tasks)
		}
	case 4:
		{
			priorityFirst.PriorityFirst(tasks)
		}
	case 5:
		{
			firstInfirstOut.FirstInfirstOut(tasks)
		}
	case 6:
		{
			priorityList.Prioritylist(tasks)
		}
	default:
		{
			fmt.Println("Nenhum selecionado!")
			return
		}
	}
}

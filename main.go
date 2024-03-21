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
	"time"
)

func main() {

	qtdTask := 1
	fmt.Println("Digite o numero de processos (1 - 50): ")
	fmt.Scanf("%d", &qtdTask)

	testType := 0
	fmt.Println("")
	fmt.Println("Digite o tipo de teste: ")
	fmt.Println("1 -> Tasks curtas")
	fmt.Println("2 -> Tasks longas")
	fmt.Println("3 -> Tasks variadas")
	fmt.Println("4 -> Tasks aleatorias (Default)")
	fmt.Println("")
	fmt.Scanf("%d", &testType)

	var tasks []taskcreate.TaskStruct

	if qtdTask < 1 || qtdTask > 50 {
		fmt.Println("Numero invalido!")
		return
	} else {
		tasks = taskcreate.TaskVetorCreator(qtdTask, testType)
	}

	fmt.Println()

	managerOption := 1

	fmt.Println("Digite o numero do algoritimo de teste: ")
	fmt.Println("1 -> Round Robin")
	fmt.Println("2 -> Round Robin - Priority")
	fmt.Println("3 -> Round Robin - Dynamic")
	fmt.Println("4 -> Shortest First")
	fmt.Println("5 -> Priority First")
	fmt.Println("6 -> First In First Out")
	fmt.Println("")
	fmt.Scanf("%d", &managerOption)

	fmt.Print("\033[H\033[2J")

	start := time.Now()

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
			priorityList.Prioritylist(tasks)

		}
	case 4:
		{
			shortestFirst.ShortestFirst(tasks)

		}
	case 5:
		{
			priorityFirst.PriorityFirst(tasks)

		}
	case 6:
		{
			firstInfirstOut.FirstInfirstOut(tasks)

		}
	default:
		{
			fmt.Println("Nenhum selecionado!")
			return
		}
	}

	end := time.Now()
	duration := end.Sub(start)
	durationInSeconds := duration.Seconds()

	fmt.Printf("\nTempo de execucao: %.2f segundos\nQuantidade de processos: %d\nProcessos concluidos por segundo: %.2f\n\n", durationInSeconds, qtdTask, float64(qtdTask)/durationInSeconds)
}

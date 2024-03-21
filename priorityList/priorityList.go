package prioritylist

import (
	"fmt"
	taskcreate "module/taskCreate"
	"sort"

	//"runtime/trace"
	"sync"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)


type taskAux struct {
	task taskcreate.TaskStruct
	index int
}

var taskChanel = make(chan []taskAux)


func Prioritylist(tasks []taskcreate.TaskStruct) {
	const (
		Reset = "\033[0m"
		Red   = "\033[31m"
		/* 		Green   = "\033[32m"
		   		Yellow  = "\033[33m"
		   		Blue    = "\033[34m"
		   		Magenta = "\033[35m"
		   		Cyan    = "\033[36m"
		   		White   = "\033[37m" */
	)


	fmt.Println(Red + "    ID   " + Reset + "|" + Red + "   PRIORITY   " + Reset + "|" + Red + "       QUANTUM       " + Reset + "|" + Red + "                                               PROGRESS                                         " + Reset)

	// Cria um progresso principal
	p := mpb.New(mpb.WithWaitGroup(&sync.WaitGroup{}))

	// Criar as barras de progresso dinamicamente
	bars := make([]*mpb.Bar, len(tasks))
	for i, task := range tasks {
		bars[i] = p.AddBar(int64(task.Quantum), mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("PID: %3.d | Priority: %2.d |", task.Id, task.Priority)),
			decor.CountersNoUnit(" Quantum: %3.d / %3.d |", decor.WCSyncSpace),
		), mpb.AppendDecorators(
			decor.Percentage(decor.WCSyncSpace),
			decor.Elapsed(decor.ET_STYLE_MMSS, decor.WCSyncSpace),
			decor.OnComplete(decor.Name("          "), " Completed"),
		))
		time.Sleep(time.Millisecond * 10)
	}

	TasksComplete := make([]taskAux, len(tasks))

	for i, t := range tasks {
		TasksComplete[i].task = t
		TasksComplete[i].index = i	
	}

	count := 1

	for {
		allCompleted := true

		for _, bar := range bars {
			if bar.Completed(){
				continue
			}
			if !bar.Completed(){
				allCompleted = false
			}
		}

		

		for _, t := range TasksComplete {
			if t.task.Quantum == 0{
				continue
			}
			for quantuns := 0; quantuns < 10 + int((10/t.task.Priority)); quantuns++ {
				t.task.QuantumI--
				bars[t.index].IncrBy(1)
				time.Sleep(time.Millisecond * 10)				
			}
			time.Sleep(time.Millisecond * 100)
			
		}

		count++

		if count >= 3{
			count = 1
			go RoutineProcessVector (taskChanel, TasksComplete)

			newTask := <- taskChanel

			TasksComplete = newTask
		}



		if allCompleted {
			break
		}
	}


	p.Wait()
}

func RoutineProcessVector(channel chan []taskAux, TaskI []taskAux) {

	for _, t := range TaskI {
		conta := float64(t.task.QuantumI) /  float64(t.task.Quantum)
		if conta >= 0.95{
			t.task.Priority = t.task.Priority * 8
		}else if conta < 0.95 && conta	>= 0.80{
			t.task.Priority = t.task.Priority * 6
		}else if conta < 0.80 && conta	>= 0.50{
			t.task.Priority = t.task.Priority * 4
		}

		

		t.task.Quantum = t.task.QuantumI
	}

	sort.Slice(TaskI, func(i, j int) bool {
		return TaskI[i].task.Priority > TaskI[j].task.Priority
	})

	fmt.Println("HI")

	
	channel <- TaskI
}




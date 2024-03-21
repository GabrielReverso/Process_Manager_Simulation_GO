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
	task      taskcreate.TaskStruct
	index     int
	interacao int
}

func Prioritylist(tasks []taskcreate.TaskStruct) {
	const (
		Reset = "\033[0m"
		Red   = "\033[31m"
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
		TasksComplete[i].interacao = 0
	}

	for {
		allComplete := true

		for _, t := range TasksComplete {
			if !bars[t.index].Completed() {
				allComplete = false
				for quantuns := 0; quantuns < 10+(100/t.task.Priority); quantuns++ {
					if bars[t.index].Completed() {
						break
					}
					t.task.QuantumI = t.task.QuantumI - 1
					bars[t.index].IncrBy(1)
					time.Sleep(time.Millisecond * 10)
					t.interacao++
				}
				time.Sleep(time.Millisecond * 100)
			}
		}

		if allComplete {
			break
		}

		newTaskComplete := RoutineProcessVector(TasksComplete)

		TasksComplete = newTaskComplete

	}

	p.Wait()
}

func RoutineProcessVector(TaskI []taskAux) []taskAux {

	newTaskI := TaskI

	for _, t := range newTaskI {

		multiply := len(newTaskI)

		if t.interacao >= 40*multiply {
			t.task.Priority = (t.task.Priority / 20) * 20
		} else if t.interacao < 40*multiply && t.interacao >= 30*multiply {
			t.task.Priority = (t.task.Priority / 10) * 10
		} else if t.interacao < 30*multiply && t.interacao > 20*multiply {
			t.task.Priority = (t.task.Priority / 5) * 10
		} else if t.interacao < 20*multiply && t.interacao > 10*multiply {
			t.task.Priority = (t.task.Priority / 2) * 10
		}

	}

	sort.Slice(newTaskI, func(i, j int) bool {
		return newTaskI[i].task.Priority > newTaskI[j].task.Priority
	})

	return newTaskI
}

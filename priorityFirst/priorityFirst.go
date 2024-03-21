package priorityFirst

import (
	"fmt"
	"sort"

	"sync"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	taskcreate "module/taskCreate"
)

func PriorityFirst(tasks []taskcreate.TaskStruct) {
	const (
		Reset = "\033[0m"
		Red   = "\033[31m"
	)

	priorityTaskIdx := make([]int, len(tasks))
	for i := range priorityTaskIdx {
		priorityTaskIdx[i] = i
	}

	sort.SliceStable(priorityTaskIdx, func(i, j int) bool {
		return tasks[priorityTaskIdx[i]].Priority < tasks[priorityTaskIdx[j]].Priority
	})

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

	// Atualiza as barras de progresso dinamicamente
	for _, index := range priorityTaskIdx {
		for !bars[index].Completed() {
			bars[index].IncrBy(1)
			time.Sleep(time.Millisecond * 10)
		}
		time.Sleep(time.Millisecond * 100)
	}

	// Finaliza o progresso principal
	p.Wait()

}

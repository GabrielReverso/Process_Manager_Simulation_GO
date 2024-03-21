package roundRobinPriority

import (
	"fmt"
	"sync"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	taskcreate "module/taskCreate"
)

func RoundRobinPriority(tasks []taskcreate.TaskStruct) {
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

	// Atualiza as barras de progresso dinamicamente
	// Continua até que todas as barras de progresso estejam concluídas.
	for {
		allCompleted := true
		// Itera sobre cada barra de progresso no slice 'bars'.
		for i, bar := range bars {
			// Verifica se a barra de progresso atual não está concluída.
			if !bar.Completed() {
				// Se a barra de progresso não estiver concluída, define 'allCompleted' como false.
				allCompleted = false
				for quantum := 0; quantum < 10+(10/tasks[i].Priority); quantum++ {
					bar.IncrBy(1)
					time.Sleep(time.Millisecond * 10)
					if bar.Completed() {
						break
					}
				}
				// Após incrementar a barra de progresso 10 vezes, aguarda 100 milissegundos.
				time.Sleep(time.Millisecond * 100)
			}
		}
		// Se todas as barras de progresso estiverem concluídas sai do loop infinito.
		if allCompleted {
			break
		}
	}

	// Finaliza o progresso principal
	p.Wait()
}

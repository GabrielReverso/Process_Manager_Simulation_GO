package main

import (
	"fmt"
	roundrobin "module/roundRobin"
)

/* 	"fmt"
"sync"
"time"

"github.com/vbauerster/mpb/v7"
"github.com/vbauerster/mpb/v7/decor" */

func main() {
	/* 	limits := []int64{600, 200, 200, 400, 500}

	   	// Cria um progresso principal
	   	p := mpb.New(mpb.WithWaitGroup(&sync.WaitGroup{}))

	   	// Criar as barras de progresso dinamicamente
	   	bars := make([]*mpb.Bar, len(limits))
	   	for i, limit := range limits {
	   		bars[i] = p.AddBar(limit, mpb.PrependDecorators(
	   			decor.Name(fmt.Sprintf("Process %d", i+1)),
	   			decor.CountersNoUnit("%d / %d", decor.WCSyncSpace),
	   		), mpb.AppendDecorators(
	   			decor.Percentage(decor.WCSyncSpace),
	   			decor.OnComplete(decor.Name("          "), " Completed"),
	   			decor.Elapsed(decor.ET_STYLE_MMSS, decor.WCSyncSpace),
	   		))
	   	}

	   	// Atualiza as barras de progresso dinamicamente
	   	// Continua até que todas as barras de progresso estejam concluídas.
	   	for {
	   		allCompleted := true
	   		// Itera sobre cada barra de progresso no slice 'bars'.
	   		for _, bar := range bars {
	   			// Verifica se a barra de progresso atual não está concluída.
	   			if !bar.Completed() {
	   				// Se a barra de progresso não estiver concluída, define 'allCompleted' como false.
	   				allCompleted = false
	   				for quantum := 0; quantum < 10; quantum++ {
	   					bar.IncrBy(1)
	   					time.Sleep(time.Millisecond * 10)
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
	   	p.Wait() */

	//TO DO
	numTask := 1
	fmt.Println("Digite o numero de processos (1 - 100): ")
	fmt.Scanf("%d", numTask)

	if numTask < 1 || numTask > 100 {
		fmt.Println("Numero invalido!")
		return
	}

	fmt.Println("Digite o numero do algoritimo de teste: ")
	fmt.Println("1 -> Round Robin")

	roundrobin.RoundRobin()
}

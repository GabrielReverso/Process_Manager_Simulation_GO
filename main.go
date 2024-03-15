package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"sync"
	"time"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

func main() {
	var limit int64 = 300 // 1MB

	// Cria dois arquivos temporários para simular dois downloads
	tempFile1, err := ioutil.TempFile("", "process1")
	if err != nil {
		panic(err)
	}
	defer tempFile1.Close() // Garante que o arquivo vai ser fechado ao terminar o programa

	tempFile2, err := ioutil.TempFile("", "process2")
	if err != nil {
		panic(err)
	}
	defer tempFile2.Close()

	tempFile3, err := ioutil.TempFile("", "process3")
	if err != nil {
		panic(err)
	}
	defer tempFile3.Close()

	// Cria um progresso principal
	p := mpb.New(mpb.WithWaitGroup(&sync.WaitGroup{}))

	// Cria duas barras de progresso
	bar1 := p.AddBar(limit, mpb.PrependDecorators(
		decor.Name("Process 1"),
		decor.CountersNoUnit("%d / %d", decor.WCSyncSpace),
	), mpb.AppendDecorators(
		decor.Percentage(decor.WCSyncSpace),
		decor.OnComplete(decor.Name("          "), " Completed"),
	))

	bar2 := p.AddBar(100, mpb.PrependDecorators(
		decor.Name("Process 2"),
		decor.CountersNoUnit("%d / %d", decor.WCSyncSpace),
	),
		mpb.AppendDecorators(
			decor.Percentage(decor.WCSyncSpace),
			decor.OnComplete(decor.Name("          "), " Completed"),
		))

	bar3 := p.AddBar(200, mpb.PrependDecorators(
		decor.Name("Process 3"),
		decor.CountersNoUnit("%d / %d", decor.WCSyncSpace),
	),
		mpb.AppendDecorators(
			decor.Percentage(decor.WCSyncSpace),
			decor.OnComplete(decor.Name("          "), " Completed"),
		))

	var wg sync.WaitGroup // Sincroniza as goroutines
	wg.Add(3)             // Adiciona 2 tarefas ao WaitGroup

	NextRoutine := make(chan bool)
	count := 3

	// Inicia os processos em paralelo usando goroutines
	go func(id int) {
		defer wg.Done()
		i := int64(0)
		for i = 0; i < limit; i += 10 { // Aumenta o progresso
			if count > 1 {
				<-NextRoutine // Aguarda o sinal para continuar
			}
			for j := 0; j < 10; j++ {
				if _, err := io.CopyN(tempFile1, rand.Reader, 10); err != nil {
					panic(err)
				}
				bar1.IncrBy(1)                    // Atualiza a primeira barra de progresso
				time.Sleep(time.Millisecond * 20) // Simula um atraso para visualizar o progresso
			}
			if count > 1 {
				NextRoutine <- true // Sinaliza a próxima goroutine para executar
			}
		}
		count--
	}(1)

	go func(id int) {
		defer wg.Done()
		i := int64(0)
		for i = 0; i < 100; i += 10 { // Aumenta o progresso
			if count > 1 {
				<-NextRoutine // Aguarda o sinal para continuar
			}
			for j := 0; j < 10; j++ {
				if _, err := io.CopyN(tempFile2, rand.Reader, 10); err != nil {
					panic(err)
				}
				bar2.IncrBy(1)                    // Atualiza a segunda barra de progresso
				time.Sleep(time.Millisecond * 20) // Simula um atraso para visualizar o progresso
			}
			if count > 1 {
				NextRoutine <- true // Sinaliza a próxima goroutine para executar
			}
		}
		count--
	}(2)

	go func(id int) {
		defer wg.Done()
		i := int64(0)
		for i = 0; i < 200; i += 10 { // Aumenta o progresso
			if count > 1 {
				<-NextRoutine // Aguarda o sinal para continuar
			}
			for j := 0; j < 10; j++ {
				if _, err := io.CopyN(tempFile3, rand.Reader, 10); err != nil {
					panic(err)
				}
				bar3.IncrBy(1)                    // Atualiza a terceira barra de progresso
				time.Sleep(time.Millisecond * 20) // Simula um atraso para visualizar o progresso
			}
			if count > 1 {
				NextRoutine <- true // Sinaliza a próxima goroutine para executar
			}
		}
		count--
	}(3)

	NextRoutine <- true // Inicia o ciclo de execução

	// Aguarda todas as goroutines terminarem
	wg.Wait()

	// Finaliza o progresso principal
	p.Wait()
}

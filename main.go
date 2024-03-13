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
	var limit int64 = 1024 * 1024 // 1MB

	// Cria dois arquivos temporários para simular dois downloads
	tempFile1, err := ioutil.TempFile("", "download1")
	if err != nil {
		panic(err)
	}
	defer tempFile1.Close() // Garante que o arquivo vai ser fechado ao terminar o programa

	tempFile2, err := ioutil.TempFile("", "download2")
	if err != nil {
		panic(err)
	}
	defer tempFile2.Close()

	// Cria um progresso principal
	p := mpb.New(mpb.WithWaitGroup(&sync.WaitGroup{})) // Cria uma nova instância de progresso usando a biblioteca mpb. A opção WithWaitGroup é usada para sincronizar o progresso com um WaitGroup.

	// Cria duas barras de progresso
	bar1 := p.AddBar(limit, mpb.PrependDecorators(
		decor.Name("Download 1"),
		decor.CountersNoUnit("%d / %d", decor.WCSyncSpace),
	), mpb.AppendDecorators(
		decor.OnComplete(decor.Name("          "), " Completed"),
	))

	bar2 := p.AddBar(limit, mpb.PrependDecorators(
		decor.Name("Download 2"),
		decor.CountersNoUnit("%d / %d", decor.WCSyncSpace),
	))

	var wg sync.WaitGroup // Sincroniza as goroutines
	wg.Add(2)             // Adiciona 2 tarefas ao WaitGroup

	// Inicia os downloads em paralelo usando goroutines
	go func() {
		defer wg.Done()
		for i := int64(0); i < limit; i += 1024 { // Aumenta o progresso a cada 1024 bytes
			if _, err := io.CopyN(tempFile1, rand.Reader, 1024); err != nil {
				panic(err)
			}
			bar1.IncrBy(10240)                // Atualiza a primeira barra de progresso
			time.Sleep(time.Millisecond * 50) // Simula um atraso para visualizar o progresso
		}
	}()

	go func() {
		defer wg.Done()
		for i := int64(0); i < limit; i += 1024 { // Aumenta o progresso a cada 1024 bytes
			if _, err := io.CopyN(tempFile2, rand.Reader, 1024); err != nil {
				panic(err)
			}
			bar2.IncrBy(1024)                 // Atualiza a segunda barra de progresso
			time.Sleep(time.Millisecond * 50) // Simula um atraso para visualizar o progresso
		}
	}()

	// Aguarda todas as goroutines terminarem
	wg.Wait()

	// Finaliza o progresso principal
	p.Wait()
}

package taskcreate

import (
	"math/rand"
	"strconv"
	"time"
)

const TamanhoPriority int = 11

type TaskStruct struct {
	Nome     string
	Id       int
	Priority int
	Quantum  int
	//Status   string
}

func Taskcreate(numero int, Nquantum int, Npriority int) TaskStruct {
	stringNumero := strconv.Itoa(numero)

	task := TaskStruct{}
	task.Nome = "Process " + stringNumero

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	task.Id = r.Intn(201) + 1

	quantuns := [100]int{
		127, 51, 73, 222, 147, 126, 135, 139, 149, 70, 77,
		243, 29, 156, 134, 120, 99, 112, 261, 54, 48, 245,
		56, 208, 243, 35, 91, 230, 39, 35, 265, 151, 26,
		95, 94, 221, 31, 181, 287, 201, 240, 166, 257, 243,
		298, 38, 250, 49, 115, 213, 215, 37, 194, 126, 21, 99,
		80, 258, 39, 76, 57, 227, 82, 150, 211, 60, 204, 293,
		248, 240, 294, 74, 224, 276, 226, 59, 258, 261, 34, 234,
		248, 93, 251, 198, 172, 111, 80, 103, 52, 282, 48, 284,
		192, 100, 184, 137, 194, 54, 272, 271,
	}

	priorities := [TamanhoPriority]int{
		7, 1, 10, 8, 3, 10, 5, 6, 2, 9, 4,
	}

	task.Quantum = quantuns[Nquantum]
	task.Priority = priorities[Npriority]
	//task.Status = "disponivel"

	return task

}

func TaskVetorCreator(QuantidadeDeTask int) []TaskStruct {

	taskVetor := make([]TaskStruct, QuantidadeDeTask)

	for i := 0; i < QuantidadeDeTask; i++ {

		PriorityIndex := i % TamanhoPriority

		newTask := Taskcreate(i+1, i, PriorityIndex)

		taskVetor[i] = newTask

	}

	return taskVetor

}

package taskcreate

import (
	"math/rand"
	"strconv"
	"time"
)

type taskStruct struct {
	Nome     string
	Id       int
	Priority int
	Quantum  float64
	Status   string
}

func Taskcreate(numero int, Nquantum int, Npriority int) taskStruct {
	stringNumero := strconv.Itoa(numero)

	task := taskStruct{}
	task.Nome = "Process" + stringNumero

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	task.Id = r.Intn(201) + 1

	quantuns := [100]int{
		20, 78, 47, 86, 93, 88, 85, 90, 82, 98, 68, 96, 52, 84, 74, 46, 62, 49, 66, 72,
		92, 73, 97, 91, 89, 99, 65, 63, 21, 27, 48, 61, 83, 51, 60, 75, 29, 26, 28, 67,
		69, 80, 22, 57, 94, 77, 55, 70, 87, 79, 23, 64, 95, 81, 24, 76, 25, 71, 30, 54,
		58, 56, 53, 31, 100, 59, 50, 45, 32, 33, 34, 36, 37, 35, 38, 39, 40, 41, 42, 43,
		44, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	}

	priorities := [11]int{
		7, 4, 11, 8, 3, 10, 5, 6, 2, 9,
	}

	task.Quantum = float64(quantuns[Nquantum])
	task.Priority = priorities[Npriority]
	task.Status = "disponivel"

	return task

}

func TaskVetorCreator(QuantidadeDeTask int) []taskStruct {

	taskVetor := make([]taskStruct, 10)

	for i := 0; i < QuantidadeDeTask; i++ {

		newTask := Taskcreate(i+1, i, i)

		taskVetor = append(taskVetor, newTask)

	}

	return taskVetor

}

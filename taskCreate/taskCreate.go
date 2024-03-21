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
	QuantumI int
}

func Taskcreate(numero int, Nquantum int, Npriority int, testType int) TaskStruct {
	stringNumero := strconv.Itoa(numero)

	task := TaskStruct{}
	task.Nome = "Process " + stringNumero

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	task.Id = r.Intn(201) + 1

	var quantuns [50]int

	switch testType {
	case 1:
		{
			quantuns = [50]int{
				30, 46, 28, 33, 50, 28, 39, 50, 60,
				51, 40, 33, 40, 51, 45, 23, 37, 23,
				25, 47, 51, 37, 56, 26, 56, 55, 46,
				22, 51, 21, 27, 21, 26, 42, 30, 41,
				44, 50, 24, 43, 54, 47, 33, 29, 59,
				46, 52, 28, 37, 47,
			}

		}
	case 2:
		{
			quantuns = [50]int{
				248, 225, 232, 210, 271, 292, 283, 245,
				266, 276, 217, 224, 208, 231, 292, 203,
				271, 200, 240, 214, 275, 259, 270, 276,
				280, 290, 272, 227, 252, 242, 277, 220,
				243, 231, 244, 280, 226, 288, 234, 299,
				247, 250, 283, 204, 278, 212, 227, 258,
				216, 204,
			}
		}
	case 3:
		{
			quantuns = [50]int{
				127, 51, 73, 222, 147, 126, 135, 139, 149, 70, 77,
				243, 29, 156, 134, 120, 99, 112, 261, 54, 48, 245,
				56, 208, 243, 35, 91, 230, 39, 35, 265, 151, 26,
				95, 94, 221, 31, 181, 287, 201, 240, 166, 257, 243,
				298, 38, 250, 49, 115, 213,
			}
		}
	default:
		{
			for i := range quantuns {
				quantuns[i] = rand.Intn(391) + 10
			}
		}
	}

	priorities := [TamanhoPriority]int{
		7, 1, 10, 8, 3, 10, 5, 6, 2, 9, 4,
	}

	task.Quantum = quantuns[Nquantum]
	task.QuantumI = task.Quantum
	task.Priority = priorities[Npriority]

	return task

}

func TaskVetorCreator(QuantidadeDeTask int, testType int) []TaskStruct {

	taskVetor := make([]TaskStruct, QuantidadeDeTask)

	for i := 0; i < QuantidadeDeTask; i++ {

		PriorityIndex := i % TamanhoPriority

		newTask := Taskcreate(i+1, i, PriorityIndex, testType)

		taskVetor[i] = newTask

	}

	return taskVetor

}

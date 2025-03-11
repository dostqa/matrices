package matrix

import (
	"fmt"
)

// Определение типа MatrixBase
type MatrixBase struct {
	data [][]int
}

/*=============== Взаимодействия с столбцами и строками ======================*/

// возвращает количество строк в матрице
func (matrix MatrixBase) CountRow() int {
	return cap(matrix.data)
}

// возвращает количество столбцов в матрице
func (matrix MatrixBase) CountCol() int {
	return cap(matrix.data[0])
}

// метод удаляет строку
func (matrix *MatrixBase) DelRow(matrixRow int) MatrixBase {

	matrixRow -= 1
	result := *matrix

	// выполняем сдвиг влево на один индекс
	copy(result.data[matrixRow:], result.data[matrixRow+1:])
	// удаляем последний элемент (записываем нулевое значение)
	result.data[cap(result.data)-1] = nil
	// усекаем срез
	result.data = result.data[:cap(result.data)-1]

	*matrix = result

	return *matrix
}

// метод удаляет столбец
func (matrix *MatrixBase) DelColumn(matrixRow int, matrixCol int) MatrixBase {
	result := *matrix
	matrixCol -= 1

	// выполняем сдвиг влево на один индекс
	copy(
		result.data[matrixRow][matrixCol:],
		result.data[matrixRow][matrixCol+1:],
	)
	// удаляем последний элемент (записываем нулевое значение)
	result.data[matrixRow][cap(result.data[matrixRow])-1] = 0
	// усекаем срез
	result.data[matrixRow] = result.data[matrixRow][:cap(result.data[matrixRow])-1]

	*matrix = result

	return *matrix
}

// метод записи значения в ячейку
func (matrix *MatrixBase) SetValue(row int, col int, value int) {
	matrix.data[row-1][col-1] = value
}

// метод чтения значения из ячейки
func (matrix MatrixBase) GetValue(row int, col int) int {
	return matrix.data[row-1][col-1]
}

/*============ Методы инициализации и заполнения матрицы =============*/

// Метод перезаписи матрицы
func (matrix *MatrixBase) Fill(countRow int, countCol int) {
	// создаем свободное место для элементов матрицы
	matrix.PrepareToFill(countRow, countCol)
	// получаем данные от пользователя
	matrix.ConsoleInput()
}

// метод создаёт матрицу
func (matrix *MatrixBase) Initialize(input [][]int) {
	result := *matrix
	result.PrepareToFill(cap(input), cap(input[0]))

	for i := range input {
		copy(result.data[i], input[i])
	}

	*matrix = result
}

// метод создаёт матрицу
func (matrix *MatrixBase) Create(input [][]int) MatrixBase {
	// matrix.PrepareToFill(cap(input), cap(input[0]))

	(*matrix).data = [][]int{
		{0, -1, 2},
		{1, 0, -2},
		{3, 1, 2},
	}
	return *matrix
}

// метод размечает массив для матрицы
func (matrix *MatrixBase) PrepareToFill(countRow int, countCol int) {
	result := *matrix
	// создаем матрицу с заданным количеством строк и помещаем её в ячейку памяти, где хранится наша матрица
	result.data = make([][]int, countRow)

	// создаем в каждой строке нужное количество мест для элементов
	for i := range result.data {
		result.data[i] = make([]int, countCol)
	}

	*matrix = result
}

// Ввод матрицы через консоль
func (matrix *MatrixBase) ConsoleInput() {
	result := *matrix

	// Вводим значения в каждую ячейку матрицы
	for matrixRow := range result.data {
		for matrixCol := range result.data[matrixRow] {

			// метод ввода из консоли
			fmt.Scan(&result.data[matrixRow][matrixCol])
		}
	}

	*matrix = result
}

// выводит матрицу в консоль в удобочитаемом виде
func (matrix MatrixBase) ShowInConsole() {
	for matrixRow := range matrix.data {
		fmt.Println("Cтрока", matrixRow+1, matrix.data[matrixRow])
	}
}

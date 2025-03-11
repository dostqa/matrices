package matrix

import (
	"fmt"
	"slices"
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
func (matrix *MatrixBase) DelRow(indexRow int) MatrixBase {
	// Сдвигаем индекс, так как индексация с 0
	indexRow--
	result := *matrix

	// Подготавливаем матрицу для уменьшения размера на 1 строку
	result.PrepareToFill(matrix.CountRow()-1, matrix.CountCol())

	// Удаляем строку из данных матрицы
	copy(result.data, slices.Delete(matrix.data, indexRow, indexRow+1))
	// result.data = append(result.data[:indexRow], result.data[indexRow+1:]...)

	// Обновляем исходную матрицу
	*matrix = result

	return *matrix
}

// метод удаляет столбец
func (matrix *MatrixBase) DelColumn(indexCol int) MatrixBase {
	// Сдвигаем индекс, так как индексация с 0
	indexCol--

	result := *matrix

	// Подготовим матрицу для уменьшения размера
	result.PrepareToFill(matrix.CountRow(), matrix.CountCol()-1)

	// fmt.Println("Размерность до:", result.CountRow(), result.CountCol())

	// Копируем данные в новую матрицу, удаляя столбец
	for matrixRow := range result.data {
		// Удаляем столбец из текущей строки
		copy(result.data[matrixRow], slices.Delete(matrix.data[matrixRow], indexCol, indexCol+1))

	}

	// Размерность не изменится, так как мы не изменяем длину среза
	// fmt.Println("Размерность после:", result.CountRow(), result.CountCol())

	*matrix = result // Обновляем исходную матрицу
	return *matrix
}

// метод записи значения в ячейку
func (matrix *MatrixBase) SetValue(row int, col int, value int) {
	matrix.data[row][col] = value
}

// метод чтения значения из ячейки
func (matrix MatrixBase) GetValue(row int, col int) int {
	return matrix.data[row][col]
}

func (matrix *MatrixBase) Get(row int, col int) *int {
	return &(matrix.data[row][col])
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
func (matrix *MatrixBase) Create(input [][]int) *MatrixBase {
	// matrix.PrepareToFill(cap(input), cap(input[0]))

	(*matrix).data = input
	return matrix
}

// метод размечает массив для матрицы
func (matrix *MatrixBase) PrepareToFill(countRow int, countCol int) {

	result := *matrix
	// создаем матрицу с заданным количеством строк и помещаем её в ячейку памяти, где хранится наша матрица
	result.data = make([][]int, countRow)

	// fmt.Println(result)
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
	for matrixRow := range result.CountRow() {
		for matrixCol := range result.CountRow() {

			// метод ввода из консоли
			fmt.Scan(&result.data[matrixRow][matrixCol])
		}
	}

	*matrix = result
}

// выводит матрицу в консоль в удобочитаемом виде
func (matrix MatrixBase) ShowInConsole() {
	for matrixRow := range matrix.CountRow() {
		fmt.Println("Cтрока", matrixRow+1, matrix.data[matrixRow])
	}
}

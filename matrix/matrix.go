package matrix

import "fmt"

// Определение типа Matrix
type Matrix [][]int

/*=============== Взаимодействия с столбцами и строками ======================*/

// возвращает количество строк в матрице
func (matrix Matrix) CountRow() int {
	return cap(matrix)
}

// возвращает количество столбцов в матрице
func (matrix Matrix) CountCol() int {
	return cap(matrix[0])
}

// метод удаляет строку
func (matrix *Matrix) DelRow(matrixRow int) {

	matrixRow -= 1
	result := *matrix

	// выполняем сдвиг влево на один индекс
	copy(result[matrixRow:], result[matrixRow+1:])
	// удаляем последний элемент (записываем нулевое значение)
	result[cap(result)-1] = nil
	// усекаем срез
	result = result[:cap(result)-1]

	*matrix = result
}

// метод удаляет столбец
func (matrix *Matrix) DelColumn(matrixRow int, matrixCol int) {
	result := *matrix
	matrixCol -= 1

	// выполняем сдвиг влево на один индекс
	copy(
		result[matrixRow][matrixCol:],
		result[matrixRow][matrixCol+1:],
	)
	// удаляем последний элемент (записываем нулевое значение)
	result[matrixRow][cap(result[matrixRow])-1] = 0
	// усекаем срез
	result[matrixRow] = result[matrixRow][:cap(result[matrixRow])-1]

	*matrix = result
}

/*============ Методы инициализации и заполнения матрицы =============*/

// возвращает подматрицу матрицы
func (matrix Matrix) FindSubMatrix(i int, j int) Matrix {
	var subMatrix Matrix
	subMatrix.PrepareToFill(matrix.CountRow()-1, matrix.CountCol()-1)

	for matrixRow := range matrix {
		switch {
		case matrixRow == i-1:
			// удаляем строку
			matrix.DelRow(i)
		default:
			continue
		}
	}

	for matrixRow := range matrix {
		for matrixCol := range matrix[matrixRow] {
			switch {
			case matrixCol == j-1:
				// удаляем столбец
				matrix.DelColumn(matrixRow, j)
			default:
				continue
			}
		}
	}

	subMatrix = matrix

	return subMatrix
}

// Метод перезаписи матрицы
func (matrix *Matrix) Fill(countRow int, countCol int) {
	// создаем свободное место для элементов матрицы
	matrix.PrepareToFill(countRow, countCol)
	// получаем данные от пользователя
	matrix.ConsoleInput()
}

// метод размечает массив для матрицы
func (matrix *Matrix) PrepareToFill(countRow int, countCol int) {

	// создаем матрицу с заданным количеством строк и помещаем её в ячейку памяти, где хранится наша матрица
	*(matrix) = make([][]int, countRow)

	// создаем в каждой строке нужное количество мест для элементов
	for i := range *(matrix) {
		(*matrix)[i] = make([]int, countCol)
	}
}

// Ввод матрицы через консоль
func (matrix *Matrix) ConsoleInput() {
	got := *matrix
	fmt.Println("Введите матрицу: ")

	// Вводим значения в каждую ячейку матрицы
	for matrixRow := range got {
		for matrixCol := range got[matrixRow] {

			// метод ввода из консоли
			fmt.Scan(&got[matrixRow][matrixCol])
		}
	}

	*matrix = got
}

// выводит матрицу в консоль в удобочитаемом виде
func (matrix Matrix) ShowInConsole() {
	for matrixRow := range matrix {
		fmt.Println("Cтрока", matrixRow+1, matrix[matrixRow])
	}
}

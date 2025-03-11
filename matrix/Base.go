package matrix

import "fmt"

// Определение типа MatrixBase
type MatrixBase [][]int

/*=============== Взаимодействия с столбцами и строками ======================*/

// возвращает количество строк в матрице
func (matrix MatrixBase) CountRow() int {
	return cap(matrix)
}

// возвращает количество столбцов в матрице
func (matrix MatrixBase) CountCol() int {
	return cap(matrix[0])
}

// метод удаляет строку
func (matrix *MatrixBase) DelRow(matrixRow int) MatrixBase {

	matrixRow -= 1
	result := *matrix

	// выполняем сдвиг влево на один индекс
	copy(result[matrixRow:], result[matrixRow+1:])
	// удаляем последний элемент (записываем нулевое значение)
	result[cap(result)-1] = nil
	// усекаем срез
	result = result[:cap(result)-1]

	*matrix = result

	return *matrix
}

// метод удаляет столбец
func (matrix *MatrixBase) DelColumn(matrixRow int, matrixCol int) MatrixBase {
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

	return *matrix
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
	got := *matrix
	got.PrepareToFill(cap(input), cap(input[0]))

	for i := range input {
		copy(got[i], input[i])
	}

	*matrix = got
}

// метод создаёт матрицу
func (matrix *MatrixBase) Create(input [][]int) MatrixBase {
	// matrix.PrepareToFill(cap(input), cap(input[0]))

	(*matrix) = [][]int{
		{0, -1, 2},
		{1, 0, -2},
		{3, 1, 2},
	}
	return *matrix
}

// метод размечает массив для матрицы
func (matrix *MatrixBase) PrepareToFill(countRow int, countCol int) {

	// создаем матрицу с заданным количеством строк и помещаем её в ячейку памяти, где хранится наша матрица
	*(matrix) = make([][]int, countRow)

	// создаем в каждой строке нужное количество мест для элементов
	for i := range *(matrix) {
		(*matrix)[i] = make([]int, countCol)
	}
}

// Ввод матрицы через консоль
func (matrix *MatrixBase) ConsoleInput() {
	got := *matrix

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
func (matrix MatrixBase) ShowInConsole() {
	for matrixRow := range matrix {
		fmt.Println("Cтрока", matrixRow+1, matrix[matrixRow])
	}
}

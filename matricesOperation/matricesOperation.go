package matricesOperation

import m "matrices/matrix"

// умножает матрицу на матрицу
func MultiplyOfMatrix(firstMatrix, secondMatrix m.Matrix) (matrix m.Matrix, err bool) {

	// проверяем возможно ли умножение
	if firstMatrix.CountCol() != secondMatrix.CountRow() {
		// умножение невозможно: возвращаем пустую матрицу и наличие ошибки
		return m.Matrix{{}}, true
	}

	// переменная в которую сохраним результат умножения
	var result m.Matrix
	// создаем свободное место для элементов матрицы
	result.PrepareToFill(firstMatrix.CountRow(), secondMatrix.CountCol())

	// непосредственно умножаем
	for matrixRow := range result {
		for matrixCol := range result[matrixRow] {
			for i := range secondMatrix.CountRow() {
				result[matrixRow][matrixCol] += firstMatrix[matrixRow][i] * secondMatrix[i][matrixCol]
			}

		}
	}

	return result, false
}

// складывает матрицы
func SumOfMatrix(firstMatrix, secondMatrix m.Matrix) (matrix m.Matrix, err bool) {
	// проверяем возможно ли сложение

	if firstMatrix.CountRow() != secondMatrix.CountRow() || firstMatrix.CountCol() != secondMatrix.CountCol() {
		// сложение невозможно: возвращаем пустую матрицу и наличие ошибки
		return m.Matrix{{}}, true
	}

	// переменная в которую сохраним результат сложения
	var result m.Matrix
	// создаем свободное место для элементов матрицы
	result.PrepareToFill(firstMatrix.CountRow(), firstMatrix.CountCol())

	// непосредственно складываем
	for matrixRow := range result {
		for matrixCol := range result[matrixRow] {
			result[matrixRow][matrixCol] = firstMatrix[matrixRow][matrixCol] + secondMatrix[matrixRow][matrixCol]
		}
	}

	return result, false
}

// находит разность матриц
func DiffOfMatrix(firstMatrix, secondMatrix m.Matrix) (matrix m.Matrix, err bool) {
	// проверяем возможно ли вычитание
	if firstMatrix.CountRow() != secondMatrix.CountRow() || firstMatrix.CountCol() != secondMatrix.CountCol() {
		// вычитание невозможно: возвращаем пустую матрицу и наличие ошибки
		return m.Matrix{{}}, true
	}

	// переменная в которую сохраним результат вычитания
	var result m.Matrix
	// создаем свободное место для элементов матрицы
	result.PrepareToFill(firstMatrix.CountRow(), firstMatrix.CountCol())

	// непосредственно вычитаем
	for matrixRow := range result {
		for matrixCol := range result[matrixRow] {
			result[matrixRow][matrixCol] = firstMatrix[matrixRow][matrixCol] - secondMatrix[matrixRow][matrixCol]
		}
	}

	return result, false
}

// умножает матрицу на число
func MultiplyImmutable(matrix m.Matrix, multiplier int) m.Matrix {

	// умножаем каждый элемент матрицы на число num
	for matrixRow := range matrix {
		for matrixCol := range matrix[matrixRow] {
			matrix[matrixRow][matrixCol] *= multiplier
		}
	}

	// возвращаем копию
	return matrix
}

// Метод транспонирует матрицу
func TransposedImmutable(matrix m.Matrix) m.Matrix {

	// переменная в которую сохраним результат транспонирования
	var result m.Matrix

	// создаем свободное место для элементов матрицы
	result.PrepareToFill(matrix.CountCol(), matrix.CountRow())

	// непосредственно транспонируем
	for matrixRow := range matrix {
		for matrixCol := range matrix[matrixRow] {
			result[matrixCol][matrixRow] = matrix[matrixRow][matrixCol]
		}
	}

	// возвращаем результат
	return result
}

// возвращает симметричную матрицу
func SymmetricImmutable(matrix m.Matrix) m.Matrix {
	result, _ := MultiplyOfMatrix(matrix, TransposedImmutable(matrix))
	return result
}

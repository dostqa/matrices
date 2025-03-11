package matrix

type MatrixEnhance struct {
	MatrixBase
}

// // Конструктор для MatrixEnhance
// func NewMatrixEnhance(data [][]int) MatrixEnhance {
// 	var matrix MatrixEnhance
// 	matrix.Initialize(data)
// 	return matrix
// }

func (matrix *MatrixEnhance) Create(input [][]int) MatrixEnhance {
	(*matrix).MatrixBase = new(MatrixBase).Create(input)
	// matrix.MatrixBase =
	// fmt.Println(&matrix)
	return *matrix
}

// умножает матрицу на матрицу
func (matrix MatrixEnhance) MultiplyOfMatrix(secondMatrix MatrixEnhance) MatrixEnhance {

	// проверяем возможно ли умножение
	if matrix.CountCol() != secondMatrix.CountRow() {
		// умножение невозможно: вызываем исключение
		panic("multiplication is not possible: the number of columns in the first matrix is not equal to the number of rows in the second matrix")
	}

	// переменная в которую сохраним результат умножения
	var result MatrixEnhance
	// создаем свободное место для элементов матрицы
	result.PrepareToFill(matrix.CountRow(), secondMatrix.CountCol())

	// непосредственно умножаем
	for matrixRow := range result.MatrixBase {
		for matrixCol := range result.MatrixBase[matrixRow] {
			for i := range secondMatrix.CountRow() {
				result.MatrixBase[matrixRow][matrixCol] +=
					matrix.MatrixBase[matrixRow][i] *
						secondMatrix.MatrixBase[i][matrixCol]
			}

		}
	}

	return result
}

// складывает матрицы
func (matrix MatrixEnhance) SumOfMatrix(secondMatrix MatrixEnhance) MatrixEnhance {
	// проверяем возможно ли сложение

	if matrix.CountRow() != secondMatrix.CountRow() || matrix.CountCol() != secondMatrix.CountCol() {
		// сложение невозможно: вызываем исключение
		panic("addition is not possible: the number of rows or columns in the matrices do not match")
	}

	// переменная в которую сохраним результат сложения
	var result MatrixEnhance
	// создаем свободное место для элементов матрицы
	result.PrepareToFill(matrix.CountRow(), matrix.CountCol())

	// непосредственно складываем
	for matrixRow := range result.MatrixBase {
		for matrixCol := range result.MatrixBase[matrixRow] {
			result.MatrixBase[matrixRow][matrixCol] =
				matrix.MatrixBase[matrixRow][matrixCol] +
					secondMatrix.MatrixBase[matrixRow][matrixCol]
		}
	}

	return result
}

// находит разность матриц
func (matrix MatrixEnhance) DiffOfMatrix(secondMatrix MatrixEnhance) MatrixEnhance {
	// проверяем возможно ли вычитание
	if matrix.CountRow() != secondMatrix.CountRow() ||
		matrix.CountCol() != secondMatrix.CountCol() {
		// вычитание невозможно: вызываем исключение
		panic("subtraction is not possible: the number of rows or columns in the matrices do not match")
	}

	// переменная в которую сохраним результат вычитания
	var result MatrixEnhance
	// создаем свободное место для элементов матрицы
	result.PrepareToFill(matrix.CountRow(), matrix.CountCol())

	// непосредственно вычитаем
	for matrixRow := range result.MatrixBase {
		for matrixCol := range result.MatrixBase[matrixRow] {
			result.MatrixBase[matrixRow][matrixCol] =
				matrix.MatrixBase[matrixRow][matrixCol] -
					secondMatrix.MatrixBase[matrixRow][matrixCol]
		}
	}

	return result
}

// умножает матрицу на число
func (matrix MatrixEnhance) MultiplyImmutable(multiplier int) MatrixEnhance {

	// переменная в которую сохраним результат вычитания
	var result MatrixEnhance
	// создаем свободное место для элементов матрицы
	result.PrepareToFill(matrix.CountRow(), matrix.CountCol())

	// умножаем каждый элемент матрицы на число num
	for matrixRow := range result.MatrixBase {
		for matrixCol := range result.MatrixBase[matrixRow] {
			result.MatrixBase[matrixRow][matrixCol] =
				matrix.MatrixBase[matrixRow][matrixCol] *
					multiplier
		}
	}

	// возвращаем копию
	return result
}

// Метод транспонирует матрицу
func (matrix MatrixEnhance) TransposedImmutable() MatrixEnhance {

	// переменная в которую сохраним результат транспонирования
	var result MatrixEnhance

	// создаем свободное место для элементов матрицы
	result.PrepareToFill(matrix.CountCol(), matrix.CountRow())

	// непосредственно транспонируем
	for matrixRow := range result.MatrixBase {
		for matrixCol := range result.MatrixBase[matrixRow] {
			result.MatrixBase[matrixCol][matrixRow] = matrix.MatrixBase[matrixRow][matrixCol]
		}
	}

	// возвращаем результат
	return result
}

// возвращает симметричную матрицу
func (matrix MatrixEnhance) SymmetricImmutable() MatrixEnhance {
	result := matrix.MultiplyOfMatrix(matrix.TransposedImmutable())
	return result
}

// возвращает подматрицу матрицы
func (matrix MatrixEnhance) GetSubMatrix(coordRow int, coordCol int) MatrixEnhance {

	var result MatrixEnhance
	// fmt.Println(matrix)
	result.PrepareToFill(matrix.CountRow()-1, matrix.CountCol()-1)

	for matrixRow := range matrix.MatrixBase {
		switch {
		case matrixRow == coordRow-1:
			// удаляем строку
			matrix.DelRow(coordRow)
		default:
			continue
		}
	}

	for matrixRow := range matrix.MatrixBase {
		for matrixCol := range matrix.MatrixBase[matrixRow] {
			switch {
			case matrixCol == coordCol-1:
				// удаляем столбец
				matrix.DelColumn(matrixRow, coordCol)
			default:
				continue
			}
		}
	}

	result = matrix

	return result
}

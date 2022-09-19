/*
Задачи:
1) Доработать функцию matrix.fill()
	1.1) Реализовать input от пользователя
2) Доработать функцию matrix.show()
	2.1) Разобраться с управляющими последовательностями
	2.2) Подписывать при выводе каждую строку и столбец.
	2.3) Возможно показывать размерность матрицы
3) Переработать код в пользу удобочитаемости с использованием
прижившихся в линейной алгебре обозначений
*/

package main

import "fmt"

func matrixMultiplication(firstMatrix, secondMatrix Matrix) Matrix { // умножает матрицу на матрицу; возвращает результат
	var result Matrix                                       // переменная в которую сохраним результат умножения
	result.prepareToFill(firstMatrix.m(), secondMatrix.n()) // создаем свободное место для элементов матрицы

	for m := range result { // непосредственно умножаем
		for n := range result[m] {
			for i := 0; i < secondMatrix.m(); i++ {
				result[m][n] += firstMatrix[m][i] * secondMatrix[i][n]
			}

		}
	}

	return result
}

func sumOfMatrix(firstMatrix, secondMatrix Matrix) Matrix { // складывает матрицы; возвращает результат
	var result Matrix                                      // переменная в которую сохраним результат сложения
	result.prepareToFill(firstMatrix.m(), firstMatrix.n()) // создаем свободное место для элементов матрицы

	for m := range result { // непосредственно складываем
		for n := range result[m] {
			result[m][n] = firstMatrix[m][n] + secondMatrix[m][n]
		}
	}

	return result
}

func diffOfMatrix(firstMatrix, secondMatrix Matrix) Matrix { // находит разность матриц; возвращает результат
	var result Matrix                                      // переменная в которую сохраним результат вычитания
	result.prepareToFill(firstMatrix.m(), firstMatrix.n()) // создаем свободное место для элементов матрицы

	for m := range result { // непосредственно вычитаем
		for n := range result[m] {
			result[m][n] = firstMatrix[m][n] - secondMatrix[m][n]
		}
	}

	return result
}

type Matrix [][]int

func (matrix Matrix) m() int { // возвращает количество строк в матрице
	return cap(matrix) // получаем и сразу же возвращаем
}

func (matrix Matrix) n() int { // возвращает количество столбцов в матрице
	return cap(matrix[0]) // получаем и сразу же возвращаем
}

func (matrix *Matrix) prepareToFill(m int, n int) { // создаем свободное место для элементов матрицы
	*(matrix) = make([][]int, m) // создаем матрицу с заданным количеством строк и помещаем её в ячейку памяти, где хранится наша матрица
	for i := range *(matrix) {   // создаем в каждой строке нужное количество мест для элементов
		(*matrix)[i] = make([]int, n)
	}
}

func (matrix *Matrix) get() { // получает матрицу
	got := *matrix
	fmt.Println("Введите матрицу: ")

	for m := range got { // непосредственно построчно получаем матрицу
		for n := range got[m] {
			fmt.Scan(&got[m][n])
		}
	}

	*matrix = got

	return
}

func (matrix *Matrix) fill(m int, n int) { // наполняет матрицу значениями, или заменяет значения уже существующей матрицы
	matrix.prepareToFill(m, n) // создаем свободное место для элементов матрицы
	matrix.get()               // получаем данные от пользователя
}

func (matrix Matrix) show() { // выводит матрицу в консоль в удобочитаемом виде
	for m := range matrix {
		fmt.Println("строка", m+1, matrix[m])
	}
}

func (matrix Matrix) multiplyByNum(num int) Matrix { // умножает матрицу на число; возвращает результат
	for m := range matrix { // работаем с копией, так что всё хорошо
		for n := range matrix[m] {
			matrix[m][n] *= num
		}
	}

	return matrix // возвращаем копию
}

func (matrix Matrix) transposed() Matrix { // возвращает транспонированную матрицу
	var result Matrix                            // переменная в которую сохраним результат транспонирования
	result.prepareToFill(matrix.n(), matrix.m()) // создаем свободное место для элементов матрицы

	for m := range matrix { // непосредственно транспонируем
		for n := range matrix[m] {
			result[n][m] = matrix[m][n]
		}
	}

	return result // возвращаем результат
}

func (matrix Matrix) symmetric() Matrix { // возвращает симметричную матрицу
	return matrixMultiplication(matrix, matrix.transposed())
}

func main() {
	/*matrix := Matrix{
		{1, 3, 0, 4},
		{1, 1, 3, 4},
		{4, 0, 0, 4},
	}
	matrix.multiplyByNum(3).show()

	fmt.Println("Первая матрица: ")
	firstMatrix := Matrix{
		{3, 2, 1},
		{4, 0, 2},
	}
	firstMatrix.show()

	fmt.Println("Вторая матрица: ")
	secondMatrix := Matrix{
		{1, 3, 0, 4},
		{1, 1, 3, 4},
		{4, 0, 0, 4},
	}
	secondMatrix.show()

	fmt.Println("Результат умножения матриц: ")
	matrixMultiplication(firstMatrix, secondMatrix).show()

	fmt.Println("Транспонированная вторая матрица: ")
	secondMatrix.transposed().show()*/

	var firstMatrix Matrix
	var secondMatrix Matrix

	firstMatrix.fill(3, 3)
	firstMatrix.show()

	secondMatrix.fill(3, 3)
	secondMatrix.show()
	fmt.Println("Результат сложения:")
	sumOfMatrix(firstMatrix, secondMatrix).show()

}

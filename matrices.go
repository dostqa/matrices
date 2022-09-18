/*
Задачи:
1) Доработать функцию matrix.fill()
	1.1) Реализовать input от пользователя
2) Доработать функцию matrix.show()
	2.1) Разобраться с управляющими последовательностями
	2.2) Подписывать при выводе каждую строку и столбец.
	2.3) Возможно показывать размерность матрицы
*/

package main

import "fmt"

func matrixMultiplication(firstMatrix, secondMatrix Matrix) Matrix { // умножает матрицу на матрицу; возвращает результат
	var result Matrix // результат умножения
	return result
}

func sumOfMatrix(firstMatrix, secondMatrix Matrix) Matrix { // складывает матрицы; возвращает результат
	var result Matrix                                      // результат сложения
	result.prepareToFill(firstMatrix.m(), firstMatrix.n()) // создаем свободное место для элементов матрицы

	for i := range result {
		for n := range result {
			result[i][n] = firstMatrix[i][n] + secondMatrix[i][n]
		}
	}

	return result
}

func getMatrix(matrix Matrix) { // получает матрицу
	fmt.Println("Введите матрицу: ")

	for i := range matrix {
		for n := range matrix[i] {
			fmt.Scan(&matrix[i][n])
		}
	}

	return
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

func (matrix *Matrix) fill(m int, n int) { // наполняет матрицу значениями, или заменяет значения уже существующей матрицы
	matrix.prepareToFill(m, n) // создаем свободное место для элементов матрицы
	getMatrix(*(matrix))       // получаем данные от пользователя
}

func (matrix Matrix) show() { // выводит матрицу в консоль в удобочитаемом виде
	for i := range matrix {
		fmt.Println("строка", i, matrix[i])
	}
}

func (matrix Matrix) transposed() Matrix { // возвращает транспонированную матрицу
	var result Matrix
	result.prepareToFill(matrix.m(), matrix.n())

	for m := range matrix {
		for n := range matrix {
			result[m][n] = matrix[n][m]
		}
	}

	return result // возвращаем результат
}

func main() {
	/*var firstMatrix Matrix
	var secondMatrix Matrix

	firstMatrix.fill(3, 3)
	firstMatrix.show()

	secondMatrix.fill(3, 3)
	secondMatrix.show()

	fmt.Println("Результат сложения:")
	sumOfMatrix(firstMatrix, secondMatrix).show()*/

	var matrix Matrix
	matrix.fill(3, 3)

	matrix.transposed().show()
}

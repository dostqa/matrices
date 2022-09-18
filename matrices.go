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

func getMatrix(matrix Matrix) { //получает матрицу
	fmt.Println("Введите матрицу: ")

	for i := 0; i < cap(matrix); i++ {
		for n := 0; n < cap(matrix[i]); n++ {
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

func (matrix *Matrix) fill(m int, n int) { // наполняет матрицу значениями, или заменяет значения уже существующей матрицы
	*(matrix) = make([][]int, m) // создаем матрицу с заданным количеством строк и помещаем её в ячейку памяти, где хранится наша матрица
	for i := range *(matrix) {   // создаем в каждой строке нужное количество мест для элементов
		(*matrix)[i] = make([]int, n)
	}

	getMatrix(*(matrix)) // получаем данные от пользователя
}

func (matrix Matrix) show() { // выводит матрицу в консоль в удобочитаемом виде
	for i := range matrix {
		fmt.Println("строка", i, matrix[i])
	}
}

func (matrix Matrix) transposed() Matrix { // возвращает транспонированную матрицу
	return matrix // работаем с копией, а не оригиналом, поэтому всё хорошо
}

func main() {
	var matrix Matrix
	matrix.fill(3, 3)
	matrix.show()
}

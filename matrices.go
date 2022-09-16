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

type Matrix [][]int

func (matrix Matrix) m() int { // возвращает количество строк в матрице
	return cap(matrix) // получаем и сразу же возвращаем
}

func (matrix Matrix) n() int { // возвращает количество столбцов в матрице
	return cap(matrix[0]) // получаем и сразу же возвращаем
}

func (matrix *Matrix) fill(m int, n int) { // наполняет матрицу значениями, или заменяет значения уже существующей матрицы
	*(matrix) = make([][]int, m) // создаем матрицу с заданным количеством строк и помещаем её в ячейку памяти, где хранится наша матрица
	for i := range *(matrix) {   // помещаем в каждую строку нужное количество элементов
		(*matrix)[i] = make([]int, n)
	}
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
	matrix.fill(4, 5)
	matrix.show()
}

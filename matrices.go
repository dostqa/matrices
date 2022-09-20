/*
Задачи:
1) реализовать пользовательский интерфейс
2) прописать условия там, где это нужно
*/

package main

import "fmt"

func det(matrix Matrix) int { // возвращает детерминант матрицы
	return 0
}

func matrixMultiplication(firstMatrix, secondMatrix Matrix) (matrix Matrix, err bool) { // умножает матрицу на матрицу; возвращает результат
	switch { // проверяем возможно ли умножение
	case firstMatrix.n() != secondMatrix.m():
		return Matrix{{}}, true // умножение невозможно: возвращаем пустую матрицу и наличие ошибки
	}

	var result Matrix                                       // переменная в которую сохраним результат умножения
	result.prepareToFill(firstMatrix.m(), secondMatrix.n()) // создаем свободное место для элементов матрицы

	for m := range result { // непосредственно умножаем
		for n := range result[m] {
			for i := 0; i < secondMatrix.m(); i++ {
				result[m][n] += firstMatrix[m][i] * secondMatrix[i][n]
			}

		}
	}

	return result, false
}

func sumOfMatrix(firstMatrix, secondMatrix Matrix) (matrix Matrix, err bool) { // складывает матрицы; возвращает результат
	switch { // проверяем возможно ли сложение
	case firstMatrix.m() != secondMatrix.m() || firstMatrix.n() != secondMatrix.n():
		return Matrix{{}}, true // сложение невозможно: возвращаем пустую матрицу и наличие ошибки
	}

	var result Matrix                                      // переменная в которую сохраним результат сложения
	result.prepareToFill(firstMatrix.m(), firstMatrix.n()) // создаем свободное место для элементов матрицы

	for m := range result { // непосредственно складываем
		for n := range result[m] {
			result[m][n] = firstMatrix[m][n] + secondMatrix[m][n]
		}
	}

	return result, false
}

func diffOfMatrix(firstMatrix, secondMatrix Matrix) (matrix Matrix, err bool) { // находит разность матриц; возвращает результат
	switch { // проверяем возможно ли вычитание
	case firstMatrix.m() != secondMatrix.m() || firstMatrix.n() != secondMatrix.n():
		return Matrix{{}}, true // вычитание невозможно: возвращаем пустую матрицу и наличие ошибки
	}

	var result Matrix                                      // переменная в которую сохраним результат вычитания
	result.prepareToFill(firstMatrix.m(), firstMatrix.n()) // создаем свободное место для элементов матрицы

	for m := range result { // непосредственно вычитаем
		for n := range result[m] {
			result[m][n] = firstMatrix[m][n] - secondMatrix[m][n]
		}
	}

	return result, false
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
	result, _ := matrixMultiplication(matrix, matrix.transposed())
	return result
}

func (matrix *Matrix) delRow(m int) { // метод удаляет строку
	m -= 1
	result := *matrix

	copy(result[m:], result[m+1:])  // выполняем сдвиг влево на один индекс
	result[cap(result)-1] = nil     // удаляем последний элемент (записываем нулевое значение)
	result = result[:cap(result)-1] // усекаем срез

	*matrix = result
}

func (matrix *Matrix) delColumn(m int, n int) { // метод удаляет столбец
	result := *matrix
	n -= 1

	copy(result[m][n:], result[m][n+1:])     // выполняем сдвиг влево на один индекс
	result[m][cap(result[m])-1] = 0          // удаляем последний элемент (записываем нулевое значение)
	result[m] = result[m][:cap(result[m])-1] // усекаем срез

	*matrix = result
}

func (matrix Matrix) findSubmatrix(i int, j int) Matrix { // возвращает подматрицу матрицы
	var submatrix Matrix
	submatrix.prepareToFill(matrix.m()-1, matrix.n()-1)

	for m := range matrix {
		switch {
		case m == i-1:
			matrix.delRow(i) // удаляем строку
			break

		default:
			continue
		}
	}

	for m := range matrix {
		for n := range matrix[m] {
			switch {
			case n == j-1:
				matrix.delColumn(m, j) // удаляем столбец

			default:
				continue
			}
		}
	}

	submatrix = matrix

	return submatrix
}

func main() {
	var A Matrix = Matrix{
		{4, 5, 6},
		{-1, 0, 3},
		{-1, 2, -1},
	}

	A.findSubmatrix(2, 2).show()

	/*var B Matrix = Matrix{
		{0, -1, 2},
		{1, 0, -2},
		{3, 1, 2},
	}

	fmt.Println("В первом действии получится: ")
	first, _ := matrixMultiplication(A, A)
	first.show()
	fmt.Println(" ")

	fmt.Println("Во втором действии получится: ")
	second, _ := sumOfMatrix(A, B)
	second.show()
	fmt.Println(" ")

	fmt.Println("В третьем действии получится: ")
	third := B.multiplyByNum(3)
	third.show()
	fmt.Println(" ")

	fmt.Println("В четвертом действии получится: ")
	fourth, _ := diffOfMatrix(A, third)
	fourth.show()
	fmt.Println(" ")

	fmt.Println("В пятом действии получится: ")
	fifth, _ := matrixMultiplication(second, fourth)
	fifth.show()
	fmt.Println(" ")

	fmt.Println("В шестом действии получится: ")
	sixth, _ := diffOfMatrix(first, fifth)
	sixth.show()
	fmt.Println(" ")

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

	/*var firstMatrix Matrix
	var secondMatrix Matrix

	firstMatrix.fill(3, 3)
	firstMatrix.show()

	secondMatrix.fill(3, 3)
	secondMatrix.show()
	fmt.Println("Результат сложения:")
	sumOfMatrix(firstMatrix, secondMatrix).show()*/
}

package main

import m "matrices/matrix"

func main() {
	var matrix m.Matrix = m.Matrix{
		{4, 5, 6},
		{-1, 0, 3},
		{-1, 2, -1},
	}

	matrix.FindSubMatrix(2, 2).ShowInConsole()

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

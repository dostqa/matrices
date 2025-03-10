package matrix

type MatrixInterface interface {
	CountRow() int
	CountCol() int
	DelRow(matrixRow int)
	DelColumn(matrixRow int, matrixCol int)
	Fill(countRow int, countCol int)
	PrepareToFill(countRow int, countCol int)
	ConsoleInput()
	ShowInConsole()
}
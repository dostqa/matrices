package matrix

type MatrixInterface interface {
	CountRow() int
	CountCol() int
	DelRow(matrixRow int)
	DelColumn(matrixRow int, matrixCol int)
	SetValue(row int, col int, value int)
	GetValue(row int, col int) int
	Get(row int, col int) *int
	Fill(countRow int, countCol int)
	PrepareToFill(countRow int, countCol int)
	ConsoleInput()
	ShowInConsole()
}
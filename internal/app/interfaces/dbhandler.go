package interfaces

type DBHandler interface {
	Create(value string) error
	Retrieve() ([][]byte, [][]byte, error)
	Delete(key int) error
}
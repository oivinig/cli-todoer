package interfaces

type DBHandler interface {
	Create(key []byte, value []byte) error
	Retrieve(key []byte) []byte
	Update(key []byte, value []byte)
	Delete(key []byte) 
}
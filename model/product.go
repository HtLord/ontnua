package model

type Inventory struct {
	Ban string
	Products []Product
}
type Product struct{
	Ban string
	Name string
	Quality int32
}

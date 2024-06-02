package main

import (
	"fmt"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("Explorando GORM!")
}

type Produto struct {
	gorm.Model
	Codigo string
	Preco  uint
}

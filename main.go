package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Explorando GORM!")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Tabela de Produtos
	db.AutoMigrate(&Produto{})

	// Cria alguns registros
	db.Create(&Produto{Codigo: "F5EA54", Preco: 123})
	db.Create(&Produto{Codigo: "DFE332", Preco: 132})
	db.Create(&Produto{Codigo: "453D4F", Preco: 213})
	db.Create(&Produto{Codigo: "FF546A", Preco: 231})
	db.Create(&Produto{Codigo: "BA12AE", Preco: 312})
	db.Create(&Produto{Codigo: "ABCEDF", Preco: 321})

	// Leitura de dados
	var produto Produto
	db.First(&produto, 1) // Buscar produto pela chave primária
	fmt.Println(&produto)
	db.First(&produto, "Codigo = ?", "F5EA54") // Buscar produto pelo código
	fmt.Println(&produto)

	// Update - Atualização de dados
	db.Model(&produto).Update("Preco", 150)
	fmt.Println(&produto)

	// Updades - Atualização de múltiplos campos
	db.Model(&produto).Updates(Produto{Preco: 572, Codigo: "BA12AE"}) // non-zero fields
	fmt.Println(&produto)
	db.Model(&produto).Updates(map[string]interface{}{"Preco": 987, "Codigo": "BA12AE"})
	fmt.Println(&produto)

	// Delete - remove o produto
	db.Delete(&produto, 1)
}

type Produto struct {
	gorm.Model
	Codigo string
	Preco  uint
}

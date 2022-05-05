package models

import (
	"app/app/Structs"
	"app/app/db"
)

func ReturnDate() []Structs.Product {
	db := db.Connect()
	row, err := db.Query("select * from product")

	if err != nil {
		panic(err.Error())
	}

	product := Structs.Product{}
	var products []Structs.Product

	for row.Next() {
		var id, qtd int
		var name, desc string
		var price float64

		//Detalhe, que neste caso, o Scan deve seguir a ordem das colunas no banco
		err := row.Scan(&id, &name, &desc, &price, &qtd)
		if err != nil {
			panic(err.Error())
		}

		product.Name = name
		product.Price = price
		product.Desc = desc
		product.Qtd = qtd

		products = append(products, product)
	}

	return products
}

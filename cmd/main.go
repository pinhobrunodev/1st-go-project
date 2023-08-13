package main

import (
	"database/sql"

	"github.com/pinhobrunodev/pfa-go/internal/order/infra/database"
	"github.com/pinhobrunodev/pfa-go/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/orders")
	if err != nil {
		panic(err)
	}
	// -> defer fecha a conexao com o banco dps de rodar tudo.
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 100,
		Tax:   10,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output.FinalPrice)
}

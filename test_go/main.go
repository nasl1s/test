package main

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

func main() {
	// Определяем схему GraphQL
	schema := `
		type Query {
			hello: String
		}
	`

	// Создаем экземпляр схемы GraphQL
	s, err := graphql.ParseSchema(schema, &Resolver{})
	if err != nil {
		panic(err)
	}

	// Выполняем запрос GraphQL
	query := `
		query {
			hello
		}
	`

	ctx := context.Background()
	result := s.Exec(ctx, query, "", nil)
	if len(result.Errors) > 0 {
		fmt.Printf("Ошибка при выполнении запроса: %v\n", result.Errors)
		return
	}

	// Выводим результат запроса
	fmt.Println(result.Data)
}

// Resolver определяет методы для обработки запросов
type Resolver struct{}

func (r *Resolver) Hello() *string {
	message := "Hello, world!"
	return &message
}

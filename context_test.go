package belajar_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T){
	background := context.Background()
	fmt.Println(background)	

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextValue(t *testing.T){
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)

	fmt.Println(contextD.Value("b"))
}
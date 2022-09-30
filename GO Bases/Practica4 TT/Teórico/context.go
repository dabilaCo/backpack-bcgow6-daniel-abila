package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main(){

	ctx := context.Background()

	ctx2 := context.WithValue(ctx, "clave", "valor")

	//fmt.Printf("%v\n",ctx)

	//Function(ctx2, 10)

	//ctx3, _ := context.WithDeadline(ctx2, time.Now().Add(5 * time.Second))
	//<-ctx3.Done()

	//fmt.Println(ctx3.Err().Error())

	ctx3, _ := context.WithTimeout(ctx2, 5 * time.Second)
	<-ctx3.Done()

	fmt.Println(ctx3.Err().Error())

}

func Function(ctx context.Context, dato int){
	valor := ctx.Value("Clave")

	//.(string) es para asegurarnos que en una funcion vacia nos envian el tipo de dato correcto
	str, ok := valor.(string)

	if !ok{
		fmt.Println("no es string")
		os.Exit(1)
	}

	fmt.Printf("%T\n",str)
}
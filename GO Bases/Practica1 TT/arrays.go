package main

import "fmt"
func main(){
	/*var myArray [3] int
	myArray[0] = 10
	fmt.Println(myArray)*/
	var arr [5] int
	var mySlice = [len(arr)]int{1,2,3,4,5}
	/*slc := mySlice[:4] //te muestra desde 0 a 4
	slc = mySlice[:] //te muestra todos los valores del slice 0 a 5*/
	slc := mySlice[0:2] //te muestra desde el indice 1 hasta el indice 2, no incluye el 3
	slc = append(slc, 5, 7, 6, 7, 4)
	fmt.Println(cap(slc), len(slc))
}
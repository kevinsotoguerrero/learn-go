package mypackage

import "fmt"

//Modificadores de acceso, con la inicial en mayuscula son publicos, visibles desde cualquier clase, en minuscula son privados y solo se pueden ver desde el mismo package
type CarPublic struct {
	Brand string
	Year  int
}

type CarPublic2 struct {
	Brand string
	Year  int
}

func Message() {
	fmt.Println("oes")
}

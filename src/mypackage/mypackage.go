package mypackage

//Modificadores de acceso, con la inicial en mayuscula son publicos, visibles desde cualquier clase, en minuscula son privados y solo se pueden ver desde el mismo package. privados con la inicial en minuscula
type CarPublic struct {
	Brand string
	Year  int
}

type CarPublic2 struct {
	brand string
	Year  int
}

package chainofresponsibility

type Department interface {
	execute(*Patient)
	setNext(Department) Department
}

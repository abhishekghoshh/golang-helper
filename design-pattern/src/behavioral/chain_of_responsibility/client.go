package chainofresponsibility

import "fmt"

func Main() {
	cashier := &Cashier{}
	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "Abhishek Ghosh"}
	//Patient visiting
	reception.execute(patient)

	fmt.Println()
	// another way of function chaining or chain of responsibility
	reception.
		setNext(doctor).
		setNext(medical).
		setNext(medical).
		setNext(cashier)
	reception.execute(patient)

	fmt.Println()
	reception.
		setNext(doctor).
		setNext(medical).
		setNext(medical).
		setNext(cashier)
	reception.execute(&Patient{name: "Abhishek"})

}

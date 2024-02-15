package chainofresponsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient for " + p.name)
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (c *Medical) setNext(next Department) Department {
	c.next = next
	return next
}

package chainofresponsibility

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done for " + p.name)
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}
func (c *Doctor) setNext(next Department) Department {
	c.next = next
	return next
}

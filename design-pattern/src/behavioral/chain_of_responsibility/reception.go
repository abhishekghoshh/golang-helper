package chainofresponsibility

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done for " + p.name)
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (c *Reception) setNext(next Department) Department {
	c.next = next
	return next
}

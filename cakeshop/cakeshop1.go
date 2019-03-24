// cake shop in serial work flow

package main

import (
	"fmt"
	"time"
)

type Shop struct {
	cakes        int           // quantity of cakes to be made
	bakeTime     time.Duration // time to bake
	iceTime      time.Duration // time to ice
	inscribeTime time.Duration // time to inscribe
}

type cake int

func (s Shop) bake(baked chan cake) {
	defer close(baked)

	for i := 0; i < s.cakes; i++ {
		c := cake(i)

		fmt.Println("baking", i)
		work(s.bakeTime)
		fmt.Println(i, "baked")

		baked <- c
	}
}

func (s Shop) ice(iced chan cake, baked chan cake) {
	defer close(iced)

	for c := range baked {
		fmt.Println("icing", c)
		work(s.iceTime)
		fmt.Println(c, "iced")

		iced <- c
	}
}

func (s Shop) inscribe(iced chan cake) {
	for c := range iced {
		fmt.Println("inscribing", c)
		work(s.inscribeTime)
		fmt.Println(c, "finished")
	}
}

func (s Shop) Work() {
	baked := make(chan cake)
	iced := make(chan cake)

	go s.bake(baked)
	go s.ice(iced, baked)
	s.inscribe(iced)
}

func work(duration time.Duration) {
	time.Sleep(duration)
}

func main() {
	shop := Shop{
		cakes:        3,
		bakeTime:     1 * time.Second,
		iceTime:      2 * time.Second,
		inscribeTime: 3 * time.Second,
	}

	shop.Work()
}

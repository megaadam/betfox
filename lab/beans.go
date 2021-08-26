package lab

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/megaadam/betfox/brain"
	"github.com/megaadam/betfox/server"
)

type beanCounts struct {
	red   int
	white int
	green int
}

func beanCounter(bc *beanCounts) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		col := rand.Intn(3)
		switch col {
		case 0:
			bc.red++

		case 1:
			bc.white++

		default:
			bc.green++
		}
	}
}

func beanWatch(bc *beanCounts) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond * 2)
		fmt.Println(bc)
	}
}

// Beans shall rule
func Beans() {
	bc := beanCounts{}
	go beanCounter(&bc)
	go beanWatch(&bc)

	time.Sleep(20 * time.Second)

	fmt.Println(brain.GetBrain())
	fmt.Println(server.GetServer())
}

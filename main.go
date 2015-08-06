package random

import (
	"math/rand"
	"time"
)

func SudoRandomNumer() int {
	return rand.Intn(time.Now().Second())
}

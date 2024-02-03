package random

import (
	"math/rand"
	"time"
)

var Number string

func Randint() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100000)
	Number = string(num)
}

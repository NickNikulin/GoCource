package brain

/* Let's create two data structure - EmptyCup and BeLikeWater
with a general method - Make() */

import "fmt"

type EmptyCupBrain struct {
}

func (c *EmptyCupBrain) Make() {
	fmt.Println("Open your mingt")
}

type BeLikeWaterBrain struct {
}

func (w *BeLikeWaterBrain) Make() {
	fmt.Println("I'm water")
}

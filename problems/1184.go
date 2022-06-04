package problems

import (
	"fmt"
	"math"
)

func (p Problem) Run1184() {
	fmt.Println(distanceBetweenBusStops([]int{1,2,3,4}, 0, 1))
}

func distanceBetweenBusStops(distance []int, start, destination int) int {
	if start > destination {
		temp := start
		start = destination
		destination = temp
	}

	var res, total int
	for i := 0; i < len(distance); i++ {
		if i >= start && i < destination {
			res += distance[i]
		}
		total += distance[i]
	}

	return int(math.Min(float64(res), float64(total - res)))
}
package main

import (
	"fmt"
	"github.com/wmatsushita/gotoolbox/collections"
	"math"
)

type point struct {
	x int
	y int
}

func (p1 point) distance(p2 point) (dist float64) {
	dist = math.Sqrt(float64(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2)))

	return
}

type pair struct {
	points []point
	dist   float64
}

func minDistPairs(p1, p2 pair) pair {
	if math.Min(p1.dist, p2.dist) == p1.dist {
		return p1
	} else {
		return p2
	}
}

func sortPointsByAxis(points []point, isX bool) (sorted []point) {

	if len(points) < 2 {
		return points
	}

	cutpoint := len(points) / 2

	left := sortPointsByAxis(points[:cutpoint], isX)
	right := sortPointsByAxis(points[cutpoint:], isX)

	i := 0
	j := 0

	for i < len(left) || j < len(right) {
		if i >= len(left) {
			sorted = append(sorted, right[j:]...)
			break
		}
		if j >= len(right) {
			sorted = append(sorted, left[i:]...)
			break
		}
		if isX {
			if left[i].x <= right[j].x {
				sorted = append(sorted, left[i])
				i++
			} else {
				sorted = append(sorted, right[j])
				j++
			}
		} else {
			if left[i].y <= right[j].y {
				sorted = append(sorted, left[i])
				i++
			} else {
				sorted = append(sorted, right[j])
				j++
			}
		}
	}

	return

}

func bruteForceClosestPair(points []point) (pair pair) {
	pair.dist = math.MaxFloat64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := points[i].distance(points[j])
			if dist < pair.dist {
				pair.dist = dist
				pair.points = []point{points[i], points[j]}
			}
		}
	}

	return
}

func splitPoints(pointsX, pointsY []point) (leftX, rightX, leftY, rightY []point) {
	cutpoint := len(pointsX) / 2

	//left, right := NewPointSet(), NewPointSet()
	left := collections.NewSet()

	for i, v := range pointsX {
		if i < cutpoint {
			left.Add(v)
			leftX = append(leftX, v)
		} else {
			//right.add(v)
			rightX = append(rightX, v)
		}
	}

	for _, v := range pointsY {
		if left.Contains(v) {
			leftY = append(leftY, v)
		} else {
			rightY = append(rightY, v)
		}
	}

	return
}

func closestPairRecursion(pointsX, pointsY []point) pair {
	// base case, using pointsX or pointsY makes no difference
	if len(pointsX) <= 3 {
		return bruteForceClosestPair(pointsX)
	}

	leftX, rightX, leftY, rightY := splitPoints(pointsX, pointsY)

	pairL := closestPairRecursion(leftX, leftY)
	pairR := closestPairRecursion(rightX, rightY)
	delta := minDistPairs(pairL, pairR)

	pairS := closestSplitPair(pointsX, pointsY, delta)

	return minDistPairs(delta, pairS)

}

func closestSplitPair(pointsX, pointsY []point, delta pair) pair {
	divide := len(pointsX)/2 - 1
	center := pointsX[divide].x
	remaining := []point{}
	for _, p := range pointsY {
		if math.Abs(p.x-center) < delta.dist {
			remaining = append(remaining, p)
		}
	}

}

func ClosestPair(points []point) pair {
	// sort points by X and Y (this runs in n.log(n))
	pointsX := sortPointsByAxis(points, true)
	pointsY := sortPointsByAxis(points, false)

	return closestPairRecursion(pointsX, pointsY)
}

func main() {
	points := []point{
		{1, 1}, {5, 4}, {6, 10}, {4, 9},
	}

	fmt.Println(sortPointsByAxis(points, true))
	fmt.Println(sortPointsByAxis(points, false))

	fmt.Println(points[0].distance(points[1]))
	fmt.Println(points[0].distance(points[2]))
	fmt.Println(points[0].distance(points[3]))
	fmt.Println(points[1].distance(points[2]))
	fmt.Println(points[1].distance(points[3]))
	fmt.Println(points[2].distance(points[3]))

	fmt.Println(bruteForceClosestPair(points))

	fmt.Println(ClosestPair(points))

}

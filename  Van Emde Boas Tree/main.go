package main

import (
	"fmt"
	"math"
)

type VEBTree struct {
	u       int // Universe size (2^k)
	min     int // Minimum element in the tree
	max     int // Maximum element in the tree
	summary *VEBTree
	cluster []*VEBTree
}

func newVEBTree(u int) *VEBTree {
	if u <= 2 {
		return &VEBTree{
			u:       u,
			min:     math.MaxInt32,
			max:     -1,
			summary: nil,
			cluster: nil,
		}
	}

	sqrtU := int(math.Sqrt(float64(u)))
	summary := newVEBTree(sqrtU)
	cluster := make([]*VEBTree, sqrtU)
	for i := 0; i < sqrtU; i++ {
		cluster[i] = newVEBTree(sqrtU)
	}

	return &VEBTree{
		u:       u,
		min:     math.MaxInt32,
		max:     -1,
		summary: summary,
		cluster: cluster,
	}
}

func (v *VEBTree) high(x int) int {
	return x / int(math.Sqrt(float64(v.u)))
}

func (v *VEBTree) low(x int) int {
	return x % int(math.Sqrt(float64(v.u)))
}

func (v *VEBTree) index(x, y int) int {
	return x*int(math.Sqrt(float64(v.u))) + y
}

func (v *VEBTree) minimum() int {
	return v.min
}

func (v *VEBTree) maximum() int {
	return v.max
}

func (v *VEBTree) member(x int) bool {
	if x == v.min || x == v.max {
		return true
	}
	if v.u <= 2 {
		return false
	}
	return v.cluster[v.high(x)].member(v.low(x))
}

func (v *VEBTree) successor(x int) int {
	if v.u == 2 {
		if x == 0 && v.max == 1 {
			return 1
		}
		return -1
	}
	if v.min != -1 && x < v.min {
		return v.min
	}
	maxLow := v.cluster[v.high(x)].maximum()
	if maxLow != -1 && v.low(x) < maxLow {
		offset := v.cluster[v.high(x)].successor(v.low(x))
		return v.index(v.high(x), offset)
	}
	successorCluster := v.summary.successor(v.high(x))
	if successorCluster == -1 {
		return -1
	}
	offset := v.cluster[successorCluster].minimum()
	return v.index(successorCluster, offset)
}

func (v *VEBTree) predecessor(x int) int {
	if v.u == 2 {
		if x == 1 && v.min == 0 {
			return 0
		}
		return -1
	}
	if v.max != -1 && x > v.max {
		return v.max
	}
	minLow := v.cluster[v.high(x)].minimum()
	if minLow != -1 && v.low(x) > minLow {
		offset := v.cluster[v.high(x)].predecessor(v.low(x))
		return v.index(v.high(x), offset)
	}
	predecessorCluster := v.summary.predecessor(v.high(x))
	if predecessorCluster == -1 {
		if v.min != -1 && x > v.min {
			return v.min
		}
		return -1
	}
	offset := v.cluster[predecessorCluster].maximum()
	return v.index(predecessorCluster, offset)
}

func (v *VEBTree) insert(x int) {
	if v.min == -1 {
		v.min = x
		v.max = x
		return
	}
	if x < v.min {
		x, v.min = v.min, x
	}
	if v.u > 2 {
		if v.cluster[v.high(x)].minimum() == -1 {
			v.summary.insert(v.high(x))
			v.cluster[v.high(x)].min, v.cluster[v.high(x)].max = v.low(x), v.low(x)
		} else {
			v.cluster[v.high(x)].insert(v.low(x))
		}
	}
	if x > v.max {
		v.max = x
	}
}

func (v *VEBTree) delete(x int) {
	if v.min == v.max {
		v.min = -1
		v.max = -1
		return
	}
	if v.u == 2 {
		if x == 0 {
			v.min = 1
		} else {
			v.min = 0
		}
		v.max = v.min
	} else {
		if x == v.min {
			firstCluster := v.summary.minimum()
			x = v.index(firstCluster, v.cluster[firstCluster].minimum())
			v.min = x
		}
		v.cluster[v.high(x)].delete(v.low(x))
		if v.cluster[v.high(x)].minimum() == -1 {
			v.summary.delete(v.high(x))
			if x == v.max {
				summaryMax := v.summary.maximum()
				if summaryMax == -1 {
					v.max = v.min
				} else {
					v.max = v.index(summaryMax, v.cluster[summaryMax].maximum())
				}
			}
		} else if x == v.max {
			v.max = v.index(v.high(x), v.cluster[v.high(x)].maximum())
		}
	}
}

func main() {
	v := newVEBTree(16)

	v.insert(5)
	v.insert(10)
	v.insert(2)

	fmt.Println("Minimum:", v.minimum())
	fmt.Println("Maximum:", v.maximum())

	fmt.Println("Successor of 5:", v.successor(5))
	fmt.Println("Predecessor of 5:", v.predecessor(5))

	v.delete(10)

	fmt.Println("Minimum:", v.minimum())
	fmt.Println("Maximum:", v.maximum())
}

package algorithm

/**
 * Reference: https://oi-wiki.org/ds/dsu/
 *            https://zhuanlan.zhihu.com/p/93647900
 */

type UFSet struct {
	UF   []int
	Rank []int
	size int
}

func UFSetConstructor(size int) *UFSet {
	uf := &UFSet{
		UF:   make([]int, size),
		Rank: make([]int, size),
		size: size,
	}
	for i := 1; i < size; i++ {
		uf.UF[i] = i
		uf.Rank[i] = 1
	}
	return uf
}

func (uf *UFSet) Find(x int) int {
	if x != uf.UF[x] {
		parent := uf.Find(uf.UF[x])
		if parent != uf.UF[x] {
			uf.Rank[uf.UF[x]]--
			uf.UF[x] = parent
		}
	}
	return uf.UF[x]
}

func (uf *UFSet) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	if uf.Rank[x] <= uf.Rank[y] {
		uf.UF[x] = y
		uf.Rank[y] += uf.Rank[x]
	} else {
		uf.UF[y] = x
		uf.Rank[x] += uf.Rank[y]
	}
}

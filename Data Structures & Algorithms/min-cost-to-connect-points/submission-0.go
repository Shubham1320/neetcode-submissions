type DSU struct {
	Parent []int
	Size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{Parent: parent, Size: size}
}

func (dsu *DSU) find(node int) int {
	if dsu.Parent[node] != node {
		dsu.Parent[node] = dsu.find(dsu.Parent[node])
	}
	return dsu.Parent[node]
}

func (dsu *DSU) union(u, v int) bool {
	pu, pv := dsu.find(u), dsu.find(v)
	if pu == pv {
		return false
	}
	if dsu.Size[pu] < dsu.Size[pv] {
		pu, pv = pv, pu
	}
	dsu.Size[pu] += dsu.Size[pv]
	dsu.Parent[pv] = pu
	return true
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dsu := NewDSU(n)
	var edges [][]int
	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dist := int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
			edges = append(edges, []int{dist, i, j})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a][0] < edges[b][0]
	})

	res := 0
	for _, edge := range edges {
		dist, u, v := edge[0], edge[1], edge[2]
		if dsu.union(u, v) {
			res += dist
		}
	}
	return res
}

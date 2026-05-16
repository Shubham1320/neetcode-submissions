type DSU struct {
    Parent, Size []int
}

func NewDSU(n int) *DSU {
    dsu := &DSU{
        Parent: make([]int, n+1),
        Size:   make([]int, n+1),
    }
    for i := 0; i <= n; i++ {
        dsu.Parent[i] = i
        dsu.Size[i] = 1
    }
    return dsu
}

func (dsu *DSU) Find(node int) int {
    if dsu.Parent[node] != node {
        dsu.Parent[node] = dsu.Find(dsu.Parent[node])
    }
    return dsu.Parent[node]
}

func (dsu *DSU) Union(u, v int) bool {
    pu, pv := dsu.Find(u), dsu.Find(v)
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

func (dsu *DSU) Connected(u, v int) bool {
    return dsu.Find(u) == dsu.Find(v)
}

func swimInWater(grid [][]int) int {
    N := len(grid)
    dsu := NewDSU(N * N)
    positions := make([][3]int, 0, N*N)
    for r := 0; r < N; r++ {
        for c := 0; c < N; c++ {
            positions = append(positions, [3]int{grid[r][c], r, c})
        }
    }
    sort.Slice(positions, func(i, j int) bool {
        return positions[i][0] < positions[j][0]
    })
    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    for _, pos := range positions {
        t, r, c := pos[0], pos[1], pos[2]
        for _, d := range directions {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nc >= 0 && nr < N && nc < N && grid[nr][nc] <= t {
                dsu.Union(r*N+c, nr*N+nc)
            }
        }
        if dsu.Connected(0, N*N-1) {
            return t
        }
    }
    return -1
}

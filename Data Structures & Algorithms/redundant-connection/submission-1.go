func findRedundantConnection(edges [][]int) []int {

    n := len(edges)
    adj := make([][]int, n+1)
    visit := make([]bool, n+1)

    var dfs func(node, par int) bool
    dfs = func(node, par int) bool {
        if visit[node] {
            return true
        }
        visit[node] = true
        for _, nei := range adj[node] {
            if nei == par {
                continue
            }
            if dfs(nei, node) {
                return true
            }
        }
        return false
    }

    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
        for i := 0; i <= n; i++ {
            visit[i] = false
        }

        if dfs(u, -1) {
            return []int{u, v}
        }
    }
    return []int{}

}




func validTree(n int, edges [][]int) bool {

    adjList := make(map[int][]int)

    for i := 0; i < len(edges); i++ {
        first := edges[i][0]
        second := edges[i][1]

        if _, ok := adjList[first]; !ok {
            adjList[first] = []int{second}
        } else {
            adjList[first] = append(adjList[first], second)
        }

        if _, ok := adjList[second]; !ok {
            adjList[second] = []int{first}
        } else {
            adjList[second] = append(adjList[second], first)
        }

    }

    fmt.Println(adjList)

    var dfs func(int, int) bool
    visit := make(map[int]int)
    fmt.Println(visit)

    dfs = func(parent int, node int) bool {

        if visit[node] == 1 {
            return false
        }

        if _, ok := adjList[node]; !ok {
            visit[node] = 1
            return true
        }

        visit[node] = 1
        for _, val := range adjList[node] {
            if parent == val {
                continue
            }
            if !dfs(node, val) {
                return false
            }
        }
        return true
    }

    return dfs(-1, 0) && len(visit) == n
}

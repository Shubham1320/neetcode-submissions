func countComponents(n int, edges [][]int) int {

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

    var dfs func(int, int)
    visit := make(map[int]int)
    fmt.Println(visit)

    dfs = func(parent int, node int) {

        if visit[node] == 1 {
            return
        }

        if _, ok := adjList[node]; !ok {
            visit[node] = 1
            return
        }

        visit[node] = 1
        for _, val := range adjList[node] {
            if parent == val {
                continue
            }
            dfs(node, val)
        }
    }

    ans := 0
    root := 0
    for len(visit) != n {
        if _, ok := visit[root]; !ok {
            dfs(-1, root)
            ans++
        }
        root++
    }
    return ans

}

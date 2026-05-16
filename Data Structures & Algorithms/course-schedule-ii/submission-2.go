func findOrder(numCourses int, prerequisites [][]int) []int {

    var ans []int

    prevMap := make(map[int][]int)

    for i := 0; i < len(prerequisites); i++ {

        first := prerequisites[i][0]
        second := prerequisites[i][1]

        if _, ok := prevMap[first]; !ok {
            prevMap[first] = []int{second}
        } else {
            prevMap[first] = append(prevMap[first], second)
        }
    }

    visited := make([]int, numCourses)
    alreadyDone := make(map[int]bool)
    fmt.Println(prevMap)

    var dfs func(int) bool

    dfs = func(node int) bool{

        if visited[node] == 1 {
            return false
        }

        if _, ok := alreadyDone[node]; ok {
            return true
        }

        if _, ok := prevMap[node]; !ok {
            ans = append(ans, node)
            alreadyDone[node] = true
            return true
        }

        visited[node] = 1

        for _, val := range prevMap[node] {
            if !dfs(val) {
                return false
            }
        }

        ans = append(ans, node)
        alreadyDone[node] = true

        visited[node] = 0
        return true
    }

    for i := 0; i < numCourses; i++ {
        dfs(i) 
        fmt.Println(visited) 
    }

    if len(ans) != numCourses {
        return []int{}
    }

    return ans

}
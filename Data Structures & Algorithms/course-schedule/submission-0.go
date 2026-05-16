func canFinish(numCourses int, prerequisites [][]int) bool {

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
    fmt.Println(prevMap)

    var dfs func(int) bool

    dfs = func(node int) bool {

        if visited[node] == 1 {
            return false
        }

        if _, ok := prevMap[node]; !ok {
            return true
        }

        visited[node] = 1

        ans := true
        for _, val := range prevMap[node] {
            ans = ans && dfs(val)
        }

        if !ans {
            return false
        }

        visited[node] = 0
        return true

    }

    for i := 0; i < numCourses; i++ {

        if !dfs(i) {
            return false
        }

    }

    return true


}

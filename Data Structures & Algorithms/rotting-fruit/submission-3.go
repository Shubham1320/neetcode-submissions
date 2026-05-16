
func orangesRotting(grid [][]int) int {

    queue := [][2]int{}
    freshFruits := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i, j})
            } else if grid[i][j] == 1 {
                freshFruits++
            }
        }
    }

    //[1,1,0],
    //[0,1,1],
    //[0,1,2]

    time := 0
    if freshFruits == 0 {
        return time
    }

    for len(queue) != 0 {

        tempQueue := [][2]int{}
        time++
        for len(queue) != 0 {
            current := queue[0]

            i := current[0]
            j := current[1]

            fmt.Println(freshFruits, i, j)

            if i-1 >= 0 {
                if grid[i-1][j] == 1 {
                    grid[i-1][j] = 2
                    freshFruits--
                    tempQueue = append(tempQueue, [2]int{i-1, j})
                }
            }

            if i+1 < len(grid) {
                if grid[i+1][j] == 1 {
                    grid[i+1][j] = 2
                    freshFruits--
                    tempQueue = append(tempQueue, [2]int{i+1, j})
                }
            }

            if j-1 >= 0 {
                if grid[i][j-1] == 1 {
                    grid[i][j-1] = 2
                    freshFruits--
                    tempQueue = append(tempQueue, [2]int{i, j-1})
                }
            }

            if j+1 < len(grid[i]) {
                if grid[i][j+1] == 1 {
                    grid[i][j+1] = 2
                    freshFruits--
                    tempQueue = append(tempQueue, [2]int{i, j+1})
                }
            }
            fmt.Println(queue, tempQueue)
            queue = queue[1:len(queue)]
        }

        fmt.Println(freshFruits)

        if freshFruits == 0 {
            return time
        }

        queue = append(queue, tempQueue...)
        fmt.Println(queue, freshFruits)
    }

    return -1
}

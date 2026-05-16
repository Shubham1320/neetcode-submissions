

func findRow(matrix [][]int, i int, j int, target int) int {

    if i > j {
        return -1
    }

    mid := int((i+j)/2)

    if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[0])-1]{
        return mid
    } else if target < matrix[mid][0] {
        return findRow(matrix, i, mid-1, target)
    } else if target > matrix[mid][len(matrix[0])-1] {
        return findRow(matrix, mid+1, j, target)
    }

    return -1
}

func binSearch(row []int, i int, j int, target int) bool {

    if i > j {
        return false
    }

    mid := int((i+j)/2)

    if target == row[mid] {
        return true
    } else if target > row[mid] {
        return binSearch(row, mid+1, j, target)
    } else {
        return binSearch(row, i, mid-1, target)
    }
}

func searchMatrix(matrix [][]int, target int) bool {

    row := findRow(matrix, 0, len(matrix)-1, target)
    if row != -1 {
        return binSearch(matrix[row], 0, len(matrix[row])-1, target)
    }
    return false
}
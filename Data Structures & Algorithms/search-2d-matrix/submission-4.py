class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        
        def find_row(matrix, i, j):

            print(i,j)

            if j < i:
                return -1

            mid = i + int((j-i)/2)

            print(mid)

            if target >= matrix[mid][0] and target <= matrix[mid][-1]:
                return mid
            elif target < matrix[mid][0]:
                return find_row(matrix, i, mid-1)
            elif target > matrix[mid][-1]:
                return find_row(matrix, mid+1, j)
                
        row_index = find_row(matrix, 0, len(matrix)-1)

        if row_index == -1:
            return False

        print(row_index)

        def find(row, i, j):
            if j < i:
                return False
            
            #print(i, j)

            mid = i + int((j-i)/2)

            #print(mid)

            if row[mid] == target:
                return True
            elif row[mid] < target:
                return find(row, mid+1, j)
            else:
                return find(row, i, mid-1)

        return find(matrix[row_index], 0, len(matrix[row_index]))

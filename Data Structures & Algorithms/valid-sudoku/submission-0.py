

1, 2, 3, 4, 5, 6, 7, 8 , 9








class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [{} for _ in range(9)]
        cols = [{} for _ in range(9)]
        sqrs = [{} for _ in range(9)]   
        
        i = 0
        j = 0

        while i != 9:
            while j != 9:
                if board[i][j] == '.':
                    j = j + 1
                    continue
                print(rows)
                print(cols)
                print(sqrs)
                try:
                    if rows[i][board[i][j]]:
                        return False
                except KeyError as ke:
                    rows[i][board[i][j]] = True
            
                try:
                    if cols[j][board[i][j]]:
                        return False
                except KeyError as ke:
                    cols[j][board[i][j]] = True
 
                try:
                    sqr_index = -1
                    if j < 3:
                        if i < 3:
                            sqr_index = 0
                        elif i < 6:
                            sqr_index = 1
                        elif i < 9:
                            sqr_index = 2
                    elif j < 6:
                        if i < 3:
                            sqr_index = 3
                        elif i < 6:
                            sqr_index = 4
                        elif i < 9:
                            sqr_index = 5
                    elif j < 9:
                        if i < 3:
                            sqr_index = 6
                        elif i < 6:
                            sqr_index = 7
                        elif i < 9:
                            sqr_index = 8
                    
                    if sqrs[sqr_index][board[i][j]]:
                        return False
                except KeyError as ke:
                    sqrs[sqr_index][board[i][j]] = True
                j = j + 1 
            j = 0
            i = i + 1
        return True



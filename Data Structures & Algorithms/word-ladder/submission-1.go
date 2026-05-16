


func ladderLength(beginWord string, endWord string, wordList []string) int {

    wordMap := make(map[string][]string)
    visited := make(map[string]bool)

    for _, word := range wordList {
        for i := 0; i < len(word); i++ {
            pattern := word[:i] + "*" + word[i+1:]
            
            if _, ok := wordMap[pattern]; !ok {
                wordMap[pattern] = []string{word}
            } else {
                wordMap[pattern] = append(wordMap[pattern], word)
            }
        }
    }

    fmt.Println(wordMap)

    var dfs func(string) int

    dfs = func(currentWord string) int {

        if currentWord == endWord {
            return 1
        }

        ans := math.MaxInt

        if _, ok := visited[currentWord]; ok {
            return ans
        }

        visited[currentWord] = true
        for i := 0; i < len(currentWord); i++ {
            pattern := currentWord[:i] + "*" + currentWord[i+1:]
            if wordArray, ok := wordMap[pattern]; ok {
                for _, word := range wordArray {
                    ans = min(ans, dfs(word))
                }
            }
        }
        delete(visited, currentWord)
        if ans == math.MaxInt {
            return ans
        }
        return ans + 1
    }

    ans := dfs(beginWord)
    if ans == math.MaxInt {
        return 0
    }
    return ans
}

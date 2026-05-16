var ans []string

func generate(i int, subString string, availableChars [][]string) {
    
    if len(subString) == len(availableChars) {
        ans = append(ans, subString)
        return
    }

    j := 0
    for j != len(availableChars[i]) {
        generate(i+1, subString + availableChars[i][j], availableChars)
        j++
    }

}

func letterCombinations(digits string) []string {

    ans = []string{}
    digitMap := make(map[byte][]string)

    digitMap = map[byte][]string {
        '2': []string{"a","b","c"},
        '3': []string{"d","e","f"},
        '4': []string{"g","h","i"},
        '5': []string{"j","k","l"},
        '6': []string{"m","n","o"},
        '7': []string{"p","q","r","s"},
        '8': []string{"t","u","v"},
        '9': []string{"w","x","y","z"},
    }

    var availableChars [][]string

    for _, val := range digits {
        availableChars = append(availableChars, digitMap[byte(val)])
    }

    if len(availableChars) == 0 {
        return []string{}
    } else if len(availableChars) == 1 {
        return availableChars[0]
    }

    generate(0, "", availableChars)

    return ans




}

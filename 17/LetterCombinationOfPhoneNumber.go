package  main 

var letterMap = map[byte]string{
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    var result []string
    backtrack(digits, 0, "", &result)
    return result
}

func backtrack(digits string, index int, combination string, result *[]string) {
    if index == len(digits) {
        *result = append(*result, combination)
        return
    }
    currentDigit := digits[index]
    letters := letterMap[currentDigit]
    for i := 0; i < len(letters); i++ {
        backtrack(digits, index+1, combination+string(letters[i]), result)
    }
}
func romanToInt(s string) int {
    result := 0
    for i := 0; i < len(s); i++ {
        switch  {
            case string(s[i]) == "I":
                if i == len(s) - 1 {
                    result += 1
                    continue
                }
                if string(s[i + 1]) == "V" || string(s[i + 1]) == "X"{
                    result -= 1
                }else {result += 1}
            case string(s[i]) == "V":
                result += 5
            case string(s[i]) == "X":
                if i == len(s) - 1 {
                    result += 10
                    continue
                }
                if string(s[i + 1]) == "L" || string(s[i + 1]) == "C"{
                    result -= 10
                }else {result += 10}
            case string(s[i]) == "L":
                result += 50
            case string(s[i]) == "C":
                if i == len(s) - 1 {
                    result += 100
                    continue
                }
                if string(s[i + 1]) == "D" || string(s[i + 1]) == "M"{
                    result -= 100
                }else {result += 100}
            case string(s[i]) == "D":
                result += 500
            case string(s[i]) == "M":
                result += 1000
            default:
                fmt.Println("存在無效的字母，只能包含I V X L C D M")
        }
    }
    return result
}
/*
其他人解法:
func romanToInt(s string) int {
    var romanMap = map[byte]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
    var result = romanMap[s[len(s)-1]]
    
    for i := len(s)-2; i >= 0; i-- {
        if romanMap[s[i]] < romanMap[s[i+1]] {
            result -= romanMap[s[i]]
        } else {
            result += romanMap[s[i]]
        }
	}
	return result
}
*/
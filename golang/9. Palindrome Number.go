func isPalindrome(x int) bool {
    str := strconv.Itoa(x)  //先將x轉為string以判斷有幾位數
    digits := len(str)      //x有幾位數
    if x < 0 {return false} //負數不為回文
    if x < 10 {return true} //個位數為回文(leetcode認定的) 
    if digits % 2 ==0 {     //偶數位數
        for i := 0; i <= (digits - 1)/2; i++ {
            if str[i] != str[digits - 1 - i] {
                return false
            }
        }
        return true
    } else {                //奇數位數
        for i := 0; i <= (digits/2) - 1; i++ {
            if str[i] != str[digits - 1 - i] {
                return false
            }
        }
        return true
    }
}
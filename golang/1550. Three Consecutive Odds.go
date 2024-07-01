func threeConsecutiveOdds(arr []int) bool {
    result := false
    for i := 0; i < len(arr) - 2; i++ {
        if arr[i] % 2 != 0 && arr[i+1] % 2 != 0 && arr[i+2] % 2 != 0{
            result = true
        }
    }

    return result
}
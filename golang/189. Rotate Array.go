func rotate(nums []int, k int) {
    n := len(nums)
    k = k % n // 確保 k 不超過陣列長度

    reverse(nums, 0, n - 1) //整個切片反轉
    reverse(nums, 0, k - 1) //反轉前k個元素
    reverse(nums, k, n - 1) //反轉前k個以外剩下的元素

}

func reverse(nums []int, start, end int) {  //將指定的index為start到end的元素進行反轉
    for true {
        if start < end {
            nums[start], nums[end] = nums[end], nums[start]     //golang特有的兩元素直接交換
            start += 1
            end -= 1
        } else {break}
    }
}

/*
範例解析
以範例1為例：
原始陣列：[1,2,3,4,5,6,7]
反轉整個陣列：[7,6,5,4,3,2,1]
反轉前3個元素：[5,6,7,4,3,2,1]
反轉剩下的元素：[5,6,7,1,2,3,4]
這樣就實現了將陣列右移k次的效果。這種方法的優點是只需O(n)的時間複雜度和O(1)的空間複雜度。
*/
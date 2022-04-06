func findDuplicates(nums []int) []int {
    m := make(map[int]int)
    arr := []int{}
    for _, v := range nums {
        m[v] += 1
    }

    for k, v := range m {
        if v == 2 {
            arr = append(arr, k)
        }
    }
    return arr
}

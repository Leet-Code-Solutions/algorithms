// func twoSum(numbers []int, target int) []int {
    
//     for l:=0; l < len(numbers); l++{
//         for r:=0; r < len(numbers); r++ {
//             if l != r && (numbers[l] + numbers[r] == target) {
//                 l++
//                 r++
//                 return []int{l, r}
//             }
           
//         }

//     }

//     return []int{}
// }

// 很明显这是一个升序数组
// 左右指针分别指向首尾
// 如果左+右 == target 直接返回
// 左+右 > target 说明右边的数大了，左移

func twoSum(numbers []int, target int) []int {
    l := 0
    r := len(numbers) - 1
    for l < r {
        sum := numbers[l]+numbers[r]
        if sum == target {
            l++
            r++
            return []int{l, r}  
        }

        if sum > target {
            r-- 
        } else {
            l++
        }

    }

    return []int{}
}
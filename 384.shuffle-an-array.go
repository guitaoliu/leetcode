type Solution struct {
    ori []int
}


func Constructor(nums []int) Solution {
    ori := make([]int, len(nums))
    copy(ori, nums)
    return Solution{ori: ori}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.ori
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    nums := make([]int, len(this.ori))
    copy(nums, this.ori)
    
    for i := 0; i < len(nums); i++ {
        idx := rand.Intn(len(nums) - i) + i
        nums[i], nums[idx] = nums[idx], nums[i]
    }
    
    return nums
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
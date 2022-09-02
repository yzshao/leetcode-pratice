package main

func main() {

}

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		remain := target - num
		if index, ok := indexMap[remain]; ok {
			return []int{i, index}
		} else {
			indexMap[num] = i
		}
	}
	return nil
}

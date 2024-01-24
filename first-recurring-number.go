// Time complexity O(n)
func findFirstRecurringCharacter(input []int) int {
  myMap := make(map[int]int)

  for _, v := range input {
    if _, ok := myMap[v]; ok {
      return v
    }
    myMap[v]++
  }

  return 0
}

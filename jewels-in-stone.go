func numJewelsInStones(jewels string, stones string) int {
    counter := make(map[rune]bool)
    sum := 0

    for _, value := range jewels {
        counter[value] = true
    }

    for _, value := range stones {
        if _, ok := counter[value]; ok {
            sum += 1
        }
    }
    
    return sum
}

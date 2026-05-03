func isAnagram(s string, t string) bool {
counts := make(map[rune]int)

for _, char := range s {
    counts[char]++
}

for _, char := range t {
    counts[char]--
    if counts[char] < 0 {
        return false
    }
}
return true
}

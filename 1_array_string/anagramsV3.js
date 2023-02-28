const anagrams = (s1, s2) => {
    const map = {}
    if (s1.length !== s2.length) return false
    for (let i = 0; i < s1.length; i++) {
        if (map[s1[i]]) {
            map[s1[i]]++
        } else {
            map[s1[i]] = 1
        }
    }
    for (let i = 0; i < s2.length; i++) {
        if (!map[s2[i]]) {
            return false
        } else {
            map[s2[i]]--
        }
    }
    return true
}   
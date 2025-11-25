package main

import "fmt"

func CountInts(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	return counts
}

func CountWorlds(nums []string) map[string]int {
	counts := make(map[string]int)
	for _, num := range nums {
		counts[num]++
	}
	return counts
}

func IsAnagram(a, b string) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	map1 := make(map[string]int)
	for i, num := range a {
		map1[string(num)] = i
	}
	map2 := make(map[string]int)
	for i, num := range b {
		map2[string(num)] = i
	}
	for key := range map1 {
		if _, ok := map2[key]; !ok {
			return false
		}
	}
	return true
}

func Invert(m map[string]int) map[int]string {
	n := make(map[int]string)
	for key, value := range m {
		n[value] = key
	}
	return n
}

func MergeMaps(dst, src map[string]int) map[string]int {
	for key, value := range src {
		dst[key] = dst[key] + value
	}
	return dst
}

type User struct {
	Name string
	City string
}

func GroupByCity(users []User) map[string][]User {
	result := make(map[string][]User)
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
		result[users[i].City] = append(result[users[i].City], users[i])
	}
	return result
}

func main() {
	//1
	nums := []int{2, 3, 4, 5, 3, 2, 2, 4, 5, 6, 7, 5, 3, 2, 2, 4, 5}
	counts := CountInts(nums)
	for num, count := range counts {
		fmt.Println(num, count)
	}

	//2
	str := "Эх, если бы да кабы, во рту росли грибы, то был бы не рот, а целый огород"
	var tmp_str string
	var worlds = []string{}
	for _, word := range str {
		if string(word) != " " && string(word) != "," {
			tmp_str += string(word)
		} else {
			if tmp_str != "" {
				worlds = append(worlds, tmp_str)
				tmp_str = ""
			}
		}
	}
	counts1 := CountWorlds(worlds)
	for num, count := range counts1 {
		fmt.Println(num, count)
	}

	//3
	a := "кабан"
	b := "банка"
	fmt.Println(IsAnagram(a, b))

	//4
	forward := map[string]int{"a": 1, "b": 2}
	back := Invert(forward)
	fmt.Println(forward)
	fmt.Println(back)

	//5
	src := map[string]int{"f": 2, "d": 4, "a": 5}
	dst := map[string]int{"a": 3, "c": 5, "f": 3}
	dst = MergeMaps(dst, src)
	fmt.Println(src)
	fmt.Println(dst)

	//6
	tmp_user := []User{
		{"User1", "Almaty"},
		{"User2", "Almaty"},
		{"User3", "Astana"},
		{"User4", "Astana"},
	}
	q := GroupByCity(tmp_user)
	for key, value := range q {
		fmt.Println(key, value)
	}

}

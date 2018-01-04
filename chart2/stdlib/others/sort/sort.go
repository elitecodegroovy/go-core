package main

import (
	"fmt"
	"sort"
)

func SortSlice(){
	a := []int{1,-10, 10, 11, 90,18}
	fmt.Println("before a :", a)
	sort.Ints(a)
	fmt.Println("after a :", a)

	strSlice := []string{"Jamaica","Estonia","Indonesia","Hong Kong", "China"} // unsorted
	fmt.Println(" BEFORE sort:",strSlice)
	sort.Strings(strSlice)
	fmt.Println(" AFTER  sort:",strSlice)

}

type User struct {
	Name string
	Age int
	Profession string
}

type UserAge []User
func (s UserAge) Len() int {
	return len(s)
}
func (s UserAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s UserAge) Less(i, j int) bool {
	if s[i].Age == s[j].Age{
		return len(s[i].Name) < len(s[j].Name)
	}
	return s[i].Age <s[j].Age
}


func sortCustomUser(){
	users := []User{
		User{"李明", 10, "学生"},
		User{"刘希", 18, "学生"},
		User{"马华", 8, "学生"},
		User{"董明", 28, "教师"},
		User{"李克勤", 38, "教授"},
		User{"马克", 38, "高级工程师"},
	}
	fmt.Println("before sort users: ", users)
	sort.Sort(UserAge(users))
	fmt.Println("after sort users: ", users)
}
func main(){
	SortSlice()
	sortCustomUser()
}

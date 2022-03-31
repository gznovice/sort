package main

import(
	"fmt"
	"sort"
)

type myStruct struct{
	name string
	value int
}

type ByNameMyStruct []myStruct

func (s ByNameMyStruct)Len() int{
	return len(s)
}

func (s ByNameMyStruct)Less(i, j int) bool{
	return s[i].name < s[j].name
}

func (s ByNameMyStruct)Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

type ByValueMyStruct []myStruct

func (s ByValueMyStruct)Len() int{
	return len(s)
}

func (s ByValueMyStruct)Less(i, j int) bool{
	return s[i].value < s[j].value
}

func (s ByValueMyStruct)Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func main(){
	var myStructSlice []myStruct = []myStruct{{"tom", 65},{"ben", 70}, {"qiang", 80}, {"hit", 60}, {"cal", 70}}
	
	fmt.Println("Before sorting:", myStructSlice)
	
	sort.Sort(ByNameMyStruct(myStructSlice))
	
	fmt.Println("After sorting1:", myStructSlice)
	
	sort.Sort(ByValueMyStruct(myStructSlice))
	
	fmt.Println("After sorting2:", myStructSlice)
}
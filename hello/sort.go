package main

import(
	"fmt"
    "sort"
)

type Squeue []int

func (s Squeue) Len()int {
	return len(s)
}

func (s Squeue) Less(i,j int)bool {
    return s[i] > s[j]
}

func (s Squeue)Swap(i,j int){
    s[i],s[j] = s[j],s[i]
}

func (s Squeue) String() string{
    sort.Sort(s)
    res := "["
    for k,v := range s {
	if k > 0 {
		res += " "
	}
        res += fmt.Sprint(v)
    }
    res += "]"
    return res
}

func main(){
    s := Squeue{1,3,2,4}
    sort.Sort(s)
    fmt.Println(s)
}

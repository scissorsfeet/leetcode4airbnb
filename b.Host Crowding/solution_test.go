package b_Host_Crowding

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	results := []string{
		"1,28,300.1,SanFrancisco",
		"1,5,209.1,SanFrancisco",
		"2,7,208.1,SanFrancisco",
		"3,8,207.1,SanFrancisco",
		"1,10,206.1,Oakland",
		"1,16,205.1,SanFrancisco",
		"6,29,204.1,SanFrancisco",
		"7,20,203.1,SanFrancisco",
	}
	res := hostCrowding(3, results)
	for _, v := range res {
		fmt.Println(v)
	}
}

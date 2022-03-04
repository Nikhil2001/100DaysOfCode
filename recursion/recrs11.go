/*
finding nth generation kth child gender
a family follows child hierachy as shown below

                       m

	    	m                                  f   // 2nd generation has 1st child as male 2nd as female


     m           f	                    f         m // 3rd generation 3rd child is female

	 ..

	 nth generation kth child ?

*/

package main

import "fmt"

func family(output, input []string, FirstGeneration, generation, k int) {
	for _, v := range input {
		if v == "M" {
			output = append(output, "M", "F")
		} else if v == "F" {
			output = append(output, "F", "M")
		}
	}
	FirstGeneration++
	if FirstGeneration == generation {
		fmt.Println(output)
		fmt.Println(output[k-1]) // kth child in k-1 index
		return
	}
	family([]string{}, output, FirstGeneration, generation, k)
}

func main() {
	output := []string{}
	input := []string{"M"}
	FirstGeneration, generation := 1, 4
	k := 4
	family(output, input, FirstGeneration, generation, k)

}

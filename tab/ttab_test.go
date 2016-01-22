package tab

import "fmt"

func ExampleCrossroadsCream() {
	motifStr := "[1:5,-,3:7,-,3:5,-,3:7,-,-,3:5,3:7,-,1:8,-,1:5]"
	motif := ParseMotif(motifStr + motifStr)
	fmt.Println(motif.String())
	// Output:
	//e||---------------|---------------||
	//B||---------------|---------------||
	//G||---------------|---------------||
	//D||--7-5-7--57----|--7-5-7--57----||
	//A||---------------|---------------||
	//E||5-----------8-5|5-----------8-5||
}

func ExampleGToDChords() {
	motifStr := "[-,-,1:3|2:2|3:0|4:0|5:3|6:3,-,-][-,-,3:0|4:2|5:3|6:2,-,-]"
	motif := ParseMotif(motifStr)
	fmt.Println(motif.String())
	// Output:
	//e||--3--|--2--||
	//B||--3--|--3--||
	//G||--0--|--2--||
	//D||--0--|--0--||
	//A||--2--|-----||
	//E||--3--|-----||
}
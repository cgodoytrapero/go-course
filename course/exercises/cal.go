package exercises

import "fmt"

type content int

const (
	oppo content = iota - 1
	blank
	computer
)

type Board [9]content

func (b Board) String() string {
	output := ""
	for irow := 0; irow < 3; irow++ {
		output += fmt.Sprintf("%v | %v | %v \n", b[0+irow*3], b[1+irow*3], b[2+irow*3])

		if irow < 2 {
			output += fmt.Sprintf("\n")
		}
	}

	return output
}

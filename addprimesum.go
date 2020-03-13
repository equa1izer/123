import (
	"fmt"
	"strconv"
)

func main() {
	args := []string{"main", "59"}
	//args := []string{"main", "5", "7"}
	//args := []string{"main", "A"}
	if len(args) != 2 {
		fmt.Println(0)
		return
	}
	a, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(0)
		return
	}
	//Prime numbers starts at 2
	if a < 2 {
		fmt.Println(0)
		return
	}
	stack := []int{}
	for i := 2; i <= a; i++ {
		if isPrime222(i) {
			stack = append(stack, i)
		}
	}
	if len(stack) == 0 {
		fmt.Println(0)
		return
	}
	n := 0
	for i := 0; i < len(stack); i++ {
		n = n + stack[i]
	}
	fmt.Println(stack)
	fmt.Println(n)
}
func isPrime222(nbr int) bool {
	for i := 2; i < nbr; i++ {
		if nbr%i == 0 {
			return false
		}
	}
	return true
}

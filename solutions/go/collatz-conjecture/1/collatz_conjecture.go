package collatzconjecture
import "errors"
const MaxSteps = 10000

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("only positive numbers")
    }
    steps := 0
	for ; steps < MaxSteps; {
        switch {
            case n == 1:
            	return steps, nil
            case n % 2 == 0:
            	n /= 2
            	steps++
            default:
            	n = n *3 +1
            	steps++
    	}
    }
    return 0, errors.New("to many steps")
}

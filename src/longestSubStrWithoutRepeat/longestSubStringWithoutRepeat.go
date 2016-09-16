package substring

import (
	"fmt"
)
/*
func hasRepeat(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		for _, value := range s {
			if value == rune(s[i]) {
				count++
			}
		}

		if count > 1 {
			return true
		} else {
			count = 0
		}
	}

	return false
}
*/

func hasRepeat(counter []int, c byte) bool {
	counter[c]++
	//fmt.Printf("counter[%v]: %d\n", c, counter[c])
	if counter[c] > 1 {
		//fmt.Printf("counter[%v]: %d\n", c, counter[c])
		return true
	} else {
		return false
	}
}

func subLeftString(s string) (sub string) {
	sub = s[:len(s)-1]
	return sub
}

func subRightString(s string) (sub string) {
	sub = s[1:len(s)]
	return sub
}

/*
func allSame(s string) bool {
	c := rune(s[0])
	for i := 1; i < len(s); i++ {
		if c != rune(s[i]) {
			return false
		}
	}

	return true
}
*/

func allSame(s string) bool {
	counter := make([]int, 128)
	c := s[0]
	for i := 0; i < len(s); i++ {
		c := s[i]
		counter[c]++
	}

	if counter[c] == len(s) {
		return true
	} else {
		return false
	}

}
func isDup(s string, r []byte) bool {
	for i := 0; i < len(s); i ++ {
		if byte(s[i]) == r[0] {
			return true
		}
	}
	fmt.Println(r)
	return false

}

func lengthOfLongestSubstring(s string) int {
	max := 0

	if s == "" {
		return 0
	}

	if allSame(s) {
		return 1
	} else {
		for i := 0; i < len(s) - 1; i++ {
			count := 0
			counter := make([]int, 128)
			for j := i; j < len(s); j++ {
				r := hasRepeat(counter, s[j])
				fmt.Printf("s: %s, repeat: %v, count: %d,    i: %d, j: %d\n", s[i:j], r, count, i, j)
				if r {
					break
				} else {
					count++
				}
			}
			if max <= count {
				max = count
				fmt.Printf("max: %d\n", max)
			}


		}
		fmt.Printf("s: %s, result: %d\n", s, max)
		return max
	}


}


/*
func lengthOfLongestSubstring(s string) int {
	max := 0
	count := 0

	if s == "" {
		return 0
	}

	if allSame(s) {
		return 1
	} else {
		for i := 0; i < len(s); i++ {
			for j := i + 1  ; j <= len(s); j++ {
				count++
				r := hasRepeat(s[i:j], s[j])
				fmt.Printf("s: %s, repeat: %v, count: %d\n", s[i:j], r, count)
				if r {
					count--
					break
				}
			}

			if max < count {
				max = count
				fmt.Printf("max: %d\n", max)
			}

			count = 0
		}
		fmt.Printf("s: %s, result: %d\n", s, max)
		return max
	}

}
*/
// wrong answer
/*
//TODO: merge hasRepeat with allSame
func hasRepeat(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		for _, value := range s {
			if value == rune(s[i]) {
				count++
			}
		}

		if count > 1 {
			return true
		} else {
			count = 0
		}
	}

	return false
}

func allSame(s string) bool {
	c := rune(s[0])
	for i := 1; i < len(s); i++ {
		if c != rune(s[i]) {
			return false
		}
	}

	return true
}

func subLeftString(s string, limit int) (sub string) {
	sub = s[:len(s)-limit]
	return sub
}

func subRightString(s string, limit int) (sub string) {
	sub = s[limit:len(s)]
	return sub
}
/*
func upLimit(s string) (limit int) {

	count := 0
	max := 0
/*
	for i := 0; i < len(s); i++ {
		for _, value := range s {
			if value == rune(s[i]) {
				count++
			}
		}

		if count > max {
			max = count
			count = 0
		} else {
			count = 0
		}

	}
	var i int
	i = 0
	for {
		for {
			if s[i] == s[i+1] {
				count++
			}
			i++
		}
		if count > max {
			max = count
			i += count
			count = 0
		} else {
			count = 0
		}
		if i == len(s) - 1 {
			break
		}
	}


	return max

}

var max int

func lengthOfLongestSubstring(s string) int {
	var limit int
	var size1, size2 int

	limit = upLimit(s) - 1
	fmt.Printf("s=%s, offset=%d", s, limit)

	if limit == 0 {
		return 1
	} else {

		fmt.Println("-------------")

		leftsub := subLeftString(s, limit)
		if allSame(leftsub) {
			size1 = 1
		} else {
			size1 = lengthOfLongestSubstring(leftsub)
		}
		fmt.Printf("s=%s, size1=%d\n", leftsub, len(leftsub))

		rightsub := subRightString(s, limit)
		if allSame(rightsub) {
			size2 = 1
		} else {
			size2 = lengthOfLongestSubstring(rightsub)
		}
		fmt.Printf("s=%s, size2=%d\n", rightsub, len(rightsub))

		fmt.Println("-------------")
		if size1 > size2 {
			return size1
		} else {
			return size2
		}

	}
	/*
		if allSame(s) {
			return 1
		}else if hasRepeat(s) {

			/*
			if size1 > max && size1 >= size2 {
				max = size1
				return size1
			} else {
				return max
			}
			if  size2 > max && size2 > size1{
				max = size2
				return size2
			} else {
				return max
			}

		} else {
			fmt.Println("-------------")
			fmt.Printf("s=%s, len=%d\n", s, len(s))
			fmt.Println("-------------")
			if len(s) > max {
				max = len(s)
				return max
			} else {
				return len(s)
			}
		}
}
*/

// Too much time
/*


func hasRepeat(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		for _, value := range s {
			if value == rune(s[i]) {
				count++
			}
		}

		if count > 1 {
			return true
		} else {
			count = 0
		}
	}

	return false
}

func subLeftString(s string) (sub string) {
	sub = s[:len(s)-1]
	return sub
}

func subRightString(s string) (sub string) {
	sub = s[1:len(s)]
	return sub
}

var max int

func lengthOfLongestSubstring(s string) int {
	var size1, size2 int

	if hasRepeat(s) {

		leftsub := subLeftString(s)
		size1 = lengthOfLongestSubstring(leftsub)

		rightsub := subRightString(s)
		size2 = lengthOfLongestSubstring(rightsub)

		if size1 > size2 {
			return size1
		} else {
			return size2
		}

	} else {
			return len(s)
	}

}

*/

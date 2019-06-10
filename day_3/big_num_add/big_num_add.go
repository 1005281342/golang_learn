package big_num_add

import "fmt"

func BigNumberAdd(s1, s2 string) (res string) {

	l1, l2 := len(s1)-1, len(s2)-1
	left := 0

	for l1 > -1 && l2 > -1 {
		c1 := s1[l1] - '0'
		c2 := s2[l2] - '0'
		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", c3, res)
		l1--
		l2--
	}

	for l1 > -1 {
		c1 := s1[l1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", c3, res)
		l1--
	}

	for l2 > -1 {

		c2 := s2[l2] - '0'
		sum := int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", c3, res)
		l2--
	}

	if left == 1 {
		res = fmt.Sprint("1%s", res)
	}

	return
}

func Multi(str1, str2 string) (result string) {

	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}

	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'

		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'

		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c1 := str2[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}

	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return
}

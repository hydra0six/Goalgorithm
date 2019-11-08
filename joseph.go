package main

//约瑟夫环问题递归解决
func Joseph(n, m int) (s int) {
	if n == 1 {
		return n
	}

	s = (Joseph(n-1, m)+m-1)%n + 1
	//fmt.Println(s)
	return s
}

func Joseph2(n, m, i int) (s int) {
	if i == 1 {
		return (n + m - 1) % n
	} else {
		s = (Joseph2(n-1, m, i-1) + m) % n
		//fmt.Println(s)
		return s
	}
}

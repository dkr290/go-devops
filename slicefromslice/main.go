package main

import "fmt"

func main() {
	l1, l2, l3 := linked()
	fmt.Println("Linked   ", l1, l2, l3)
	nl1, nl2 := noLinks()
	fmt.Println("No linked   ", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap link", cl1, cl2)
	cn1, cn2 := capNoLink()
	fmt.Println("Cap no link ", cn1, cn2)
	copy1, copy2, copied := copyNoLink()
	fmt.Println("Copy no link:", copy1, copy2, copied)
	a1, a2 := appendNolink()
	fmt.Println("Append No link ", a1, a2)

}

func linked() (int, int, int) {
	//slice with some
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLinks() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]

}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {

	s1 := make([]int, 5, 10)
	for i := 1; i < 6; i++ {
		s1[i-1] = i
	}
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNolink() (int, int) {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := append([]int{}, a1...)
	a1[3] = 99
	return a1[3], a2[3]
}

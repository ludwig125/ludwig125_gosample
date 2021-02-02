package main

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	// tests := map[string]struct {
	// 	str1 string
	// 	str2 string
	// 	want bool
	// }{
	// 	"1": {
	// 		str1: "waterbottle",
	// 		str2: "erbottlewat",
	// 		want: true,
	// 	},
	// 	// "2": {
	// 	// 	str1: "Alfa Bravo Charlie Delta Echo Foxtrot Golf",
	// 	// 	str2: "lta Echo Foxtrot GolfAlfa Bravo Charlie De",
	// 	// 	want: true,
	// 	// },
	// 	// "3": {
	// 	// 	str1: "Alfa Bravo Charlie Delta Echo Foxtrot Golf",
	// 	// 	str2: " lta Echo Foxtrot GolfAlfa Bravo Charlie De",
	// 	// 	want: false,
	// 	// },
	// }
	// for name, tt := range tests {
	// 	t.Run(name, func(t *testing.T) {
	// 		l := singleLinkedList{}
	// 		fmt.Println(l)

	// 	})
	// }
	l := singleLinkedList{value: 10}
	l.print()

	fmt.Println("---")
	l.add(11)
	l.print()

	fmt.Println("---")
	l.add(12)
	l.print()

	fmt.Println("---")
	l.add(13)
	l.print()

	fmt.Println("---")
	fmt.Println(l.search(10, 0))
	fmt.Println(l.search(11, 0))
	fmt.Println(l.search(12, 0))
	fmt.Println(l.search(13, 0))
	fmt.Println(l.search(14, 0))

	fmt.Println("---")
	l.delete(13)
	// fmt.Println("this ", l.value)
	l.print()

}

type singleLinkedList struct {
	next  *singleLinkedList
	value int
}

func newSingleLinkedList(is ...int) *singleLinkedList {
	l := &singleLinkedList{value: is[0]}
	for i := 1; i < len(is); i++ {
		l.add(is[i])
	}
	return l
}

func (s *singleLinkedList) print() {
	fmt.Println(s.value)
	if s.next != nil {
		s.next.print()
	}
}

func (s *singleLinkedList) add(target int) {
	if s.next == nil {
		s.next = &singleLinkedList{value: target}
		return
	}
	s.next.add(target)
}

func (s *singleLinkedList) search(target, cnt int) int {
	if s.value == target {
		return cnt
	}
	if s.next == nil {
		return -1
	}
	return s.next.search(target, cnt+1)
}

func (s *singleLinkedList) delete(target int) {
	if s.value == target {
		if s.next != nil {
			if s.next.next != nil {
				*s = singleLinkedList{value: s.next.value, next: s.next.next}
				return
			}
			*s = singleLinkedList{value: s.next.value}
			return
		}

	}
	if s.next != nil {
		// 次が終端だったら次の次はないはず
		if s.next.next == nil {
			// 終端が対象の値と一致したらlistを現在のNodeで最終にする
			if s.next.value == target {
				*s = singleLinkedList{value: s.value}
				return
			}
		}
	}
	s.next.delete(target)
}

// func (s *singleLinkedList) delete(target int) {
// 	// fmt.Println("here0", s.value, target)

// 	if s.value == target {
// 		if s.next != nil {
// 			if s.next.next != nil {
// 				// fmt.Printf("%v %#v\n", s.next.value, *(s.next.next))
// 				// fmt.Println("here")
// 				*s = singleLinkedList{value: s.next.value, next: s.next.next}
// 				return
// 			}
// 			*s = singleLinkedList{value: s.next.value}
// 			return
// 		}

// 	}
// 	if s.next != nil {
// 		// 次が終端だったら次の次はないはず
// 		if s.next.next == nil {
// 			// 終端が対象の値と一致したらlistを現在のNodeで最終にする
// 			if s.next.value == target {
// 				*s = singleLinkedList{value: s.value}
// 				return
// 			}
// 		}
// 	}
// 	s.next.delete(target)
// }

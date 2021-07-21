package main

import "fmt"

type rectangle struct {
	a int
	b int
}

func S(rec rectangle) (s int) {
	s = rec.a + rec.b
	return s
	
}

func P(rec rectangle) (p int) {
	p = 2 * (rec.a + rec.b)
	return p

}

func main() {
  answer := rectangle{
	  a: 5,
	  b: 5,
  }
fmt.Println(answer)
}

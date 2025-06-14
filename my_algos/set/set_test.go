package set

import (
	"fmt"
	"testing"
)

func Test_Set(t *testing.T) {
	setA := NewSet("alaska", "ny")
	setB := NewSet("ny", "kansas")
	fmt.Println(intersection(setA, setB).GetElements()) // ny
	fmt.Println(difference(setA, setB).GetElements())   // alaska
	fmt.Println(union(setA, setB).GetElements())        // alaska, ny, kansas

	setC := NewSet("california", "ny")
	fmt.Println(intersectionMulti(setA, setB, setC).GetElements()) // ny
}

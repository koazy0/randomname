package randomname

import (
	"fmt"
	"testing"
)

func TestGenerateNameAndReturnFirst(t *testing.T) {
	name, first := GenerateNameAndReturnFirst()
	fmt.Println(name, first)

	name, first = GenerateNameAndReturnFirst()
	fmt.Println(name, first)

	name, first = GenerateNameAndReturnFirst()
	fmt.Println(name, first)
}

package contrib

import (
	"fmt"

	"github.com/fako1024/effuncert"
)

// Dummy just is a dummy function
func Dummy() {
	fmt.Println("Hello from a contrib plugin:", effuncert.New(1, 100).Mode)
}

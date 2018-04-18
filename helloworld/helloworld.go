package helloworld

import (
	"fmt"
)

func Hello(world string) string {
	return fmt.Sprintf("hello %s", world)
}

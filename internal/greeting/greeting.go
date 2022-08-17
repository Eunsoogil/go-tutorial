package greeting

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %s", name)
	fmt.Println(message)
	return message
}

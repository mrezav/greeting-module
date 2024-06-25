package greetingmodule

import "fmt"

func SayHello(name string) string {
	result := fmt.Sprintf("Hello %v, selamat datang di go programming", name)
	return result
}

func SayGreeting() string {
	return "Selamat Belajar Golang Module"
}

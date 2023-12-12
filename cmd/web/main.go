package main

import (
	"fmt"
	"net/http"

	"github.com/openmymai/modern_webapp_in_golang/pkg/handlers"
)

const portNumber = ":8080"

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 10.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by 0")
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("Cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, world!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // add _ , ignore if error
}

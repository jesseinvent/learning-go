package main;

import (
	"fmt"
	// "errors"
	// "math"
);

// Struct
type person struct {
	name string;
	age int;
}

func main()  {
	// Variables
	x := 8;
	if (x > 6) {
		fmt.Println("More than x");
	} else if (x < 6) {
		fmt.Println("Less than x");
	} else {
		fmt.Println("Equal to x");
	}

	 arr := []int {1,2,4,2,4};
	 arr = append(arr, 13);
	fmt.Println(arr);

	// Maps
	vertices := make(map[string]int);

	vertices["triangle"] = 2;
	vertices["square"] = 3;
	vertices["rectanlgle"] = 4;

	fmt.Println(vertices);

	// Loops

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i);
	}

	// while loop
	j := 0;
	for j < 5 {
		fmt.Println(j);
		j++;
	}

	// Range
	arr2 := []string{"a", "b",  "c"}

	for index, value := range arr2 {
		fmt.Println("index:", index, "value:", value);
	}

	// Functions
	result := sum(2,4);
	fmt.Println(result);

	// var value float64 = 16;
	// result, err := sqrt(-16);

	// if (err != nil) {
	// 	fmt.Println(err);
	// } else {
	// 	fmt.Println(result)
	// }

	p := person{name: "Jake", age:23};

	fmt.Println(p.name);

	// Pointers
	z := 7;
	inc(&z);
	fmt.Println(z);
} 

func inc(x *int) {
	// Increments value at memory address
	*x++;
}

func sum(x int, y int) int {
	return x + y;
}

// Functions returning errors
// func sqrt(x float64) (float64, error) {
// 	if (x < 0) {
// 		return 0, errors.New("Undefined for negative numbers");
// 	}

// 	return math.Sqrt(x), nil;
// }
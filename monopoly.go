//This is a monopololy program to help track and handle the monolpoly pieces for the safeway monopoly giveaway.

//Author: Allen Byerly
//Name: monopoly

//The Main Package Definition
package main

//The Imported Libraries
import (
  //import format library
  "fmt"
  //import time library
  "time"
  //import the math library
  "math"
  //import random number functions from the math library
  "math/rand"
  //import string conversion
  "strconv"
)

//added a basic addition function to perform addition of two numbers
//Shortend parameter inputs to only one int declaration
//note: this is similar to how C declares variables of the same type
func add(x, y int) int {
  return x + y
}

//added a swap function to swap two variables for one another
func swap(x, y string) (string, string) {
	return y, x
}

//Main Package Function Definition
func main() {
  //Seed the random number generator based on the current time (UTC Unix Nano)
  // Use
  rand.Seed(time.Now().UTC().UnixNano())
  //Print a Name Prompt
  fmt.Println("Hello World!")
  //Print a Description Prompt
  fmt.Println("This is my Go Playground!")
  //Print a current time Prompt
  fmt.Println("The time is: ", time.Now())
  //Print a random number prompt
  //note: the random number will always return the same number until it is seeded
  fmt.Println("A random number is: ", rand.Intn(100))
  //Print a problems Promt based off the square root of the random number
  //Display the number as a defualt variable
  fmt.Printf("Now you have %v problems.\n", int(math.Sqrt(rand.Float64()*100)))
  //Print a prompt that displays Pi
  fmt.Printf("Pi is: %v\n", math.Pi)
  //add some of my favorite numbers together
  //use variables this time
  favorite_numbers := []int {
    42,
    22,
    13,
    222,
  }
  fmt.Println(favorite_numbers, "are some of my favorite numbers")
  //Print a prompt with the solution to adding some numbers together using the add function
  fmt.Println(favorite_numbers[0], "+", favorite_numbers[2], "=", add(favorite_numbers[0], favorite_numbers[2]))
  // put these numbers active numbers into an array
  active_variables := []string {
    "hello",
    "world",
  }
  a, b := swap(active_variables[0], active_variables[1])
  swapped_variables := []string {
    a,
    b,
  }

  // Use a swap funciton  to swap to variables
  fmt.Println("now lets swap strings", active_variables, "for", swapped_variables)
  x := strconv.Itoa(favorite_numbers[0])
  y := strconv.Itoa(favorite_numbers[2])
  active_variables = []string {x,y}
  x,y = swap(x, y)
  swapped_variables = []string {x,y}
  fmt.Println("now lets swap ints", active_variables, "for", swapped_variables)
}

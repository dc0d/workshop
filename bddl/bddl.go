// Package bddl is just a throw-away playground
package bddl

import "fmt"

func Describe(desc string, actions ...func()) {
	fmt.Printf("%s:\n", desc)
	Doing(actions...)
}

func Given(desc string, actions ...func()) {
	fmt.Println("  given", desc)
	Doing(actions...)
}

func When(desc string, actions ...func()) {
	fmt.Println("  when", desc)
	Doing(actions...)
}

func And(desc string, actions ...func()) {
	fmt.Println("  and", desc)
	Doing(actions...)
}

func Then(desc string, actions ...func()) {
	fmt.Println("  then", desc)
	Doing(actions...)
}

func Doing(actions ...func()) {
	for _, act := range actions {
		act()
	}
}

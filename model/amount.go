package model

import "fmt"

type Amount int

func (a Amount) String() string {
	if a == 0 {
		return ""
	}
	return fmt.Sprintf("%.2f", float64(a))
}

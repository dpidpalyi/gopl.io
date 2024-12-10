package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%.2fkg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%.2flb", p)
}

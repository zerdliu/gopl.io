package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
    OnePound2Kilograms Kilogram = 0.45359237
)

func (p Pound) toKilogram() Kilogram {
    return Kilogram(p) * OnePound2Kilograms
}

func (k Kilogram) toPound() Pound {
    return Pound(k / OnePound2Kilograms)
}

func (p Pound) String() string {
    return fmt.Sprintf("%g Pound", p)
}

func (k Kilogram) String() string {
    return fmt.Sprintf("%g Kilogram", k)
}
package lengthconv

import "fmt"

type Foot float64
type Meter float64

const (
    OneFoot2Meter float64 = 0.3048
)

func FToM (f Foot) Meter {
    return Meter(float64(f) * OneFoot2Meter)
}

func MToF (m Meter) Foot {
    return Foot(float64(m) / OneFoot2Meter)
}

func (f Foot) String() string {
    return fmt.Sprintf("%g Feet", f)
}

func (m Meter) String() string {
    return fmt.Sprintf("%g Meters", m)
}
package tempconv

import "fmt"

type K float64

const (
    AbsoluteZeroK K = 0
)

func KToC(k K) Celsius { return Celsius(k + K(-AbsoluteZeroC))}


func (k K) String() string    { return fmt.Sprintf("%g°K", k) }


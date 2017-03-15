package miscellany


import(
    "testing"
    "fmt"
    "math"
    "time"
)

func TestModf(t *testing.T) {
    
    var f float64 = 100.56
    var i float64
    var d float64

    i, d = math.Modf(f)

    fmt.Printf("Parte entera %f - Parte decimal %f\n", i, d)
}

func TestUnix(t *testing.T) {
    var now = time.Now().Unix()

    fmt.Printf("Now %d\n", now);
}

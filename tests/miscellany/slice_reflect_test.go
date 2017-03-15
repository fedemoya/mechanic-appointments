package miscellany


import(
    "testing"
    "reflect"
    "fmt"
)

func TestSliceReflect(t *testing.T) {
    var intSlice []int = make([]int, 5)

    sliceValue := reflect.ValueOf(intSlice)

    fmt.Println("Name " + sliceValue.Slice(1, 1).Kind().String())
}

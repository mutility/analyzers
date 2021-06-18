package a

import (
	"a/pkg"
	"fmt"
)

var (
	v1 = "v1"
	v2 = "v2"
	v3 = "v3"

	ar  = []string{"a", "b", "c"}
	pv1 = &v1
)

const (
	c1 = "c1"
	c2 = "c2"
	c3 = "c3"
)

func bad() {
	// note: plusses, square brackets, other regex chars replaced with .

	fmt.Sprintf(v1)           // want "variable `v1` used for fmt.Sprintf format parameter"
	passthrough(v1 + v2)      // want "variable `v1 . v2` used for passthrough format parameter"
	fmt.Sprintf(v1 + v2 + v3) // want "variable `v1 . v2 . v3` used for fmt.Sprintf format parameter"
	fmt.Sprintf(v1 + c1)      // want "variable `v1` used for fmt.Sprintf format parameter"
	fmt.Sprintf((v1 + c1))    // want "non-constant expression used for fmt.Sprintf format parameter"
	fmt.Sprintf(ar[0])        // want "variable `ar.0.` used for fmt.Sprintf format parameter"
	fmt.Sprintf(*pv1)         // want "variable `.pv1` used for fmt.Sprintf format parameter"
	fmt.Sprintf(v3[:2])       // want "variable `v3` used for fmt.Sprintf format parameter"
	fmt.Sprintf(v3[2:])       // want "variable `v3` used for fmt.Sprintf format parameter"
	fmt.Sprintf(v3[0:2])      // want "variable `v3` used for fmt.Sprintf format parameter"
	fmt.Sprintf(pkg.V)        // want "variable `pkg.V` used for fmt.Sprintf format parameter"
	lv := "lv"
	fmt.Sprintf(lv) // want "variable `lv` used for fmt.Sprintf format parameter"
	lookup := map[bool]string{false: "false", true: "true"}
	fmt.Sprintf(lookup[false])                                        // want "variable `lookup.false.` used for fmt.Sprintf format parameter"
	fmt.Sprintf(map[bool]string{false: "false", true: "true"}[false]) // want "non-constant expression used for fmt.Sprintf format parameter"
	complicated(v1, v2, v3)                                           // want "variable `v3` used for complicated format parameter"
}

func goodf(format string, args ...interface{}) {
	fmt.Sprint()
	fmt.Sprintf(c1)
	fmt.Sprintf(c1 + c2)
	fmt.Sprintf((c1 + c2))
	fmt.Sprintf(c1 + c2 + c3)
	fmt.Sprintf("abc")
	fmt.Sprintf(string("abc"))
	fmt.Sprintf(string(pkg.String("abc")))
	fmt.Sprintf("abc" + c3)
	fmt.Sprintf("abc"[:2])
	fmt.Sprintf("abc"[2:])
	fmt.Sprintf("abc"[0:2])
	fmt.Sprintf(c3[:2])
	fmt.Sprintf(c3[2:])
	fmt.Sprintf(c3[0:2])
	const lc = "lc"
	fmt.Sprintf(lc)
	fmt.Sprintf(pkg.C)
	complicated(v1, v2, c3)
}

func passthrough(format string, args ...interface{}) {
	// wrappers are allowed
	fmt.Sprintf(format, args...)
}

func modifiedpassthrough(format string, args ...interface{}) {
	fmt.Sprintf(format, args...)
	format = "blah"
	// but not after they modify the format
	fmt.Sprintf(format, args...) // want "variable `format` used for fmt.Sprintf format parameter"
}

func complicated(a, b, format string, args ...interface{}) {
	fmt.Sprintf(format, args...)
}

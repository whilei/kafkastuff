package main

import (
	"fmt"
	"reflect"
)

func compile(t reflect.Type) {

	fmt.Printf("func (r *%s) ToWire() (b []byte) {\n", t.Name())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		switch f.Type.Kind() {
		case reflect.Int16:
			fmt.Printf("  b = append(b, 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.%s))\n", f.Name)

		case reflect.Int32:
			fmt.Printf("  b = append(b, 0, 0, 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.%s))\n", f.Name)

		case reflect.String:
			fmt.Printf("  b = append(b, 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.%s)))\n", f.Name)
			fmt.Printf("  b = append(b, []byte(r.%s)...)\n", f.Name)

		case reflect.Array:
			fmt.Printf("  b = append(b, 0, 0 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.%s)))\n", f.Name)
			fmt.Printf("  for _, el := range r.%s {\n", f.Name)
			fmt.Printf("    b = append(b, []byte(r.%s.ToWire())...)\n", f.Name)
			fmt.Printf("  }\n")
		}
	}

	fmt.Println("  return")
	fmt.Println("}")

}

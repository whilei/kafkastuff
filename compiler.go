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

		case reflect.Int8:
			fmt.Printf("  b = append(b, byte(r.%s))\n", f.Name)

		case reflect.Int16:
			fmt.Printf("  b = append(b, 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint16(b[len(b)-2:], uint16(r.%s))\n", f.Name)

		case reflect.Int32:
			fmt.Printf("  b = append(b, 0, 0, 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint32(b[len(b)-4:], uint32(r.%s))\n", f.Name)

		case reflect.Int64:
			fmt.Printf("  b = append(b, 0, 0, 0, 0, 0, 0, 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint64(b[len(b)-8:], uint64(r.%s))\n", f.Name)

		case reflect.String:
			fmt.Printf("  b = append(b, 0, 0)\n")
			fmt.Printf("  binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(r.%s)))\n", f.Name)
			fmt.Printf("  b = append(b, []byte(r.%s)...)\n", f.Name)

		case reflect.Slice:
			switch f.Type.Elem().Kind() {
			case reflect.Uint8: // byte
				fmt.Printf("  b = append(b, 0, 0, 0, 0)\n")
				fmt.Printf("  binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.%s)))\n", f.Name)
				fmt.Printf("  b = append(b, r.%s...)\n", f.Name)
			case reflect.Int32: // byte
				fmt.Printf("  b = append(b, 0, 0, 0, 0)\n")
				fmt.Printf("  binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.%s)))\n", f.Name)
				fmt.Printf("  for _, el := range r.%s {\n", f.Name)
				fmt.Printf("    b = append(b, 0, 0, 0, 0)\n")
				fmt.Printf("    binary.BigEndian.PutUint32(b[len(b)-4:], uint32(el))\n")
				fmt.Printf("  }\n")
			case reflect.String:
				fmt.Printf("  b = append(b, 0, 0, 0, 0)\n")
				fmt.Printf("  binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.%s)))\n", f.Name)
				fmt.Printf("  for _, el := range r.%s {\n", f.Name)
				fmt.Printf("    b = append(b, 0, 0)\n")
				fmt.Printf("    binary.BigEndian.PutUint16(b[len(b)-2:], uint16(len(el)))\n")
				fmt.Printf("    b = append(b, []byte(el)...)\n")
				fmt.Printf("  }\n")
			default:
				fmt.Printf("  b = append(b, 0, 0, 0, 0)\n")
				fmt.Printf("  binary.BigEndian.PutUint32(b[len(b)-4:], uint32(len(r.%s)))\n", f.Name)
				fmt.Printf("  for _, el := range r.%s {\n", f.Name)
				fmt.Printf("    b = append(b, []byte(el.ToWire())...)\n")
				fmt.Printf("  }\n")
			}

		default:
			fmt.Printf("  b = append(b, []byte(r.%s.ToWire())...)\n", f.Name)

		}
	}

	fmt.Println("  return")
	fmt.Println("}")

}

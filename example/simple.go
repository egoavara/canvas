package main

import (
	"strings"
	"github.com/iamGreedy/canvas/psvg"
	"fmt"
)

func main()  {
	svgpath := strings.NewReader(`M0,0 L16,16 L16,0 Z`)
	parser := psvg.NewParser(svgpath)
	fmt.Println(parser.Next())

	//for e := parser.Next(); e != nil; e = parser.Next() {
	//	switch ee := e.(type) {
	//	case psvg.UnknownError:
	//		panic(ee.Error())
	//	default:
	//		fmt.Println(ee)
	//	}
	//}
	for _, e := range parser.Stream() {
		switch ee := e.(type) {
		case psvg.UnknownError:
			panic(ee.Error())
		default:
			fmt.Println(ee)
		}
	}
}
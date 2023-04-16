package main

import "fmt"

type Pos struct {
	posx, posy int
}

type Objects struct {
	objName  string
	objShape string
	Pos
}

type CustumObjects struct {
	Objects
	newProperty string
}

func main() {
	O := Objects{
		objName:  "obj 1",
		objShape: "circle",
		Pos: Pos{
			posx: 10,
			posy: 10,
		},
	}

	fmt.Printf("object O : %+v\n", O)

	C := CustumObjects{
		newProperty: "multi edge box",
		Objects: Objects{
			objShape: "3d",
			Pos: Pos{
				posx: 4,
			},
		},
	}
	C.Objects.objName = "multi edge"

	fmt.Printf("object  C : %+v\n", C)
}

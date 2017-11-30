package main

import (
	rm "github.com/resurtm/AstroFight/resourceman"
	"fmt"
)

func main() {
	resourceman := rm.New()

	capsule := resourceman.Load("capsule/capsule.obj", rm.MeshResource)
	cube := resourceman.Load("cube/cube.obj", rm.TextureResource)

	fmt.Printf("%#v\n", capsule);
	fmt.Printf("%#v\n", cube);
}

package main

import "swigtests/template_opaque"

func main() {
	v := template_opaque.NewOpaqueVectorType(int64(10))

	template_opaque.FillVector(v)
}

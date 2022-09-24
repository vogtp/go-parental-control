//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			// gen.FeaturePrivacy,
			// gen.FeatureEntQL,
			gen.FeatureUpsert,
		},
		// Templates: []*gen.Template{
		// 	gen.MustParse(gen.NewTemplate("static").
		// 		Funcs(template.FuncMap{"title": strings.ToTitle}).
		// 		ParseFiles("template/static.tmpl")),
		// },
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

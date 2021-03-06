//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		// enables filter generation
		entgql.WithWhereFilters(true),
		// configures the path to the gqlgen config file
		entgql.WithConfigPath("./gqlgen.yml"),
		// generate gql schema from the entc schema
		entgql.WithSchemaGenerator(),
		// configures a path to a new or existing GraphQL schema to write the generated filters to
		entgql.WithSchemaPath("./ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
		// if we wanted to use templates, we could add the following:
		// entc.TemplateDir("./template"),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

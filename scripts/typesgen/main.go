package main

import (
	"fmt"
	"log"

	"github.com/coder/guts"
	"github.com/coder/guts/config"
)

func main() {
	gen, err := guts.NewGolangParser()
	if err != nil {
		log.Fatalf("go parser: %v", err)
	}

	generatePackages := map[string]string{
		"github.com/JollyGrin/postgres-attendance/internal/model": "",
	}

	for pkg, prefix := range generatePackages {
		err = gen.IncludeGenerateWithPrefix(pkg, prefix)
		if err != nil {
			log.Fatalf("include generate package %q: %v", pkg, err)
		}
	}

	// Standard type mappings
	gen.IncludeCustomDeclaration(config.StandardMappings())

	ts, err := gen.ToTypescript()
	if err != nil {
		log.Fatalf("to typescript: %v", err)
	}

	ts.ApplyMutations(
		config.SimplifyOmitEmpty,
		config.ExportTypes,
		config.ReadOnly,
	)

	output, err := ts.Serialize()
	if err != nil {
		log.Fatalf("serialize: %v", err)
	}

	fmt.Println("// Run `make gen` from the root of the repository to regenerate this file.")
	fmt.Println(output)
}

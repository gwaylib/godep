// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importers

import (
	"log"

	"github.com/gwaylib/godep"
	"github.com/gwaylib/godep/gps"
	"github.com/gwaylib/godep/internal/importers/glide"
	"github.com/gwaylib/godep/internal/importers/glock"
	"github.com/gwaylib/godep/internal/importers/godep"
	"github.com/gwaylib/godep/internal/importers/govend"
	"github.com/gwaylib/godep/internal/importers/govendor"
	"github.com/gwaylib/godep/internal/importers/gvt"
	"github.com/gwaylib/godep/internal/importers/vndr"
)

// Importer handles importing configuration from other dependency managers into
// the dep configuration format.
type Importer interface {
	// Name of the importer.
	Name() string

	// Import the config found in the directory.
	Import(path string, pr gps.ProjectRoot) (*dep.Manifest, *dep.Lock, error)

	// HasDepMetadata checks if a directory contains config that the importer can handle.
	HasDepMetadata(dir string) bool
}

// BuildAll returns a slice of all the importers.
func BuildAll(logger *log.Logger, verbose bool, sm gps.SourceManager) []Importer {
	return []Importer{
		glide.NewImporter(logger, verbose, sm),
		godep.NewImporter(logger, verbose, sm),
		vndr.NewImporter(logger, verbose, sm),
		govend.NewImporter(logger, verbose, sm),
		gvt.NewImporter(logger, verbose, sm),
		govendor.NewImporter(logger, verbose, sm),
		glock.NewImporter(logger, verbose, sm),
	}
}

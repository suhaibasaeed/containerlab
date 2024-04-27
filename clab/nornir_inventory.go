// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package clab

import (
	_ "embed"
	"io"
	"os"
	"sort"
	"text/template"

	"github.com/srl-labs/containerlab/types"
)

//go:embed inventory_nornir.go.tpl
var nornirInvT string
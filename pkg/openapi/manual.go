// Copyright 2018 The CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/encoding/openapi"
	cueyaml "cuelang.org/go/internal/encoding/yaml"
)

// Gen generates the set OpenAPI schema for all top-level types of the given instance.
func Gen(s cue.Value) (ast.Expr, error) {
	b, err := openapi.Gen(s, &openapi.Config{ExpandReferences: true})
	if err != nil {
		return nil, err
	}
	e, err := cueyaml.Unmarshal("", b)
	if err != nil {
		return nil, err
	}
	return e, err
}

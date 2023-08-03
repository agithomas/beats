// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kubernetes.
func AssetKubernetes() string {
	return "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"
}

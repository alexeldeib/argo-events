/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"github.com/argoproj/argo-events/gateways"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	configKey   = "testConfig"
	configValue = `
directory: "/bin/"
type: CREATE
path: x.txt
`
)

func TestFileWatcherConfigExecutor_Validate(t *testing.T) {
	ce := &FileWatcherConfigExecutor{}
	ctx := &gateways.ConfigContext{
		Data: &gateways.ConfigData{},
	}
	ctx.Data.Config = configValue
	err := ce.Validate(ctx)
	assert.Nil(t, err)
	configValue = `
directory: "/bin/"
type: CREATE
`
	ctx.Data.Config = configValue
	err = ce.Validate(ctx)
	assert.NotNil(t, err)
	configValue = ``
	ctx.Data.Config = configValue
	err = ce.Validate(ctx)
	assert.NotNil(t, err)
}

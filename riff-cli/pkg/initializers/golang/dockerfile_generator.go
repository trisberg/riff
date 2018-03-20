/*
 * Copyright 2017 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package golang

import (
	"github.com/projectriff/riff/riff-cli/pkg/options"
	"path/filepath"
	"github.com/projectriff/riff/riff-cli/pkg/initializers/core"
)

const goFunctionDockerfileTemplate = `
FROM projectriff/go-function-invoker:{{.InvokerVersion}}
ADD {{.Artifact}} /
ENV FUNCTION_URI file:///{{.ArtifactBase}}?handler={{.Handler}}
`

func generateGoFunctionDockerFile(opts options.InitOptions) (string, error) {
	dockerFileTokens := core.DockerFileTokens{
		Artifact:       opts.Artifact,
		ArtifactBase:   filepath.Base(opts.Artifact),
		InvokerVersion: opts.InvokerVersion,
		Handler:        opts.Handler,
	}
	return core.GenerateFunctionDockerFileContents(goFunctionDockerfileTemplate, "docker-go", dockerFileTokens)
}

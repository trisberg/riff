/*
 * Copyright 2018 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"github.com/projectriff/riff/pkg/core/kustomize"
	"github.com/projectriff/riff/pkg/kubectl"
	"k8s.io/client-go/kubernetes"
	"time"
)

type KubectlClient interface {
	SystemInstall(manifests map[string]*Manifest, options SystemInstallOptions) (bool, error)
	SystemUninstall(options SystemUninstallOptions) (bool, error)
	NamespaceInit(manifests map[string]*Manifest, options NamespaceInitOptions) error
	NamespaceCleanup(options NamespaceCleanupOptions) error
}

type kubectlClient struct {
	kubeClient kubernetes.Interface
	kubeCtl    kubectl.KubeCtl
	kustomizer kustomize.Kustomizer
}

func NewKubectlClient(kubeClient kubernetes.Interface) KubectlClient {
	httpTimeout := 30 * time.Second
	return &kubectlClient{
		kubeClient: kubeClient,
		kubeCtl:    kubectl.RealKubeCtl(),
		kustomizer: kustomize.MakeKustomizer(httpTimeout),
	}
}

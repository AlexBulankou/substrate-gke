// Copyright 2026 Google LLC
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

// Package k8srelease holds the Kubernetes release floor Agent Substrate is
// supported on. It is the one place that number is written.
//
// It is its own module, outside any internal/ tree, because the two things
// that enforce the floor are in two different modules: the installer
// (github.com/ai-on-gke/substrate-gke/installer) and the setup-gcp CLI
// (github.com/ai-on-gke/substrate-gke/tools/setup-gcp). The installer's copy
// is in installer/internal/gcp, which setup-gcp cannot import even in
// principle, so the two were hand-synced and free to drift apart.
package k8srelease

// The oldest Kubernetes release Agent Substrate is supported on, for the
// control plane and for every node pool. A cluster with either below it is
// rejected; the remedy is an upgrade.
//
// Node pools are checked because GKE lets them trail the control plane, and a
// kubelet below this serves pod certificate projection only behind the beta
// gate: Substrate's pods on such a node fail to mount with "unimplemented".
//
// https://docs.cloud.google.com/kubernetes-engine/ai-ml/install-overview-substrate#cluster-requirements
const (
	MinSupportedMajor = 1
	MinSupportedMinor = 37
)

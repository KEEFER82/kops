/*
Copyright 2019 The Kubernetes Authors.

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

package kubeapiserver

import (
	"fmt"
	"path/filepath"
	"testing"

	"k8s.io/kops/pkg/assets"
	"k8s.io/kops/pkg/featureflag"
	"k8s.io/kops/pkg/model"
	"k8s.io/kops/pkg/model/iam"
	"k8s.io/kops/pkg/testutils"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/util/pkg/vfs"
)

func Test_RunKubeApiserverBuilder(t *testing.T) {
	featureflag.ParseFlags("-ImageDigest")
	tests := []string{
		"tests/minimal",
	}
	for _, basedir := range tests {
		basedir := basedir

		t.Run(fmt.Sprintf("basedir=%s", basedir), func(t *testing.T) {
			context := &fi.CloudupModelBuilderContext{
				Tasks: make(map[string]fi.CloudupTask),
			}
			kopsModelContext, err := LoadKopsModelContext(basedir)
			if err != nil {
				t.Fatalf("error loading model %q: %v", basedir, err)
				return
			}

			builder := KubeApiserverBuilder{
				KopsModelContext: kopsModelContext,
				AssetBuilder:     assets.NewAssetBuilder(vfs.Context, kopsModelContext.Cluster.Spec.Assets, kopsModelContext.Cluster.Spec.KubernetesVersion, false),
			}

			if err := builder.Build(context); err != nil {
				t.Fatalf("error from Build: %v", err)
				return
			}

			testutils.ValidateTasks(t, filepath.Join(basedir, "tasks.yaml"), context)
		})
	}
}

func LoadKopsModelContext(basedir string) (*model.KopsModelContext, error) {
	spec, err := testutils.LoadModel(basedir)
	if err != nil {
		return nil, err
	}

	if spec.Cluster == nil {
		return nil, fmt.Errorf("no cluster found in %s", basedir)
	}

	kopsContext := &model.KopsModelContext{
		IAMModelContext: iam.IAMModelContext{Cluster: spec.Cluster},
		InstanceGroups:  spec.InstanceGroups,
	}

	return kopsContext, nil
}

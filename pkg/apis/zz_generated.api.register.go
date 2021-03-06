/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by apiregister-gen. DO NOT EDIT.

package apis

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
	"steamingmind.com/apiserver-demo/pkg/apis/demo"
	_ "steamingmind.com/apiserver-demo/pkg/apis/demo/install" // Install the demo group
	demov1alpha1 "steamingmind.com/apiserver-demo/pkg/apis/demo/v1alpha1"
)

var (
	localSchemeBuilder = runtime.SchemeBuilder{
		demov1alpha1.AddToScheme,
	}
	AddToScheme = localSchemeBuilder.AddToScheme
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetDemoAPIBuilder(),
	}
}

var demoApiGroup = builders.NewApiGroupBuilder(
	"demo.streamingmind.com",
	"steamingmind.com/apiserver-demo/pkg/apis/demo").
	WithUnVersionedApi(demo.ApiVersion).
	WithVersionedApis(
		demov1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetDemoAPIBuilder() *builders.APIGroupBuilder {
	return demoApiGroup
}

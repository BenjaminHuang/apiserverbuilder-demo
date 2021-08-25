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

package install

import (
	"k8s.io/apiserver/pkg/admission"
	genericserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/cmd/server"
	aggregatedclientset "steamingmind.com/apiserver-demo/pkg/client/clientset_generated/clientset"
	aggregatedinformerfactory "steamingmind.com/apiserver-demo/pkg/client/informers_generated/externalversions"
	initializer "steamingmind.com/apiserver-demo/plugin/admission"
	. "steamingmind.com/apiserver-demo/plugin/admission/foo"
)

func init() {
	server.AggregatedAdmissionInitializerGetter = GetAggregatedResourceAdmissionControllerInitializer
	server.AggregatedAdmissionPlugins["Foo"] = NewFooPlugin()

}

func GetAggregatedResourceAdmissionControllerInitializer(config *rest.Config) (admission.PluginInitializer, genericserver.PostStartHookFunc) {
	// init aggregated resource clients
	aggregatedResourceClient := aggregatedclientset.NewForConfigOrDie(config)
	aggregatedInformerFactory := aggregatedinformerfactory.NewSharedInformerFactory(aggregatedResourceClient, 0)
	aggregatedResourceInitializer := initializer.New(aggregatedResourceClient, aggregatedInformerFactory)

	return aggregatedResourceInitializer, func(context genericserver.PostStartHookContext) error {
		aggregatedInformerFactory.Start(context.StopCh)
		return nil
	}
}

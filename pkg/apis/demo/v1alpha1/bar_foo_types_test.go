
/*
Copyright YEAR The Kubernetes Authors.

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


package v1alpha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "steamingmind.com/apiserver-demo/pkg/apis/demo/v1alpha1"
	. "steamingmind.com/apiserver-demo/pkg/client/clientset_generated/clientset/typed/demo/v1alpha1"
)

var _ = Describe("Foo", func() {
	var instance Foo
	var expected Foo
	var client FooInterface

	BeforeEach(func() {
		instance = Foo{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(context.TODO(), instance.Name, metav1.DeleteOptions{})
	})

	Describe("when sending a bar request", func() {
		It("should return success", func() {
			client = cs.DemoV1alpha1().Foos("foo-test-bar")
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			bar := &FooBar{}
			bar.Name = instance.Name
			restClient := cs.DemoV1alpha1().RESTClient()
			err = restClient.Post().Namespace("foo-test-bar").
				Name(instance.Name).
				Resource("foos").
				SubResource("bar").
				Body(bar).
				Do(context.TODO()).
				Error()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})

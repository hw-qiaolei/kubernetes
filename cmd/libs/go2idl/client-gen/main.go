/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// client-gen makes the individual typed clients using go2idl.
package main

import (
	"k8s.io/kubernetes/cmd/libs/go2idl/args"
	"k8s.io/kubernetes/cmd/libs/go2idl/client-gen/generators"

	"github.com/golang/glog"
)

func main() {
	arguments := args.Default()

	// Override defaults. These are Kubernetes specific input and output
	// locations.
	arguments.InputDirs = []string{
		"k8s.io/kubernetes/pkg/api",
		"k8s.io/kubernetes/pkg/apis/extensions",
		"k8s.io/kubernetes/pkg/fields",
		"k8s.io/kubernetes/pkg/labels",
		"k8s.io/kubernetes/pkg/watch",
	}
	// We may change the output path later.
	arguments.OutputPackagePath = "k8s.io/kubernetes/pkg/client/clientset/unversioned"

	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
}

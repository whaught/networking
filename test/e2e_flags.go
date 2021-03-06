/*
Copyright 2020 The Knative Authors

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

// This file contains logic to encapsulate flags which are needed to specify
// what cluster, etc. to use for e2e tests.

package test

import (
	"flag"
)

// ServingFlags holds the flags or defaults for knative/serving settings in the user's environment.
var ServingFlags = initializeServingFlags()

// ServingEnvironmentFlags holds the e2e flags needed only by the serving repo.
type ServingEnvironmentFlags struct {
	ResolvableDomain bool   // Resolve Route controller's `domainSuffix`
	Https            bool   // Indicates where the test service will be created with https
	IngressClass     string // Indicates the class of Ingress provider to test.
	CertificateClass string // Indicates the class of Certificate provider to test.
	SystemNamespace  string // Indicates the system namespace, in which Knative Serving is installed.
}

func initializeServingFlags() *ServingEnvironmentFlags {
	var f ServingEnvironmentFlags

	// Only define and set flags here. Flag values cannot be read at package init time.
	flag.BoolVar(&f.ResolvableDomain, "resolvabledomain", false,
		"Set this flag to true if you have configured the `domainSuffix` on your Route controller to a domain that will resolve to your test cluster.")
	flag.BoolVar(&f.Https, "https", false,
		"Set this flag to true to run all tests with https.")

	flag.StringVar(&f.IngressClass, "ingressClass", "istio.ingress.networking.knative.dev",
		"Set this flag to the ingress class to test against.")
	flag.StringVar(&f.CertificateClass, "certificateClass", "cert-manager.certificate.networking.knative.dev",
		"Set this flag to the certificate class to test against.")

	return &f
}

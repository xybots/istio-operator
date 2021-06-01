/*
Copyright 2021 Banzai Cloud.

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

package gvr

import "k8s.io/apimachinery/pkg/runtime/schema"

var Service = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "services",
}

var Pod = schema.GroupVersionResource{
	Group:    "",
	Version:  "v1",
	Resource: "pods",
}

var Deployment = schema.GroupVersionResource{
	Group:    "apps",
	Version:  "v1",
	Resource: "deployments",
}

var HorizontalPodAutoscaler = schema.GroupVersionResource{
	Group:    "autoscaling",
	Version:  "v1",
	Resource: "horizontalpodautoscalers",
}

var ClusterRole = schema.GroupVersionResource{
	Group:    "rbac.authorization.k8s.io",
	Version:  "v1",
	Resource: "clusterroles",
}

var ClusterRoleBinding = schema.GroupVersionResource{
	Group:    "rbac.authorization.k8s.io",
	Version:  "v1",
	Resource: "clusterrolebindings",
}

var ValidatingWebhookConfiguration = schema.GroupVersionResource{
	Group:    "admissionregistration.k8s.io",
	Version:  "v1",
	Resource: "validatingwebhookconfigurations",
}

var MutatingWebhookConfiguration = schema.GroupVersionResource{
	Group:    "admissionregistration.k8s.io",
	Version:  "v1",
	Resource: "mutatingwebhookconfigurations",
}

var DestinationRule = schema.GroupVersionResource{
	Group:    "networking.istio.io",
	Version:  "v1alpha3",
	Resource: "destinationrules",
}

var VirtualService = schema.GroupVersionResource{
	Group:    "networking.istio.io",
	Version:  "v1alpha3",
	Resource: "virtualservices",
}

var PeerAuthentication = schema.GroupVersionResource{
	Group:    "security.istio.io",
	Version:  "v1beta1",
	Resource: "peerauthentications",
}

var Gateway = schema.GroupVersionResource{
	Group:    "networking.istio.io",
	Version:  "v1alpha3",
	Resource: "gateways",
}

var EnvoyFilter = schema.GroupVersionResource{
	Group:    "networking.istio.io",
	Version:  "v1alpha3",
	Resource: "envoyfilters",
}

var AttributeManifest = schema.GroupVersionResource{
	Group:    "config.istio.io",
	Version:  "v1alpha2",
	Resource: "attributemanifests",
}

var IstioConfigHandler = schema.GroupVersionResource{
	Group:    "config.istio.io",
	Version:  "v1alpha2",
	Resource: "handlers",
}

var IstioConfigInstance = schema.GroupVersionResource{
	Group:    "config.istio.io",
	Version:  "v1alpha2",
	Resource: "instances",
}

var IstioConfigRule = schema.GroupVersionResource{
	Group:    "config.istio.io",
	Version:  "v1alpha2",
	Resource: "rules",
}

var Istio = schema.GroupVersionResource{
	Group:    "istio.banzaicloud.io",
	Version:  "v1beta1",
	Resource: "istios",
}

var MeshGateway = schema.GroupVersionResource{
	Group:    "istio.banzaicloud.io",
	Version:  "v1beta1",
	Resource: "meshgateways",
}

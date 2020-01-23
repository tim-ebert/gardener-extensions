/*
Copyright 2018 The Kubernetes Authors.

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

package main

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// svcValidator intercepts updates to the service status of the kube-apiserver svc in the shoot's control plane
type svcValidator struct {
	client  client.Client
	decoder *admission.Decoder
	logger  logr.Logger
}

// svcValidator admits a pod iff a specific annotation exists.
func (v *svcValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	v.logger.Info("Received request")
	if req.Resource.Resource != "services" || req.Resource.Group != "" || req.SubResource != "status" {
		return admission.Allowed("")
	}

	service := &corev1.Service{}

	err := v.decoder.Decode(req, service)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	v.logger.Info("received update for service status", "namespace", service.Namespace, "name", service.Name)

	if ingress := service.Status.LoadBalancer.Ingress; len(ingress) > 0 {
		for _, i := range ingress {
			if i.Hostname == "localhost" {
				return admission.Errored(400, fmt.Errorf("not allowed to set localhost as LoadBalancer ingress hostname"))
			}
		}
	}

	return admission.Allowed("")
}

func (v *svcValidator) InjectClient(c client.Client) error {
	v.client = c
	return nil
}

func (v *svcValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}

#!/bin/bash

echo "Generating openapi definitions"
/mnt/c/Users/i313632/go/bin/openapi-gen \
  --v 1 \
  --logtostderr \
  --input-dirs=github.com/gardener/gardener/pkg/apis/core/v1alpha1 \
  --input-dirs=github.com/gardener/gardener/pkg/apis/core/v1beta1 \
  --input-dirs=github.com/gardener/gardener/pkg/apis/settings/v1alpha1 \
  --input-dirs=github.com/gardener/gardener/pkg/apis/garden/v1beta1 \
  --input-dirs=k8s.io/api/core/v1 \
  --input-dirs=k8s.io/api/rbac/v1 \
  --input-dirs=k8s.io/apimachinery/pkg/apis/meta/v1 \
  --input-dirs=k8s.io/apimachinery/pkg/api/resource \
  --input-dirs=k8s.io/apimachinery/pkg/types \
  --input-dirs=k8s.io/apimachinery/pkg/version \
  --input-dirs=k8s.io/apimachinery/pkg/runtime \
  --input-dirs=k8s.io/apimachinery/pkg/util/intstr \
  --report-filename=${PROJECT_ROOT}/pkg/openapi/api_violations.report \
  --output-package=github.com/gardener/gardener/pkg/openapi 
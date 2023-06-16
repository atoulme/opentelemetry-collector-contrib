// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package container // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/container"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	corev1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/docker"
	metadataPkg "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	imetadata "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/container/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/utils"
)

const (
	// Keys for container metadata.
	containerKeyStatus       = "container.status"
	containerKeyStatusReason = "container.status.reason"

	// Values for container metadata
	containerStatusRunning    = "running"
	containerStatusWaiting    = "waiting"
	containerStatusTerminated = "terminated"
)

// GetSpecMetrics metricizes values from the container spec.
// This includes values like resource requests and limits.
func GetSpecMetrics(set receiver.CreateSettings, c corev1.Container, pod *corev1.Pod) pmetric.Metrics {
	mb := imetadata.NewMetricsBuilder(imetadata.DefaultMetricsBuilderConfig(), set)
	ts := pcommon.NewTimestampFromTime(time.Now())
	mb.RecordK8sContainerCPURequestDataPoint(ts, float64(c.Resources.Requests.Cpu().MilliValue())/1000.0)
	mb.RecordK8sContainerCPULimitDataPoint(ts, float64(c.Resources.Limits.Cpu().MilliValue())/1000.0)
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.Name == c.Name {
			mb.RecordK8sContainerRestartsDataPoint(ts, int64(cs.RestartCount))
			mb.RecordK8sContainerReadyDataPoint(ts, boolToInt64(cs.Ready))
			break
		}
	}
	image, err := docker.ParseImageName(c.Image)
	if err != nil {
		docker.LogParseError(err, c.Image, set.Logger)
	}
	var containerID string
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.Name == c.Name {
			containerID = cs.ContainerID
		}
	}
	return mb.Emit(imetadata.WithK8sPodUID(string(pod.UID)),
		imetadata.WithK8sPodName(pod.Name),
		imetadata.WithK8sNodeName(pod.Spec.NodeName),
		imetadata.WithK8sNamespaceName(pod.Namespace),
		imetadata.WithOpencensusResourcetype("container"),
		imetadata.WithContainerID(utils.StripContainerID(containerID)),
		imetadata.WithK8sContainerName(c.Name),
		imetadata.WithContainerImageName(image.Repository),
		imetadata.WithContainerImageTag(image.Tag),
	)
}

func GetMetadata(cs corev1.ContainerStatus) *metadata.KubernetesMetadata {
	mdata := map[string]string{}

	if cs.State.Running != nil {
		mdata[containerKeyStatus] = containerStatusRunning
	}

	if cs.State.Terminated != nil {
		mdata[containerKeyStatus] = containerStatusTerminated
		mdata[containerKeyStatusReason] = cs.State.Terminated.Reason
	}

	if cs.State.Waiting != nil {
		mdata[containerKeyStatus] = containerStatusWaiting
		mdata[containerKeyStatusReason] = cs.State.Waiting.Reason
	}

	return &metadata.KubernetesMetadata{
		ResourceIDKey: conventions.AttributeContainerID,
		ResourceID:    metadataPkg.ResourceID(utils.StripContainerID(cs.ContainerID)),
		Metadata:      mdata,
	}
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package splunkhecexporter

import (
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/translator/conventions"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk"
)

// HecEvent is a data structure holding a span event to export explicitly to Splunk HEC.
type HecEvent struct {
	Attributes map[string]interface{}  `json:"attributes,omitempty"`
	Name       string                  `json:"name"`
	Timestamp  pdata.TimestampUnixNano `json:"timestamp"`
}

// HecLink is a data structure holding a span link to export explicitly to Splunk HEC.
type HecLink struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	TraceID    []byte                 `json:"trace_id"`
	SpanID     []byte                 `json:"span_id"`
	TraceState pdata.TraceState       `json:"trace_state"`
}

// HecSpanStatus is a data structure holding the status of a span to export explicitly to Splunk HEC.
type HecSpanStatus struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// HecSpan is a data structure used to export explicitly a span to Splunk HEC.
type HecSpan struct {
	TraceID    []byte                  `json:"trace_id"`
	SpanID     []byte                  `json:"span_id"`
	ParentSpan []byte                  `json:"parent_span_id"`
	Name       string                  `json:"name"`
	Attributes map[string]interface{}  `json:"attributes,omitempty"`
	EndTime    pdata.TimestampUnixNano `json:"end_time"`
	Kind       string                  `json:"kind"`
	Status     HecSpanStatus           `json:"status,omitempty"`
	StartTime  pdata.TimestampUnixNano `json:"start_time"`
	Events     []HecEvent              `json:"events,omitempty"`
	Links      []HecLink               `json:"links,omitempty"`
}

func traceDataToSplunk(logger *zap.Logger, data pdata.Traces, config *Config) ([]*splunk.Event, int) {
	numDroppedSpans := 0
	splunkEvents := make([]*splunk.Event, 0, data.SpanCount())
	for i := 0; i < data.ResourceSpans().Len(); i++ {
		rs := data.ResourceSpans().At(i)
		host := unknownHostName
		source := config.Source
		sourceType := config.SourceType
		if !rs.Resource().IsNil() {
			if conventionHost, isSet := rs.Resource().Attributes().Get(conventions.AttributeHostHostname); isSet {
				host = conventionHost.StringVal()
			}
			if sourceSet, isSet := rs.Resource().Attributes().Get(conventions.AttributeServiceName); isSet {
				source = sourceSet.StringVal()
			}
			if sourcetypeSet, isSet := rs.Resource().Attributes().Get(splunk.SourcetypeLabel); isSet {
				sourceType = sourcetypeSet.StringVal()
			}
		}
		for sils := 0; sils < rs.InstrumentationLibrarySpans().Len(); sils++ {
			ils := rs.InstrumentationLibrarySpans().At(sils)
			for si := 0; si < ils.Spans().Len(); si++ {
				span := ils.Spans().At(si)
				if span.IsNil() {
					continue
				}

				se := &splunk.Event{
					Time:       timestampToEpochMilliseconds(span.StartTime()),
					Host:       host,
					Source:     source,
					SourceType: sourceType,
					Index:      config.Index,
					Event:      toHecSpan(logger, span),
				}
				splunkEvents = append(splunkEvents, se)
			}
		}
	}

	return splunkEvents, numDroppedSpans
}

func toHecSpan(logger *zap.Logger, span pdata.Span) HecSpan {
	attributes := map[string]interface{}{}
	span.Attributes().ForEach(func(k string, v pdata.AttributeValue) {
		attributes[k] = convertAttributeValue(v, logger)
	})

	links := make([]HecLink, span.Links().Len())
	for i := 0; i < span.Links().Len(); i++ {
		link := span.Links().At(i)
		linkAttributes := map[string]interface{}{}
		link.Attributes().ForEach(func(k string, v pdata.AttributeValue) {
			linkAttributes[k] = convertAttributeValue(v, logger)
		})
		links[i] = HecLink{
			Attributes: linkAttributes,
			TraceID:    convert16BytesToBytes(link.TraceID().Bytes()),
			SpanID:     convert8BytesToBytes(link.SpanID().Bytes()),
			TraceState: link.TraceState(),
		}
	}
	events := make([]HecEvent, span.Events().Len())
	for i := 0; i < span.Events().Len(); i++ {
		event := span.Events().At(i)
		eventAttributes := map[string]interface{}{}
		event.Attributes().ForEach(func(k string, v pdata.AttributeValue) {
			eventAttributes[k] = convertAttributeValue(v, logger)
		})
		events[i] = HecEvent{
			Attributes: eventAttributes,
			Name:       event.Name(),
			Timestamp:  event.Timestamp(),
		}
	}
	var status HecSpanStatus
	if !span.Status().IsNil() {
		status = HecSpanStatus{
			Message: span.Status().Message(),
			Code:    span.Status().Code().String(),
		}
	}
	return HecSpan{
		TraceID:    convert16BytesToBytes(span.TraceID().Bytes()),
		SpanID:     convert8BytesToBytes(span.SpanID().Bytes()),
		ParentSpan: convert8BytesToBytes(span.ParentSpanID().Bytes()),
		Name:       span.Name(),
		Attributes: attributes,
		StartTime:  span.StartTime(),
		EndTime:    span.EndTime(),
		Kind:       span.Kind().String(),
		Status:     status,
		Links:      links,
		Events:     events,
	}
}

func convert16BytesToBytes(bytes [16]byte) []byte {
	return bytes[:]
}

func convert8BytesToBytes(bytes [8]byte) []byte {
	return bytes[:]
}

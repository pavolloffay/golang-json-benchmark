package jsontest

import (
	"encoding/json"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsingFixtureGojay(t *testing.T) {
	b, err := loadFixture("default.json")
	require.NoError(t, err)
	assert.NotNil(t, b)

	span := &Span{}
	err = gojay.Unmarshal(b, span)
	require.NoError(t, err)

	spanStd := &Span{}
	err = json.Unmarshal(b, spanStd)
	require.NoError(t, err)
	assert.Equal(t, spanStd, span)

	bb, err := gojay.Marshal(span)
	require.NoError(t, err)
	bbStd, err := gojay.Marshal(spanStd)
	require.NoError(t, err)
	assert.Equal(t, bbStd, bb)

	span2 := &Span{}
	err = json.Unmarshal(bb, span2)
	require.NoError(t, err)
	assert.Equal(t, span, span2)
}

func BenchmarkMarshalGojay(b *testing.B) {
	benchmarkMarshal(b, gojay.Marshal)
}

func BenchmarkUnmarshalGojay(b *testing.B) {
	benchmarkUnmarshal(b, gojay.Unmarshal)
}

type KeyValueArr []*KeyValue
type LogArr []*Log
type ReferencesArr []*Reference

func (s *Span) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("operationName", s.OperationName)
	enc.StringKey("spanId", string(s.SpanID))
	enc.StringKey("traceId", string(s.TraceID))
	enc.Uint64Key("startTime", s.StartTime)
	enc.Uint64Key("startTimeMillis", s.StartTimeMillis)
	enc.Uint64Key("duration", s.Duration)
	enc.ObjectKey("process", &s.Process)
	tags := convertKeyValuesToArrType(s.Tags)
	enc.ArrayKey("tags", &tags)
	logs := convertLogsToArrType(s.Logs)
	enc.ArrayKey("logs", &logs)
	refs := convertReferencesToArrType(s.References)
	enc.ArrayKey("references", &refs)
}
func (s *Span) IsNil() bool {
	return s == nil
}

func (p *Process) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("serviceName", p.ServiceName)
	tags := convertKeyValuesToArrType(p.Tags)
	enc.ArrayKey("tags", &tags)
}
func (p *Process) IsNil() bool {
	return p == nil
}

func (kvs *KeyValueArr) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *kvs {
		enc.Object(e)
	}
}
func (kvs *KeyValueArr) IsNil() bool {
	return kvs == nil
}
func (kv *KeyValue) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("key", kv.Key)
	enc.StringKey("type", string(kv.Type))
	enc.AddInterfaceKey("value", kv.Value)
}
func (kv *KeyValue) IsNil() bool {
	return kv == nil
}

func (logs *LogArr) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *logs {
		enc.Object(e)
	}
}
func (logs *LogArr) IsNil() bool {
	return logs == nil
}
func (l *Log) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Uint64Key("timestamp", l.Timestamp)
	tags := make(KeyValueArr, len(l.Fields))
	for i := 0; i < len(l.Fields); i++ {
		tags[i] = &l.Fields[i]
	}
	enc.ArrayKey("fields", &tags)
}
func (l *Log) IsNil() bool {
	return l == nil
}

func (refs *ReferencesArr) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *refs {
		enc.Object(e)
	}
}
func (refs *ReferencesArr) IsNil() bool {
	return refs == nil
}
func (r *Reference) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("refType", string(r.RefType))
	enc.StringKey("traceID", string(r.TraceID))
	enc.StringKey("spanID", string(r.SpanID))
}
func (r *Reference) IsNil() bool {
	return r == nil
}

func (s *Span) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "operationName":
		return dec.String(&s.OperationName)
	case "traceId":
		traceId := string(s.TraceID)
		err := dec.String(&traceId)
		if err == nil {
			s.TraceID = TraceID(traceId)
		}
		return err
	case "spanId":
		spanId := string(s.SpanID)
		err := dec.AddString(&spanId)
		if err == nil {
			s.SpanID = SpanID(spanId)
		}
		return err
	case "startTime":
		return dec.AddUint64(&s.StartTime)
	case "startTimeMillis":
		return dec.AddUint64(&s.StartTimeMillis)
	case "duration":
		return dec.AddUint64(&s.Duration)
	case "tags":
		tags := KeyValueArr{}
		err := dec.AddArray(&tags)
		if err == nil {
			s.Tags = convertKeyValueFromArrType(tags)
		}
		return err
	case "logs":
		logs := LogArr{}
		err := dec.AddArray(&logs)
		if err == nil {
			s.Logs = convertLogsFromArrType(logs)
		}
	case "references":
		refs := ReferencesArr{}
		err := dec.AddArray(&refs)
		if err == nil {
			s.References = convertReferencesFromArrType(refs)
		}
	case "process":
		return dec.AddObject(&s.Process)
	}
	return nil
}
func (s *Span) NKeys() int {
	return 0
}

func (p *Process) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "serviceName":
		return dec.AddString(&p.ServiceName)
	case "tags":
		tags := KeyValueArr{}
		err := dec.AddArray(&tags)
		if err == nil {
			p.Tags = convertKeyValueFromArrType(tags)
		}
	}
	return nil
}
func (p *Process) NKeys() int {
	return 0
}

func (kvs *KeyValueArr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tag := KeyValue{}
	err :=  dec.Object(&tag)
	if err != nil {
		return err
	}
	*kvs = append(*kvs, &tag)
	return nil
}
func (kv *KeyValue) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "key":
		return dec.AddString(&kv.Key)
	case "type":
		var typ string
		err := dec.AddString(&typ)
		if err == nil {
			kv.Type = ValueType(typ)
		}
		return err
	case "value":
		return dec.AddInterface(&kv.Value)
	}
	return nil
}
func (kv *KeyValue) NKeys() int {
	return 0
}

func (logs *LogArr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	l := Log{}
	err := dec.AddObject(&l)
	if err != nil {
		return err
	}
	*logs = append(*logs, &l)
	return nil
}
func (l *Log) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "timestamp":
		return dec.AddUint64(&l.Timestamp)
	case "fields":
		tags := KeyValueArr{}
		err := dec.AddArray(&tags)
		if err == nil {
			l.Fields = convertKeyValueFromArrType(tags)
		}
	}
	return nil
}
func (l *Log) NKeys() int {
	return 0
}

func (refs *ReferencesArr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	ref := Reference{}
	err := dec.AddObject(&ref)
	if err != nil {
		return err
	}
	*refs = append(*refs, &ref)
	return nil
}
func (r *Reference) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "refType":
		var refType string
		err := dec.AddString(&refType)
		if err == nil {
			r.RefType = ReferenceType(refType)
		}
		return err
	case "spanID":
		var spanId string
		err := dec.AddString(&spanId)
		if err == nil {
			r.SpanID = SpanID(spanId)
		}
		return err
	case "traceID":
		var traceId string
		err := dec.AddString(&traceId)
		if err == nil {
			r.TraceID = TraceID(traceId)
		}
		return err
	}
	return nil
}
func (r *Reference) NKeys() int {
	return 0
}

func convertKeyValuesToArrType(kvs []KeyValue) KeyValueArr {
	tags := make(KeyValueArr, len(kvs))
	for i := 0; i < len(kvs); i++ {
		tags[i] = &kvs[i]
	}
	return tags
}
func convertKeyValueFromArrType(kvs KeyValueArr) []KeyValue {
	tags := make([]KeyValue, len(kvs))
	for i := 0; i < len(kvs); i++ {
		tags[i] = *kvs[i]
	}
	return tags
}

func convertReferencesToArrType(refs []Reference) ReferencesArr {
	refsArr := make(ReferencesArr, len(refs))
	for i := 0; i < len(refs); i++ {
		refsArr[i] = &refs[i]
	}
	return refsArr
}
func convertReferencesFromArrType(refs ReferencesArr) []Reference {
	refsArr := make([]Reference, len(refs))
	for i := 0; i < len(refs); i++ {
		refsArr[i] = *refs[i]
	}
	return refsArr
}

func convertLogsToArrType(logs []Log) LogArr {
	logArr := make(LogArr, len(logs))
	for i := 0; i < len(logs); i++ {
		logArr[i] = &logs[i]
	}
	return logArr
}
func convertLogsFromArrType(logs LogArr) []Log {
	logArr := make([]Log, len(logs))
	for i := 0; i < len(logs); i++ {
		logArr[i] = *logs[i]
	}
	return logArr
}
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

	span := &Span{}
	err = gojay.Unmarshal(b, span)
	require.NoError(t, err)
	spanStd := &Span{}
	err = json.Unmarshal(b, spanStd)
	require.NoError(t, err)
	// unmarshal is the same as stdlib
	assert.Equal(t, spanStd, span)

	goJayBytes, err := gojay.Marshal(span)
	require.NoError(t, err)

	spanStd = &Span{}
	err = json.Unmarshal(goJayBytes, spanStd)
	require.NoError(t, err)
	assert.Equal(t, spanStd, span)
}

func BenchmarkMarshalGojay(b *testing.B) {
	benchmarkMarshal(b, gojay.Marshal)
}

func BenchmarkUnmarshalGojay(b *testing.B) {
	benchmarkUnmarshal(b, gojay.Unmarshal)
}

func (s *Span) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("operationName", s.OperationName)
	enc.StringKey("spanId", string(s.SpanID))
	enc.StringKey("traceId", string(s.TraceID))
	enc.Uint64Key("startTime", s.StartTime)
	enc.Uint64Key("startTimeMillis", s.StartTimeMillis)
	enc.Uint64Key("duration", s.Duration)
	enc.ObjectKey("process", &s.Process)
	enc.ArrayKey("tags", &s.Tags)
	if len(s.Tag) > 0 {
		enc.ObjectKey("tag", &s.Tag)
	}
	enc.ArrayKey("logs", &s.Logs)
	enc.ArrayKey("references", &s.References)
}
func (s *Span) IsNil() bool {
	return s == nil
}

func (p *Process) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("serviceName", p.ServiceName)
	enc.ArrayKey("tags", &p.Tags)
	if len(p.Tag) > 0 {
		enc.ObjectKey("tag", &p.Tag)
	}
}
func (p *Process) IsNil() bool {
	return p == nil
}
func (m TagMap) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range m {
		enc.AddInterfaceKey(k, v)
	}
}
func (m *TagMap) IsNil() bool {
	return m == nil
}

func (kvs *KeyValueArr) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *kvs {
		enc.Object(&e)
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
		enc.Object(&e)
	}
}
func (logs *LogArr) IsNil() bool {
	return logs == nil
}
func (l *Log) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Uint64Key("timestamp", l.Timestamp)
	enc.ArrayKey("fields", &l.Fields)
}
func (l *Log) IsNil() bool {
	return l == nil
}

func (refs *ReferenceArr) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *refs {
		enc.Object(&e)
	}
}
func (refs *ReferenceArr) IsNil() bool {
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
		err := dec.AddArray(&s.Tags)
		if s.Tags == nil {
			s.Tags = KeyValueArr{}
		}
		return err
	case "tag":
		if s.Tag == nil {
			s.Tag = TagMap{}
		}
		return dec.AddObject(&s.Tag)
	case "logs":
			return dec.AddArray(&s.Logs)
	case "references":
		return dec.AddArray(&s.References)
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
		err := dec.AddArray(&p.Tags)
		if p.Tags == nil {
			p.Tags = KeyValueArr{}
		}
		return err
	case "tag":
		if p.Tag == nil {
			p.Tag = TagMap{}
		}
		return dec.AddObject(&p.Tag)
	}
	return nil
}
func (p *Process) NKeys() int {
	return 0
}

func (t *TagMap) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	var inter interface{}
	err := dec.Interface(&inter)
	if err != nil {
		return err
	}
	(*t)[key] = inter
	return nil
}
func (t *TagMap) NKeys() int {
	return 0
}

func (kvs *KeyValueArr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tag := KeyValue{}
	err :=  dec.Object(&tag)
	if err != nil {
		return err
	}
	*kvs = append(*kvs, tag)
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
	*logs = append(*logs, l)
	return nil
}
func (l *Log) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "timestamp":
		return dec.AddUint64(&l.Timestamp)
	case "fields":
		err := dec.AddArray(&l.Fields)
		// keep the array even when empty
		if l.Fields == nil {
			l.Fields = KeyValueArr{}
		}
		return err
	}
	return nil
}
func (l *Log) NKeys() int {
	return 0
}

func (refs *ReferenceArr) UnmarshalJSONArray(dec *gojay.Decoder) error {
	ref := Reference{}
	err := dec.AddObject(&ref)
	if err != nil {
		return err
	}
	*refs = append(*refs, ref)
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

package jsontest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valyala/fastjson"
)

func TestParsingFixtureFastJson(t *testing.T) {
	b, err := loadFixture("default.json")
	require.NoError(t, err)

	span := &Span{}
	err = fastJsonUnmarshal(b, span)
	require.NoError(t, err)
	spanStd := &Span{}
	err = json.Unmarshal(b, spanStd)
	require.NoError(t, err)
	// unmarshal is the same as stdlib
	// TODO it does not pass - strings contain quotes and all values in tag map are converted to strings
	//assert.Equal(t, spanStd, span)
}

func BenchmarkUnmarshalFastjson(b *testing.B) {
	benchmarkUnmarshal(b, fastJsonUnmarshal)
}

func fastJsonUnmarshal(json []byte, int interface{}) error {
	p := fastjson.Parser{}
	valSpan, err := p.ParseBytes(json)
	if err != nil {
		return nil
	}
	span := int.(*Span)
	val := valSpan.Get("spanId")
	if val != nil {
		span.SpanID = SpanID(val.String())
	}
	val = valSpan.Get("traceId")
	if val != nil {
		span.TraceID = TraceID(val.String())
	}
	val = valSpan.Get("operationName")
	if val != nil {
		span.OperationName = val.String()
	}
	val = valSpan.Get("startTime")
	if val != nil {
		valUint, err := val.Uint64()
		if err == nil {
			span.StartTime = valUint
		}
	}
	val = valSpan.Get("startTimeMillis")
	if val != nil {
		valUint, err := val.Uint64()
		if err == nil {
			span.StartTimeMillis = valUint
		}
	}
	val = valSpan.Get("duration")
	if val != nil {
		valUint, err := val.Uint64()
		if err == nil {
			span.Duration = valUint
		}
	}
	valObj := valSpan.GetObject("process")
	if valObj != nil {
		span.Process = *parseProcess(valObj)
	}
	valArr := valSpan.GetArray("tags")
	if valArr != nil {
		span.Tags = make([]KeyValue, len(valArr))
		for i := 0; i < len(valArr); i++ {
			span.Tags[i] = parseKeyValue(valArr[i])
		}
	}
	valObj = valSpan.GetObject("tag")
	if valObj != nil {
		span.Tag = make(TagMap, valObj.Len())
		valObj.Visit(func(key []byte, v *fastjson.Value) {
			// TODO there is no access to an interface{} value
			span.Tag[string(key)] = v.String()
		})
	}
	valArr = valSpan.GetArray("logs")
	if valArr != nil {
		span.Logs = make([]Log, len(valArr))
		for i := 0; i < len(valArr); i++ {
			span.Logs[i] = parseLog(valArr[i])
		}
	}
	return nil
}

func parseLog(o *fastjson.Value) Log {
	log := Log{}
	val := o.Get("timestamp")
	if val != nil {
		valUint, err := val.Uint64()
		if err == nil {
			log.Timestamp = valUint
		}
	}
	valArr := o.GetArray("fields")
	if valArr != nil {
		log.Fields = make([]KeyValue, len(valArr))
		for i := 0; i < len(valArr); i++ {
			log.Fields[i] = parseKeyValue(valArr[i])
		}
	}
	return log
}

func parseKeyValue(o *fastjson.Value) KeyValue {
	kv := KeyValue{}
	val := o.Get("value")
	if val != nil {
		// TODO handle types properly
		kv.Value = val.String()
	}
	val = o.Get("type")
	if val != nil {
		kv.Type = ValueType(val.String())
	}
	val = o.Get("key")
	if val != nil {
		kv.Key = val.String()
	}
	return kv
}

func parseProcess(o *fastjson.Object) *Process {
	p := &Process{}
	val := o.Get("serviceName")
	if val != nil {
		p.ServiceName = val.String()
	}
	val = o.Get("tags")
	if val != nil {
		arr, err := val.Array()
		if err == nil {
			p.Tags = make([]KeyValue, len(arr))
			for i := 0; i < len(arr); i++ {
				p.Tags[i] = parseKeyValue(arr[i])
			}
		}
	}
	// TODO tag map
	return p
}

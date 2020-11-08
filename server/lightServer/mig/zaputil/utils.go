package zaputil

import (
	"go.uber.org/zap/zapcore"
	"net/http"
)

type ZapStrings []string

func (s ZapStrings) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	for _, v := range s {
		encoder.AppendString(v)
	}
	return nil
}

type ZapHeader http.Header

func (s ZapHeader) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	for k, v := range s {
		// ZapStrings은 에러 없음
		_ = encoder.AddArray(k, ZapStrings(v))
	}
	return nil
}

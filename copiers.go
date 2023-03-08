package copier

import (
	"errors"
	"strconv"
	"time"
)

var (
	NotMatchErr       = errors.New("src type not matching")
	FmtDateTime       = "2006-01-02 15:04:05"
	Int64       int64 = 0
	// 类型转换规则
	tcs = []TypeConverter{
		{
			SrcType: time.Time{},
			DstType: Int64,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(time.Time); ok {
					return s.Unix(), nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: time.Time{},
			DstType: String,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(time.Time); ok {
					return s.Format(FmtDateTime), nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: Int64,
			DstType: time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(int64); ok {
					time := time.Unix(s, 0)
					return time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: String,
			DstType: time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(string); ok {
					time, _ := time.Parse(FmtDateTime, s)
					return time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: Int64,
			DstType: &time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(int64); ok {
					time := time.Unix(s, 0)
					return &time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: String,
			DstType: &time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(string); ok {
					time, _ := time.Parse(FmtDateTime, s)
					return &time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: String,
			DstType: Int64,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(string); ok {
					i, _ := strconv.ParseInt(s, 10, 64)
					return i, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: Int64,
			DstType: String,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(int64); ok {
					return strconv.FormatInt(s, 10), nil
				}
				return nil, NotMatchErr
			},
		},
	}
)

// Copiers Set default conversion rules：
func Copiers(toValue interface{}, fromValue interface{}) error {
	return CopyWithOption(toValue, fromValue, Option{
		IgnoreEmpty: true,
		DeepCopy:    false,
		Converters:  tcs,
	})
}

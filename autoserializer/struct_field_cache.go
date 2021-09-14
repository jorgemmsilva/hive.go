package autoserializer

import (
	"fmt"
	"go/types"
	"reflect"
	"strconv"
	"sync"
	"unicode"
)

type fieldCache struct {
	lock             sync.Mutex
	structFieldCache map[reflect.Type][]FieldMetadata
}

// NewFieldCache creates and returns new fieldCache.
func NewFieldCache() *fieldCache {
	return &fieldCache{
		structFieldCache: make(map[reflect.Type][]FieldMetadata),
	}
}

type FieldMetadata struct {
	idx              int
	unpack           bool
	LengthPrefixType types.BasicKind
	MaxSliceLength   int
	MinSliceLength   int
	AllowNil         bool
	LexicalOrder     bool
}

// GetFields returns struct fields that are available for serialization. It caches the fields so consecutive calls for the same time can use previously extracted values.
func (c *fieldCache) GetFields(structType reflect.Type) ([]FieldMetadata, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.structFieldCache == nil {
		c.structFieldCache = make(map[reflect.Type][]FieldMetadata)
	}
	if cachedFields, ok := c.structFieldCache[structType]; ok {
		return cachedFields, nil
	}
	numFields := structType.NumField()
	cachedFields := make([]FieldMetadata, 0, numFields)
	for i := 0; i < numFields; i++ {
		field := structType.Field(i)
		var sm FieldMetadata
		sm.LengthPrefixType = types.Uint8
		switch field.Tag.Get("serialize") {
		case "unpack":
			sm.unpack = true
			fallthrough
		case "true":
			if !sm.unpack && unicode.IsLower(rune(field.Name[0])) {
				return nil, fmt.Errorf("can't marshal un-exported field '%s'", structType.Field(i).Name)
			}
			sm.idx = i

			if v := field.Tag.Get("minLen"); v != "" {
				minLen, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				sm.MinSliceLength = minLen
			}

			if v := field.Tag.Get("maxLen"); v != "" {
				maxLen, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				sm.MaxSliceLength = maxLen
			}

			if v := field.Tag.Get("allowNil"); v != "" {
				allowNil, err := strconv.ParseBool(v)
				if err != nil {
					return nil, err
				}
				sm.AllowNil = allowNil
			}
			if v := field.Tag.Get("lexicalOrder"); v != "" {
				lexicalOrder, err := strconv.ParseBool(v)
				if err != nil {
					return nil, err
				}
				sm.LexicalOrder = lexicalOrder
			}
			if v := field.Tag.Get("lenPrefixBytes"); v != "" {
				prefixBytes, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				switch prefixBytes {
				case 1:
					sm.LengthPrefixType = types.Uint8
				case 2:
					sm.LengthPrefixType = types.Uint16
				case 4:
					sm.LengthPrefixType = types.Uint32
				}
			}
			cachedFields = append(cachedFields, sm)

		}
	}
	c.structFieldCache[structType] = cachedFields
	return cachedFields, nil
}

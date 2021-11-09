package serializer

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"
)

// Serializable is something which knows how to serialize/deserialize itself from/into bytes.
// This is almost analogous to BinaryMarshaler/BinaryUnmarshaler.
type Serializable interface {
	json.Marshaler
	json.Unmarshaler
	// Deserialize deserializes the given data (by copying) into the object and returns the amount of bytes consumed from data.
	// If the passed data is not big enough for deserialization, an error must be returned.
	// During deserialization additional validation may be performed if the given modes are set.
	Deserialize(data []byte, deSeriMode DeSerializationMode) (int, error)
	// Serialize returns a serialized byte representation.
	// This function does not check the serialized data for validity.
	// During serialization additional validation may be performed if the given modes are set.
	Serialize(deSeriMode DeSerializationMode) ([]byte, error)
}

// Serializables is a slice of Serializable.
type Serializables []Serializable

// SerializableSlice is a slice of a type which can convert itself to Serializables.
type SerializableSlice interface {
	// ToSerializables returns the representation of the slice as a Serializables.
	ToSerializables() Serializables
	// FromSerializables updates the slice itself with the given Serializables.
	FromSerializables(seris Serializables)
}

// SerializableSelectorFunc is a function that given a type byte, returns an empty instance of the given underlying type.
// If the type doesn't resolve, an error is returned.
type SerializableSelectorFunc func(ty uint32) (Serializable, error)

// DeSerializationMode defines the mode of de/serialization.
type DeSerializationMode byte

const (
	// DeSeriModeNoValidation instructs de/serialization to perform no validation.
	DeSeriModeNoValidation DeSerializationMode = 0
	// DeSeriModePerformValidation instructs de/serialization to perform validation.
	DeSeriModePerformValidation DeSerializationMode = 1 << 0
	// DeSeriModePerformLexicalOrdering instructs de/deserialization to perform ordering of certain struct arrays by their lexical serialized form.
	DeSeriModePerformLexicalOrdering DeSerializationMode = 1 << 1
)

// HasMode checks whether the de/serialization mode includes the given mode.
func (sm DeSerializationMode) HasMode(mode DeSerializationMode) bool {
	return sm&mode > 0
}

// ArrayValidationMode defines the mode of array validation.
type ArrayValidationMode byte

const (
	// ArrayValidationModeNone instructs the array validation to perform no validation.
	ArrayValidationModeNone ArrayValidationMode = 0
	// ArrayValidationModeNoDuplicates instructs the array validation to check for duplicates.
	ArrayValidationModeNoDuplicates ArrayValidationMode = 1 << 0
	// ArrayValidationModeLexicalOrdering instructs the array validation to check for lexical order.
	ArrayValidationModeLexicalOrdering ArrayValidationMode = 1 << 1
	// ArrayValidationModeAtMostOneOfEachTypeByte instructs the array validation to allow a given byte type to occur only once in the array.
	ArrayValidationModeAtMostOneOfEachTypeByte ArrayValidationMode = 1 << 2
	// ArrayValidationModeAtMostOneOfEachTypeUint32 instructs the array validation to allow a given uint32 type to occur only once in the array.
	ArrayValidationModeAtMostOneOfEachTypeUint32 ArrayValidationMode = 1 << 3
)

// HasMode checks whether the array element validation mode includes the given mode.
func (av ArrayValidationMode) HasMode(mode ArrayValidationMode) bool {
	return av&mode > 0
}

// ArrayRules defines rules around a to be deserialized array.
// Min and Max at 0 define an unbounded array.
type ArrayRules struct {
	// The min array bound.
	Min uint
	// The max array bound.
	Max uint
	// The mode of validation.
	ValidationMode ArrayValidationMode
}

// ToWrittenObjectConsumer wraps this ArrayRules's ElementValidationFunc (according to ValidationMode) to a WrittenObjectConsumer.
// If the passed in mode has no validation, the returned WrittenObjectConsumer is a no-op.
func (ar *ArrayRules) ToWrittenObjectConsumer(mode DeSerializationMode) WrittenObjectConsumer {
	if !mode.HasMode(DeSeriModePerformValidation) {
		return func(index int, written []byte) error { return nil }
	}
	return WrittenObjectConsumer(ar.ElementValidationFunc(ar.ValidationMode))
}

// CheckBounds checks whether the given count violates the array bounds.
func (ar *ArrayRules) CheckBounds(count uint) error {
	if ar.Min != 0 && count < ar.Min {
		return fmt.Errorf("%w: min is %d but count is %d", ErrArrayValidationMinElementsNotReached, ar.Min, count)
	}
	if ar.Max != 0 && count > ar.Max {
		return fmt.Errorf("%w: max is %d but count is %d", ErrArrayValidationMaxElementsExceeded, ar.Max, count)
	}
	return nil
}

// ElementValidationFunc is a function which runs during array validation (e.g. lexical ordering).
type ElementValidationFunc func(index int, next []byte) error

// ElementUniqueValidator returns an ElementValidationFunc which returns an error if the given element is not unique.
func (ar *ArrayRules) ElementUniqueValidator() ElementValidationFunc {
	set := map[string]int{}
	return func(index int, next []byte) error {
		k := string(next)
		if j, has := set[k]; has {
			return fmt.Errorf("%w: element %d and %d are duplicates", ErrArrayValidationViolatesUniqueness, j, index)
		}
		set[k] = index
		return nil
	}
}

// LexicalOrderValidator returns an ElementValidationFunc which returns an error if the given byte slices
// are not ordered lexicographically.
func (ar *ArrayRules) LexicalOrderValidator() ElementValidationFunc {
	var prev []byte
	var prevIndex int
	return func(index int, next []byte) error {
		switch {
		case prev == nil:
			prev = next
			prevIndex = index
		case bytes.Compare(prev, next) > 0:
			return fmt.Errorf("%w: element %d should have been before element %d", ErrArrayValidationOrderViolatesLexicalOrder, index, prevIndex)
		default:
			prev = next
			prevIndex = index
		}
		return nil
	}
}

// LexicalOrderWithoutDupsValidator returns an ElementValidationFunc which returns an error if the given byte slices
// are not ordered lexicographically or any elements are duplicated.
func (ar *ArrayRules) LexicalOrderWithoutDupsValidator() ElementValidationFunc {
	var prev []byte
	var prevIndex int
	return func(index int, next []byte) error {
		if prev == nil {
			prev = next
			prevIndex = index
			return nil
		}
		switch bytes.Compare(prev, next) {
		case 1:
			return fmt.Errorf("%w: element %d should have been before element %d", ErrArrayValidationOrderViolatesLexicalOrder, index, prevIndex)
		case 0:
			// dup
			return fmt.Errorf("%w: element %d and %d are duplicates", ErrArrayValidationViolatesUniqueness, index, prevIndex)
		}
		prev = next
		prevIndex = index
		return nil
	}
}

// AtMostOneOfEachTypeValidator returns an ElementValidationFunc which returns an error if a given type occurs multiple
// times within the array.
func (ar *ArrayRules) AtMostOneOfEachTypeValidator(typeDenotation TypeDenotationType) ElementValidationFunc {
	seen := map[uint32]int{}
	return func(index int, next []byte) error {
		var key uint32
		switch typeDenotation {
		case TypeDenotationUint32:
			if len(next) < UInt32ByteSize {
				return fmt.Errorf("%w: not enough bytes to check type uniquness in array", ErrInvalidBytes)
			}
			key = binary.LittleEndian.Uint32(next)
		case TypeDenotationByte:
			if len(next) < OneByte {
				return fmt.Errorf("%w: not enough bytes to check type uniquness in array", ErrInvalidBytes)
			}
			key = uint32(next[0])
		default:
			panic(fmt.Sprintf("unknown type denotation in AtMostOneOfEachTypeValidator passed: %d", typeDenotation))
		}
		prevIndex, has := seen[key]
		if has {
			return fmt.Errorf("%w: element %d and %d have the same type", ErrArrayValidationViolatesTypeUniqueness, index, prevIndex)
		}
		seen[key] = index
		return nil
	}
}

// ElementValidationFunc returns a new ElementValidationFunc according to the given mode.
func (ar *ArrayRules) ElementValidationFunc(mode ArrayValidationMode) ElementValidationFunc {
	var arrayElementValidator ElementValidationFunc

	wrap := func(f ElementValidationFunc, f2 ElementValidationFunc) ElementValidationFunc {
		return func(index int, next []byte) error {
			if f != nil {
				if err := f(index, next); err != nil {
					return err
				}
			}
			return f2(index, next)
		}
	}

	for i := byte(1); i != 0; i <<= 1 {
		switch ArrayValidationMode(byte(mode) & i) {
		case ArrayValidationModeNone:
		case ArrayValidationModeNoDuplicates:
			if mode.HasMode(ArrayValidationModeLexicalOrdering) {
				continue
			}
			arrayElementValidator = wrap(arrayElementValidator, ar.ElementUniqueValidator())
		case ArrayValidationModeLexicalOrdering:
			// optimization: if lexical order and no dups are enforced, then byte comparison
			// to the previous element can be done instead of using a map
			if mode.HasMode(ArrayValidationModeNoDuplicates) {
				arrayElementValidator = wrap(arrayElementValidator, ar.LexicalOrderWithoutDupsValidator())
				continue
			}
			arrayElementValidator = wrap(arrayElementValidator, ar.LexicalOrderValidator())
		case ArrayValidationModeAtMostOneOfEachTypeByte:
			arrayElementValidator = wrap(arrayElementValidator, ar.AtMostOneOfEachTypeValidator(TypeDenotationByte))
		case ArrayValidationModeAtMostOneOfEachTypeUint32:
			arrayElementValidator = wrap(arrayElementValidator, ar.AtMostOneOfEachTypeValidator(TypeDenotationUint32))
		}
	}

	return arrayElementValidator
}

// LexicalOrderedByteSlices are byte slices ordered in lexical order.
type LexicalOrderedByteSlices [][]byte

func (l LexicalOrderedByteSlices) Len() int {
	return len(l)
}

func (l LexicalOrderedByteSlices) Less(i, j int) bool {
	return bytes.Compare(l[i], l[j]) < 0
}

func (l LexicalOrderedByteSlices) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// LexicalOrdered32ByteArrays are 32 byte arrays ordered in lexical order.
type LexicalOrdered32ByteArrays [][32]byte

func (l LexicalOrdered32ByteArrays) Len() int {
	return len(l)
}

func (l LexicalOrdered32ByteArrays) Less(i, j int) bool {
	return bytes.Compare(l[i][:], l[j][:]) < 0
}

func (l LexicalOrdered32ByteArrays) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// RemoveDupsAndSortByLexicalOrderArrayOf32Bytes returns a new SliceOfArraysOf32Bytes sorted by lexical order and without duplicates.
func RemoveDupsAndSortByLexicalOrderArrayOf32Bytes(slice SliceOfArraysOf32Bytes) SliceOfArraysOf32Bytes {

	seen := make(map[string]struct{})
	orderedArray := make(LexicalOrdered32ByteArrays, len(slice))

	uniqueElements := 0
	for _, v := range slice {
		k := string(v[:])
		if _, has := seen[k]; has {
			continue
		}
		seen[k] = struct{}{}
		orderedArray[uniqueElements] = v
		uniqueElements++
	}
	orderedArray = orderedArray[:uniqueElements]
	sort.Sort(orderedArray)

	return SliceOfArraysOf32Bytes(orderedArray)
}

// SortedSerializables are Serializables sorted by their serialized form.
type SortedSerializables Serializables

func (ss SortedSerializables) Len() int {
	return len(ss)
}

func (ss SortedSerializables) Less(i, j int) bool {
	iData, _ := ss[i].Serialize(DeSeriModeNoValidation)
	jData, _ := ss[j].Serialize(DeSeriModeNoValidation)
	return bytes.Compare(iData, jData) < 0
}

func (ss SortedSerializables) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

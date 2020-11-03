package rangemap

import (
	"fmt"
	"testing"

	"github.com/iotaledger/hive.go/datastructure/valuerange"
)

func Test(t *testing.T) {
	rangeMap := New(Int64Comparator)

	rangeMap.Set(valuerange.AtLeast(valuerange.Int64Value(9)), "hello9")
	rangeMap.Set(valuerange.AtLeast(valuerange.Int64Value(7)), "hello7")

	fmt.Println(rangeMap.tree)

	fmt.Println(rangeMap.Get(valuerange.Int64Value(8)))
	fmt.Println(rangeMap.Get(valuerange.Int64Value(11)))
	fmt.Println(rangeMap.Get(valuerange.Int64Value(3)))
}

func Int64Comparator(a interface{}, b interface{}) int {
	switch aCasted := a.(type) {
	case valuerange.Int64Value:
		switch bCasted := b.(type) {
		case valuerange.Int64Value:
			return aCasted.Compare(bCasted)
		case *valuerange.ValueRange:
			return -1 * bCasted.Compare(aCasted)
		}
	case *valuerange.ValueRange:
		switch bCasted := b.(type) {
		case valuerange.Int64Value:
			return aCasted.Compare(bCasted)
		case *valuerange.ValueRange:
			return GenericValueRangeComparator(aCasted, bCasted)
		default:
			panic(fmt.Sprintf("%v", bCasted))
		}
	default:
		panic(fmt.Sprintf("%v", aCasted))
	}

	return 0
}

func lowerEndPointsAreEqual(a *valuerange.ValueRange, b *valuerange.ValueRange) bool {
	if !a.HasLowerBound() && !b.HasLowerBound() {
		return true
	}

	if a.HasLowerBound() && b.HasLowerBound() {
		return a.LowerEndPoint().Value() == b.LowerEndPoint()
	}
}

func GenericValueRangeComparator(a *valuerange.ValueRange, b *valuerange.ValueRange) int {
	if !a.HasLowerBound() && !b.HasLowerBound() || (a.HasLowerBound() && b.HasLowerBound() && a.LowerEndPoint().Value() == b.LowerEndPoint().Value() && ) {
		if b.HasLowerBound() {
			return -1
		}

		if !a.HasUpperBound() {
			if !b.HasUpperBound() {
				return 0
			}
		}

		if
		// b doesn't have a lower bound either
	}

	if !aCasted.HasLowerBound() && !bCasted.HasLowerBound() {
		if !aCasted.HasUpperBound() && !bCasted.HasUpperBound() {
			return 0
		}

		if aCasted.HasUpperBound() && bCasted.HasUpperBound() && aCasted.UpperEndPoint().Value() == bCasted.UpperEndPoint().Value() && aCasted.UpperBoundType() == bCasted.UpperBoundType() {
			return 0
		}
	}

	if aCasted.HasLowerBound() && bCasted.HasLowerBound() && aCasted.LowerEndPoint().Value() == bCasted.LowerEndPoint().Value() && aCasted.LowerBoundType() == bCasted.LowerBoundType() {
		if !aCasted.HasUpperBound() && !bCasted.HasUpperBound() {
			return 0
		}

		if aCasted.HasUpperBound() && bCasted.HasUpperBound() && aCasted.UpperEndPoint().Value() == bCasted.UpperEndPoint().Value() && aCasted.UpperBoundType() == bCasted.UpperBoundType() {
			return 0
		}
	}

	if !aCasted.HasUpperBound() {

	}
	return aCasted.LowerEndPoint().Value().Compare(bCasted.LowerEndPoint().Value())
}

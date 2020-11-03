package rangemap

import (
	"github.com/iotaledger/hive.go/datastructure/genericcomparator"
	"github.com/iotaledger/hive.go/datastructure/redblacktree"
	"github.com/iotaledger/hive.go/datastructure/valuerange"
)

type RangeMap struct {
	tree *redblacktree.Tree
}

func New(comparator genericcomparator.Type) *RangeMap {
	return &RangeMap{
		tree: redblacktree.New(comparator),
	}
}

func (r *RangeMap) Set(key *valuerange.ValueRange, value interface{}) {
	node, inserted := r.tree.Set(key, value)
	if inserted {
		if predecessor := node.Predecessor(); predecessor != nil {
			predecessorValueRange := predecessor.Key().(*valuerange.ValueRange)
			if !key.HasLowerBound() {
				// remove all previous ranges + return
			}

			if predecessorValueRange.Compare(key.LowerEndPoint().Value()) == 0 {
				switch key.LowerEndPoint().BoundType() {
				case valuerange.BoundTypeOpen:
					predecessor.Key().(*valuerange.ValueRange).SetUpperEndPoint(valuerange.NewEndPoint(key.LowerEndPoint().Value(), valuerange.BoundTypeClosed))
				case valuerange.BoundTypeClosed:
					predecessor.Key().(*valuerange.ValueRange).SetUpperEndPoint(valuerange.NewEndPoint(key.LowerEndPoint().Value(), valuerange.BoundTypeOpen))
				}
			}
		}

		if successor := node.Successor(); successor != nil {
			successorValueRange := successor.Key().(*valuerange.ValueRange)
			if !key.HasUpperBound() || successorValueRange.Compare(key.UpperEndPoint().Value()) == 0 {
				switch successorValueRange.LowerEndPoint().BoundType() {
				case valuerange.BoundTypeOpen:
					key.SetUpperEndPoint(valuerange.NewEndPoint(successorValueRange.LowerEndPoint().Value(), valuerange.BoundTypeClosed))
				case valuerange.BoundTypeClosed:
					key.SetUpperEndPoint(valuerange.NewEndPoint(successorValueRange.LowerEndPoint().Value(), valuerange.BoundTypeOpen))
				}
			}
		}
	}
}

func (r *RangeMap) Get(key valuerange.Value) (interface{}, bool) {
	return r.tree.Get(key)
}

func (r *RangeMap) Size() int {
	return r.tree.Size()
}

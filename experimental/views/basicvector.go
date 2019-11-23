package views

import (
	"fmt"
	. "github.com/protolambda/zrnt/experimental/tree"
)

type BasicVectorType struct {
	ElementType BasicTypeDef
	Length      uint64
}

func (cd *BasicVectorType) DefaultNode() Node {
	depth := GetDepth(cd.BottomNodeLength())
	inner := &Commit{}
	inner.ExpandInplaceDepth(&ZeroHashes[0], depth)
	return inner
}

func (cd *BasicVectorType) ViewFromBacking(node Node) View {
	depth := GetDepth(cd.BottomNodeLength())
	return &BasicVectorView{
		SubtreeView: SubtreeView{
			BackingNode: node,
			depth:       depth,
		},
		BasicVectorType: cd,
	}
}

func (cd *BasicVectorType) ElementsPerBottomNode() uint64 {
	return 32 / cd.ElementType.ByteLength()
}

func (cd *BasicVectorType) BottomNodeLength() uint64 {
	perNode := cd.ElementsPerBottomNode()
	return (cd.Length + perNode - 1) / perNode
}

func (cd *BasicVectorType) TranslateIndex(index uint64) (nodeIndex uint64, intraNodeIndex uint8) {
	perNode := cd.ElementsPerBottomNode()
	return index / perNode, uint8(index & (perNode - 1))
}

func (cd *BasicVectorType) New() *BasicVectorView {
	return cd.ViewFromBacking(cd.DefaultNode()).(*BasicVectorView)
}

type BasicVectorView struct {
	SubtreeView
	*BasicVectorType
}

func (cv *BasicVectorView) ViewRoot(h HashFn) Root {
	return cv.BackingNode.MerkleRoot(h)
}

func (cv *BasicVectorView) subviewNode(i uint64) (r *Root, bottomIndex uint64, subIndex uint8, err error) {
	bottomIndex, subIndex = cv.TranslateIndex(i)
	v, err := cv.SubtreeView.Get(bottomIndex)
	if err != nil {
		return nil,  0, 0, err
	}
	r, ok := v.(*Root)
	if !ok {
		return nil, 0, 0, fmt.Errorf("basic vector bottom node is not a root, at index %d", i)
	}
	return r, bottomIndex, subIndex, nil
}

func (cv *BasicVectorView) Get(i uint64) (SubView, error) {
	if i >= cv.Length {
		return nil, fmt.Errorf("basic vector has length %d, cannot get index %d", cv.Length, i)
	}
	r, _, subIndex, err := cv.subviewNode(i)
	if err != nil {
		return nil, err
	}
	return cv.ElementType.SubViewFromBacking(r, subIndex), nil
}
func (cv *BasicVectorView) copyChunk(i uint64, offset uint8, dest []byte) error {
	v, err := cv.SubtreeView.Get(i)
	if err != nil {
		return err
	}
	r, ok := v.(*Root)
	if !ok {
		return fmt.Errorf("basic vector bottom node is not a root, at bottom node index %d", i)
	}
	copy(dest, r[offset:])
	return nil
}

func (cv *BasicVectorView) IntoBytes(skip uint64, dest []byte) error {
	startChunk, subStart := cv.TranslateIndex(skip)
	// copy over partial first chunk
	if subStart != 0 {
		if err := cv.copyChunk(startChunk, subStart, dest[startChunk<< 5 + uint64(subStart):]); err != nil {
			return err
		}
		startChunk += 1
	}
	endChunk, subEnd := cv.TranslateIndex(skip + uint64(len(dest)))
	// copy over full chunks
	for i := startChunk; i < endChunk; i++ {
		if err := cv.copyChunk(i, 0, dest[i << 5:(i + 1) << 5]); err != nil {
			return err
		}
	}
	// copy over partial last chunk
	if subEnd != 0 {
		if err := cv.copyChunk(endChunk, 0, dest[endChunk<< 5:endChunk<< 5 + uint64(subEnd)]); err != nil {
			return err
		}
	}
	return nil
}

func (cv *BasicVectorView) Set(i uint64, view SubView) error {
	if i >= cv.Length {
		return fmt.Errorf("cannot set item at element index %d, basic vector only has %d elements", i, cv.Length)
	}
	r, bottomIndex, subIndex, err := cv.subviewNode(i)
	if err != nil {
		return err
	}
	return cv.SubtreeView.Set(bottomIndex, view.BackingFromBase(r, subIndex))
}

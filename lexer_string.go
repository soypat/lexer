package lexer

import (
	"fmt"
	"strconv"
)

func (i Item) String() string {
	switch i.typ {
	case ItemEOF:
		return "EOF"
	case ItemError:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}


// Generated with https://godoc.org/golang.org/x/tools/cmd/stringer
// stringer -type=ItemType -output
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ItemError-0]
	_ = x[ItemNil-1]
	_ = x[ItemEOF-2]
	_ = x[ItemNumber-3]
	_ = x[ItemVar-4]
	_ = x[ItemOperator-5]
	_ = x[ItemFunc-6]
	_ = x[ItemCommaSep-7]
	_ = x[ItemColonOp-8]
	_ = x[ItemSemiSep-9]
	_ = x[ItemLeftFuncMeta-10]
	_ = x[ItemRightFuncMeta-11]
	_ = x[ItemLeftMatMeta-12]
	_ = x[ItemRightMatMeta-13]
	_ = x[ItemLeftIdxMeta-14]
	_ = x[ItemRightIdxMeta-15]
	_ = x[ItemLeftPemdas-16]
	_ = x[ItemRightPemdas-17]
	_ = x[ItemIdentifier-18]
	_ = x[ItemText-19]
	_ = x[itemEnd-20]
	_ = x[itemVarIdx-21]
	_ = x[itemString-22]
	_ = x[itemAnon-23]
	_ = x[itemIf-24]
	_ = x[itemElse-25]
}

const _itemType_name = "ItemErrorItemNilItemEOFItemNumberItemVarItemOperatorItemFuncItemCommaSepItemColonOpItemSemiSepItemLeftFuncMetaItemRightFuncMetaItemLeftMatMetaItemRightMatMetaItemLeftIdxMetaItemRightIdxMetaItemLeftPemdasItemRightPemdasItemIdentifierItemTextitemEnditemVarIdxitemStringitemAnonitemIfitemElse"

var _itemType_index = [...]uint16{0, 9, 16, 23, 33, 40, 52, 60, 72, 83, 94, 110, 127, 142, 158, 173, 189, 203, 218, 232, 240, 247, 257, 267, 275, 281, 289}

func (i ItemType) String() string {
	if i < 0 || i >= ItemType(len(_itemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _itemType_name[_itemType_index[i]:_itemType_index[i+1]]
}

// Code generated by thriftrw

package unions

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/gen/testdata/typedefs"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type ArbitraryValue struct {
	BoolValue   *bool
	Int64Value  *int64
	ListValue   []*ArbitraryValue
	MapValue    map[string]*ArbitraryValue
	StringValue *string
}
type _List_ArbitraryValue_ValueList []*ArbitraryValue

func (v _List_ArbitraryValue_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(x.ToWire())
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _List_ArbitraryValue_ValueList) Close() {
}

type _Map_String_ArbitraryValue_MapItemList map[string]*ArbitraryValue

func (m _Map_String_ArbitraryValue_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: wire.NewValueString(k), Value: v.ToWire()})
		if err != nil {
			return err
		}
	}
	return nil
}
func (m _Map_String_ArbitraryValue_MapItemList) Close() {
}
func (v *ArbitraryValue) ToWire() wire.Value {
	var fields [5]wire.Field
	i := 0
	if v.BoolValue != nil {
		fields[i] = wire.Field{ID: 1, Value: wire.NewValueBool(*(v.BoolValue))}
		i++
	}
	if v.Int64Value != nil {
		fields[i] = wire.Field{ID: 2, Value: wire.NewValueI64(*(v.Int64Value))}
		i++
	}
	if v.ListValue != nil {
		fields[i] = wire.Field{ID: 4, Value: wire.NewValueList(wire.List{ValueType: wire.TStruct, Size: len(v.ListValue), Items: _List_ArbitraryValue_ValueList(v.ListValue)})}
		i++
	}
	if v.MapValue != nil {
		fields[i] = wire.Field{ID: 5, Value: wire.NewValueMap(wire.Map{KeyType: wire.TBinary, ValueType: wire.TStruct, Size: len(v.MapValue), Items: _Map_String_ArbitraryValue_MapItemList(v.MapValue)})}
		i++
	}
	if v.StringValue != nil {
		fields[i] = wire.Field{ID: 3, Value: wire.NewValueString(*(v.StringValue))}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func _ArbitraryValue_Read(w wire.Value) (*ArbitraryValue, error) {
	var v ArbitraryValue
	err := v.FromWire(w)
	return &v, err
}
func _List_ArbitraryValue_Read(l wire.List) ([]*ArbitraryValue, error) {
	if l.ValueType != wire.TStruct {
		return nil, nil
	}
	o := make([]*ArbitraryValue, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := _ArbitraryValue_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}
func _Map_String_ArbitraryValue_Read(m wire.Map) (map[string]*ArbitraryValue, error) {
	if m.KeyType != wire.TBinary {
		return nil, nil
	}
	if m.ValueType != wire.TStruct {
		return nil, nil
	}
	o := make(map[string]*ArbitraryValue, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}
		v, err := _ArbitraryValue_Read(x.Value)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}
func (v *ArbitraryValue) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.BoolValue = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.Int64Value = &x
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TList {
				v.ListValue, err = _List_ArbitraryValue_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TMap {
				v.MapValue, err = _Map_String_ArbitraryValue_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.StringValue = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *ArbitraryValue) String() string {
	var fields [5]string
	i := 0
	if v.BoolValue != nil {
		fields[i] = fmt.Sprintf("BoolValue: %v", *(v.BoolValue))
		i++
	}
	if v.Int64Value != nil {
		fields[i] = fmt.Sprintf("Int64Value: %v", *(v.Int64Value))
		i++
	}
	if v.ListValue != nil {
		fields[i] = fmt.Sprintf("ListValue: %v", v.ListValue)
		i++
	}
	if v.MapValue != nil {
		fields[i] = fmt.Sprintf("MapValue: %v", v.MapValue)
		i++
	}
	if v.StringValue != nil {
		fields[i] = fmt.Sprintf("StringValue: %v", *(v.StringValue))
		i++
	}
	return fmt.Sprintf("ArbitraryValue{%v}", strings.Join(fields[:i], ", "))
}

type Document struct {
	Pdf       typedefs.Pdf
	PlainText *string
}

func (v *Document) ToWire() wire.Value {
	var fields [2]wire.Field
	i := 0
	if v.Pdf != nil {
		fields[i] = wire.Field{ID: 1, Value: v.Pdf.ToWire()}
		i++
	}
	if v.PlainText != nil {
		fields[i] = wire.Field{ID: 2, Value: wire.NewValueString(*(v.PlainText))}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func _Pdf_Read(w wire.Value) (typedefs.Pdf, error) {
	var x typedefs.Pdf
	err := x.FromWire(w)
	return x, err
}
func (v *Document) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Pdf, err = _Pdf_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.PlainText = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *Document) String() string {
	var fields [2]string
	i := 0
	if v.Pdf != nil {
		fields[i] = fmt.Sprintf("Pdf: %v", v.Pdf)
		i++
	}
	if v.PlainText != nil {
		fields[i] = fmt.Sprintf("PlainText: %v", *(v.PlainText))
		i++
	}
	return fmt.Sprintf("Document{%v}", strings.Join(fields[:i], ", "))
}

type EmptyUnion struct{}

func (v *EmptyUnion) ToWire() wire.Value {
	var fields [0]wire.Field
	i := 0
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func (v *EmptyUnion) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}
func (v *EmptyUnion) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("EmptyUnion{%v}", strings.Join(fields[:i], ", "))
}

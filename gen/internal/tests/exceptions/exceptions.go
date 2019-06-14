// Code generated by thriftrw v1.21.0-dev. DO NOT EDIT.
// @generated

package exceptions

import (
	errors "errors"
	fmt "fmt"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// Raised when something doesn't exist.
type DoesNotExistException struct {
	// Key that was missing.
	Key    string  `json:"key,required"`
	Error2 *string `json:"Error,omitempty"`
}

// ToWire translates a DoesNotExistException struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *DoesNotExistException) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Key), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Error2 != nil {
		w, err = wire.NewValueString(*(v.Error2)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a DoesNotExistException struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a DoesNotExistException struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v DoesNotExistException
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *DoesNotExistException) FromWire(w wire.Value) error {
	var err error

	keyIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Key, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				keyIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Error2 = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !keyIsSet {
		return errors.New("field Key of DoesNotExistException is required")
	}

	return nil
}

// String returns a readable string representation of a DoesNotExistException
// struct.
func (v *DoesNotExistException) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Key: %v", v.Key)
	i++
	if v.Error2 != nil {
		fields[i] = fmt.Sprintf("Error2: %v", *(v.Error2))
		i++
	}

	return fmt.Sprintf("DoesNotExistException{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this DoesNotExistException match the
// provided DoesNotExistException.
//
// This function performs a deep comparison.
func (v *DoesNotExistException) Equals(rhs *DoesNotExistException) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Key == rhs.Key) {
		return false
	}
	if !_String_EqualsPtr(v.Error2, rhs.Error2) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of DoesNotExistException.
func (v *DoesNotExistException) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("key", v.Key)
	if v.Error2 != nil {
		enc.AddString("Error", *v.Error2)
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *DoesNotExistException) GetKey() (o string) {
	if v != nil {
		o = v.Key
	}
	return
}

// GetError2 returns the value of Error2 if it is set or its
// zero value if it is unset.
func (v *DoesNotExistException) GetError2() (o string) {
	if v != nil && v.Error2 != nil {
		return *v.Error2
	}

	return
}

// IsSetError2 returns true if Error2 is not nil.
func (v *DoesNotExistException) IsSetError2() bool {
	return v != nil && v.Error2 != nil
}

func (v *DoesNotExistException) Error() string {
	return v.String()
}

type EmptyException struct {
}

// ToWire translates a EmptyException struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *EmptyException) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a EmptyException struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a EmptyException struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v EmptyException
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *EmptyException) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a EmptyException
// struct.
func (v *EmptyException) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("EmptyException{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this EmptyException match the
// provided EmptyException.
//
// This function performs a deep comparison.
func (v *EmptyException) Equals(rhs *EmptyException) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EmptyException.
func (v *EmptyException) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

func (v *EmptyException) Error() string {
	return v.String()
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "exceptions",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/exceptions",
	FilePath: "exceptions.thrift",
	SHA1:     "743daa9bfc5a3d69637e7c67dd6f35a7d10e79a3",
	Raw:      rawIDL,
}

const rawIDL = "exception EmptyException {}\n\n/**\n * Raised when something doesn't exist.\n */\nexception DoesNotExistException {\n    /** Key that was missing. */\n    1: required string key\n    2: optional string Error (go.name=\"Error2\")\n}\n"

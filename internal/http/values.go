package http

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type values map[string]string

var (
	errNoPtr  = errors.New("destination not a pointer")
	errNilPtr = errors.New("destination pointer is nil")
)

func requestVar(r *http.Request, k string, dest interface{}) error {
	return values(mux.Vars(r)).scan(k, dest)
}

// https://play.golang.org/p/hpRQ638XwZF
func (v values) scan(k string, dest interface{}) error {
	val, found := v[k]
	if !found {
		return fmt.Errorf("no key for %s", k)
	}

	dpv := reflect.ValueOf(dest)
	if dpv.Kind() != reflect.Ptr {
		return errNoPtr
	}
	if dpv.IsNil() {
		return errNilPtr
	}

	spv := reflect.ValueOf(val)

	// If the type of the value can be directly converted to the type of the
	// destination then perform the conversion and return. We're done here and
	// can go home for a lovely up of tea.
	dv := reflect.Indirect(dpv)
	if spv.Kind() == dv.Kind() && spv.Type().ConvertibleTo(dv.Type()) {
		dv.Set(spv.Convert(dv.Type()))
		return nil
	}

	switch d := dest.(type) {
	case *string:
		*d = val
	case *[]byte:
		*d = []byte(val)
	case *bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		*d = b
	case *int, *int8, *int16, *int32, *int64:
		i64, err := strconv.ParseInt(val, 10, dv.Type().Bits())
		if err != nil {
			return err
		}
		dv.SetInt(i64)
	case *uint, *uint8, *uint16, *uint32, *uint64:
		u64, err := strconv.ParseUint(val, 10, dv.Type().Bits())
		if err != nil {
			return err
		}
		dv.SetUint(u64)
	case *interface{}:
		*d = val
	}

	return nil
}

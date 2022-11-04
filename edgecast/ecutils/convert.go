package ecutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	jsonMarshal   = json.Marshal
	jsonUnmarshal = json.Unmarshal
)

// This function is used to convert a struct to the provided type.
// Dest must be a pointer otherwise an error is returned.
// Please refer to the json unmarshall documentation for additional information
// on expected behavior.
//
// Field names need to match between the source and destination structs. This
// function should be used wisely when the source and destination are well known.
func Convert(src, dest any) error {
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New("dest must be a pointer and not nil")
	}

	// Marshal src to json bytes
	srcBytes, err := jsonMarshal(src)
	if err != nil {
		return fmt.Errorf(
			"unable to marshal src: %s err: %w", src, err)
	}

	// Unmarshal into dest
	err = jsonUnmarshal(srcBytes, dest)
	if err != nil {
		return fmt.Errorf("failed to convert %T to %T: %w", src, dest, err)
	}
	return nil
}

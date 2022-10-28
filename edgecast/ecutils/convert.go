package ecutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// This function is used to convert a struct to the provided type.
// If dest must be a pointer otherwise an error is returned.
// Please refer to the json unmashall documentation for additional information
// on expected behavior.
//
// Field names need to match between the source and destination structs. This
// function should be used wisely when the source and destination are well known.
func Convert(src, dest any) error {
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New("dest must be a pointer")
	}

	// Marshal src to json bytes
	srcBytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf(
			"unable to marshal get origin group response: %w", err)
	}

	// Unmarshal into dest
	err = json.Unmarshal(srcBytes, dest)
	if err != nil {
		return fmt.Errorf("failed to convert %T to %T: %w", src, dest, err)
	}
	return nil
}

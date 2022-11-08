package ecutils

import (
	"encoding/json"
	"errors"
	"testing"
)

type marshal func(v any) ([]byte, error)
type unmarshal func(data []byte, v any) error

type srcStruct struct {
	ID   int
	Name string
}

type destStruct struct {
	ID   int
	Name string
}

func TestConvert(t *testing.T) {
	cases := []struct {
		Name          string
		SrcStruct     *srcStruct
		DestStruct    *destStruct
		ExpectedError bool
		Marshal       marshal
		Unmarshal     unmarshal
	}{
		{
			Name:          "Happy path",
			SrcStruct:     &srcStruct{ID: 1, Name: "First"},
			DestStruct:    &destStruct{},
			ExpectedError: false,
			Marshal:       json.Marshal,
			Unmarshal:     json.Unmarshal,
		},
		{
			Name:          "Nil src returns error",
			SrcStruct:     nil,
			DestStruct:    nil,
			ExpectedError: true,
			Marshal:       json.Marshal,
			Unmarshal:     json.Unmarshal,
		},
		{
			Name:          "Marshal failure returns error",
			SrcStruct:     &srcStruct{ID: 1, Name: "First"},
			DestStruct:    &destStruct{},
			ExpectedError: true,
			Marshal:       failJsonMarshal,
			Unmarshal:     json.Unmarshal,
		},
		{
			Name:          "Unmarshal failure returns error",
			SrcStruct:     &srcStruct{ID: 1, Name: "First"},
			DestStruct:    &destStruct{},
			ExpectedError: true,
			Marshal:       jsonMarshal,
			Unmarshal:     failJsonUnmarshal,
		},
	}

	for _, v := range cases {
		jsonMarshal = v.Marshal
		jsonUnmarshal = v.Unmarshal

		err := Convert(v.SrcStruct, v.DestStruct)

		if !v.ExpectedError && err != nil {
			t.Fatalf("Failed for case: '%+v'. Error not expected, but got `%s`", v.Name, err)
		}

		if v.ExpectedError && err == nil {
			t.Fatalf("Failed for case: '%+v'. Expected an error, but did not get one", v.Name)
		}

		if v.ExpectedError && err != nil {
			t.Logf("Received expected error: %v", err)
			continue
		}

		if !SrcEqualsDest(*v.SrcStruct, *v.DestStruct) {
			t.Fatalf("Failed for case: '%+v'. Expected '%+v', but got '%+v'", v.Name, *v.SrcStruct, *v.DestStruct)
		}
	}

}

func SrcEqualsDest(source srcStruct, dest destStruct) bool {
	if (source.ID != dest.ID) || (source.Name != dest.Name) {
		return false
	}
	return true
}

func failJsonMarshal(v any) ([]byte, error) {
	return []byte{}, errors.New("marshal failure")
}

func failJsonUnmarshal(data []byte, v any) error {
	return errors.New("unmarshal failure")
}

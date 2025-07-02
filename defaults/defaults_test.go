package defaults

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taodev/pkg/types"
)

const YamlText = `
a: 123
field_slice_binary: !!binary SGVsbG8sIFdvcmxkIQ==
`

func TestDefault(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		t1 := struct {
			FieldInt         int            `default:"1"`
			FieldIntPtr      *int           `default:"2"`
			FieldString      string         `default:"2"`
			FieldStringPtr   *string        `default:"3234234"`
			FieldDuration    time.Duration  `default:"1s"`
			FieldDurationPtr *time.Duration `default:"3s"`

			FieldTime    time.Time  `default:"2023-01-01T00:00:00Z"`
			FieldTimePtr *time.Time `default:"2023-01-02T00:00:00Z"`

			FieldSkip int `default:"-"`
		}{}
		require.NoError(t, Set(&t1))
		assert.Equal(t, 1, t1.FieldInt)
		assert.Equal(t, 2, *t1.FieldIntPtr)
		assert.Equal(t, "2", t1.FieldString)
		assert.Equal(t, "3234234", *t1.FieldStringPtr)
		assert.Equal(t, time.Second, t1.FieldDuration)
		assert.Equal(t, time.Second*3, *t1.FieldDurationPtr)
		assert.Equal(t, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), t1.FieldTime)
		assert.Equal(t, time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), *t1.FieldTimePtr)
		assert.Equal(t, 0, t1.FieldSkip)
	})

	t.Run("Slice", func(t *testing.T) {
		t1 := struct {
			FieldByte   []byte       `default:"[1, 2, 3]"`
			FieldBinary types.Binary `default:"!!binary SGVsbG8gV29ybGQh"`

			FieldSliceInt       []int     `default:"[1, 2, 3]"`
			FieldSliceIntPtr    *[]int    `default:"[1, 2, 3]"`
			FieldSliceString    []string  `default:"['123', '234', '345']"`
			FieldSliceStringPtr *[]string `default:"['123', '234', '345']"`

			FieldSliceNil   []int `default:""`
			FieldSliceEmpty []int `default:"[]"`
		}{}
		require.NoError(t, Set(&t1))
		assert.Equal(t, []byte{1, 2, 3}, t1.FieldByte)
		assert.Equal(t, types.Binary("Hello World!"), t1.FieldBinary)
		assert.Equal(t, []int{1, 2, 3}, t1.FieldSliceInt)
		assert.Equal(t, []int{1, 2, 3}, *t1.FieldSliceIntPtr)
		assert.Equal(t, []string{"123", "234", "345"}, t1.FieldSliceString)
		assert.Equal(t, []string{"123", "234", "345"}, *t1.FieldSliceStringPtr)

		assert.Nil(t, t1.FieldSliceNil)
		assert.Equal(t, []int{}, t1.FieldSliceEmpty)
	})

	t.Run("Map", func(t *testing.T) {
		t1 := struct {
			FieldMapInt       map[string]int          `default:"{a: 1, b: 2}"`
			FieldMapIntPtr    map[string]*int         `default:"{a: 1, b: 2}"`
			FieldMapString    map[string]string       `default:"{a: '123', b: '234'}"`
			FieldMapStringPtr map[string]*string      `default:"{a: '123', b: '234'}"`
			FieldMapBinary    map[string]types.Binary `default:"{a: SGVsbG8gV29ybGQh, b: !!binary SGVsbG8gV29ybGQh}"`
			FieldMapNil       map[string]int          `default:""`
			FieldMapEmpty     map[string]int          `default:"{}"`
		}{}
		require.NoError(t, Set(&t1))
		assert.Equal(t, map[string]int{"a": 1, "b": 2}, t1.FieldMapInt)
		mapIntPtr := make(map[string]*int)
		mapIntPtr["a"] = new(int)
		mapIntPtr["b"] = new(int)
		*mapIntPtr["a"] = 1
		*mapIntPtr["b"] = 2
		assert.Equal(t, mapIntPtr, t1.FieldMapIntPtr)

		assert.Equal(t, map[string]string{"a": "123", "b": "234"}, t1.FieldMapString)
		mapStringPtr := make(map[string]*string)
		mapStringPtr["a"] = new(string)
		mapStringPtr["b"] = new(string)
		*mapStringPtr["a"] = "123"
		*mapStringPtr["b"] = "234"
		assert.Equal(t, mapStringPtr, t1.FieldMapStringPtr)

		assert.Equal(t, map[string]types.Binary{"a": []byte("Hello World!"), "b": []byte("Hello World!")}, t1.FieldMapBinary)

		assert.Nil(t, t1.FieldMapNil)
		assert.Equal(t, map[string]int{}, t1.FieldMapEmpty)
	})

	t.Run("Struct", func(t *testing.T) {
		type SubStruct struct {
			A int `default:"1"`
			B int `default:"25"`
			C int
			D *int
		}

		t1 := struct {
			FieldStruct struct {
				A int `default:"1"`
			} `default:"{a: 2}"`
			FieldStructPtr *struct {
				A int `default:"1"`
			} `default:"{a: 3}"`

			FieldSubStruct          SubStruct `default:"{a: 4}"`
			FieldSubStructNoDefault SubStruct

			FieldSubStructNil   *SubStruct `default:""`
			FieldSubStructEmpty *SubStruct `default:"{}"`
		}{}
		require.NoError(t, Set(&t1))
		assert.Equal(t, 2, t1.FieldStruct.A)

		assert.Equal(t, 3, t1.FieldStructPtr.A)

		assert.Equal(t, 4, t1.FieldSubStruct.A)
		assert.Equal(t, 25, t1.FieldSubStruct.B)
		assert.Equal(t, 0, t1.FieldSubStruct.C)

		assert.Equal(t, 1, t1.FieldSubStructNoDefault.A)
		assert.Equal(t, 25, t1.FieldSubStructNoDefault.B)
		assert.Equal(t, 0, t1.FieldSubStructNoDefault.C)

		assert.Nil(t, t1.FieldSubStructNil)
		assert.Equal(t, &SubStruct{A: 1, B: 25}, t1.FieldSubStructEmpty)
	})

	t.Run("SliceStruct", func(t *testing.T) {
		type SubStruct struct {
			A int `default:"1"`
			B int `default:"25"`
			C int
		}

		t1 := struct {
			FieldSliceStruct    []SubStruct  `default:"[{a: 4}]"`
			FieldSliceStructPtr []*SubStruct `default:"[{a: 5}]"`
		}{}
		require.NoError(t, Set(&t1))
		assert.Equal(t, 4, t1.FieldSliceStruct[0].A)
		assert.Equal(t, 25, t1.FieldSliceStruct[0].B)
		assert.Equal(t, 0, t1.FieldSliceStruct[0].C)

		assert.Equal(t, 5, t1.FieldSliceStructPtr[0].A)
		assert.Equal(t, 25, t1.FieldSliceStructPtr[0].B)
		assert.Equal(t, 0, t1.FieldSliceStructPtr[0].C)
	})

	t.Run("MapStruct", func(t *testing.T) {
		type SubStruct struct {
			A int `default:"1"`
			B int `default:"25"`
			C int
		}

		t1 := struct {
			FieldMapStruct    map[string]SubStruct  `default:"{a: {a: 4}}"`
			FieldMapStructPtr map[string]*SubStruct `default:"{a: {a: 4}}"`

			FieldMapSlice    map[string][]SubStruct  `default:"{a: [{a: 4}]}"`
			FieldMapSlicePtr map[string][]*SubStruct `default:"{a: [{a: 4}]}"`

			FieldMapMap    map[string]map[string]SubStruct  `default:"{a: {a: {a: 4}}}"`
			FieldMapMapPtr map[string]map[string]*SubStruct `default:"{a: {a: {a: 4}}}"`
		}{}
		require.NoError(t, Set(&t1))
		assert.Equal(t, 4, t1.FieldMapStruct["a"].A)
		assert.Equal(t, 25, t1.FieldMapStruct["a"].B)
		assert.Equal(t, 0, t1.FieldMapStruct["a"].C)

		assert.Equal(t, 4, t1.FieldMapStructPtr["a"].A)
		assert.Equal(t, 25, t1.FieldMapStructPtr["a"].B)
		assert.Equal(t, 0, t1.FieldMapStructPtr["a"].C)

		assert.Equal(t, 4, t1.FieldMapSlice["a"][0].A)
		assert.Equal(t, 25, t1.FieldMapSlice["a"][0].B)
		assert.Equal(t, 0, t1.FieldMapSlice["a"][0].C)

		assert.Equal(t, 4, t1.FieldMapSlicePtr["a"][0].A)
		assert.Equal(t, 25, t1.FieldMapSlicePtr["a"][0].B)
		assert.Equal(t, 0, t1.FieldMapSlicePtr["a"][0].C)

		assert.Equal(t, 4, t1.FieldMapMap["a"]["a"].A)
		assert.Equal(t, 25, t1.FieldMapMap["a"]["a"].B)
		assert.Equal(t, 0, t1.FieldMapMap["a"]["a"].C)

		assert.Equal(t, 4, t1.FieldMapMapPtr["a"]["a"].A)
		assert.Equal(t, 25, t1.FieldMapMapPtr["a"]["a"].B)
		assert.Equal(t, 0, t1.FieldMapMapPtr["a"]["a"].C)
	})
}

func TestDefaultsError(t *testing.T) {
	t.Run("TypeError", func(t *testing.T) {
		var t1 int
		assert.Equal(t, errInvalidType, Set(t1))
		assert.Equal(t, errInvalidType, Set(&t1))
	})

	t.Run("setFieldError", func(t *testing.T) {
		var canSetVal map[string]struct{ canSet bool } = map[string]struct{ canSet bool }{
			"a": {canSet: true},
		}
		assert.NoError(t, setField(reflect.ValueOf(canSetVal), ""))
	})

	t.Run("yamlError", func(t *testing.T) {
		type subStruct struct {
			A int `default:"adfsdf sdfsdf"`
		}

		t1 := struct {
			A int `default:"sdf sdkfjsdkf"`
		}{}
		assert.Error(t, Set(&t1))

		t2 := struct {
			A subStruct `default:"{}"`
		}{}
		assert.Error(t, Set(&t2))

		t3 := struct {
			A []subStruct `default:"[{}]"`
		}{}
		assert.Error(t, Set(&t3))

		t4 := struct {
			A map[string]subStruct `default:"{a: {}}"`
		}{}
		assert.Error(t, Set(&t4))

		t5 := struct {
			A map[string]*subStruct `default:"{a: {}}"`
		}{}
		assert.Error(t, Set(&t5))
	})
}

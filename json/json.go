package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"../fs"
	"reflect"
)

type JSON struct {
	data interface{}
}

/**
字符串转换成interface{} 对象赋值给data
*/
func (j *JSON) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(p))
	dec.UseNumber()
	return dec.Decode(&j.data)
}

func ReadFile(fileName string) (*JSON, error) {
	bs, err := fs.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return Byte2Json(bs)

}
func Byte2Json(content []byte) (*JSON, error) {
	j := new(JSON)
	err := j.UnmarshalJSON(content)
	if err != nil {
		return nil, err
	}

	return j, nil
}
func String2Json(content string) (*JSON, error) {
	return Byte2Json([]byte(content))
}
func (j *JSON) SetData(m map[string]interface{}) {
	j.data = m
}

func (j *JSON) Map() (map[string]interface{}, error) {
	if j == nil {
		return nil, errors.New("type  assertion to map[string]interface{} failed")
	}
	m, ok := (j.data).(map[string]interface{})
	if ok == true {
		return m, nil
	}
	return nil, errors.New("type  assertion to map[string]interface{} failed")
}

func (j *JSON) Array() ([]interface{}, error) {
	if j == nil {
		return nil, errors.New("type  assertion to []interface{} failed")
	}
	a, ok := (j.data).([]interface{})
	if ok {
		return a, nil
	}
	return nil, errors.New("type  assertion to []interface{} failed")
}

func (j *JSON) GetIndex(key int) *JSON {
	m, err := j.Array()
	if err == nil {
		if len(m) > key {
			return &JSON{m[key]}
		}
	}
	return nil

}

func (j *JSON) Get(key string) *JSON {
	m, err := j.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &JSON{val}
		}
	}
	return nil
}

func (j *JSON) Del(key string) error {
	m, err := j.Map()
	if err != nil {
		return err
	}
	delete(m, key)
	return nil
}

/**
@method Set
设置
*/
func (j *JSON) Set(key string, value interface{}) error {
	if j == nil {
		return errors.New("invalid value type")
	}
	m, err := j.Map()
	if err != nil {
		return err
	}
	m[key] = value
	return nil

}

/**
@method Int
将data转换成int型
*/
func (j *JSON) Int() (int, error) {
	if j == nil {
		return 0, errors.New("invalid value type")
	}

	switch j.data.(type) {
	case json.Number:
		i, err := j.data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(j.data).Float()), nil
		break
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(j.data).Int()), nil
		break
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(j.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

func (j *JSON) String() (string, error) {
	if j == nil {
		return "", errors.New("invalid value type")
	}
	b, err := json.Marshal(j.data)
	if err != nil {
		return "", err
	}
	return string(b), nil
	/**
 	targetValue := reflect.ValueOf(target)
    switch reflect.TypeOf(target).Kind() {
		case reflect.Slice, reflect.Array, reflect.:
			b, err := json.Marshal(m)
			if err != nil {
				return "", err
			}
			return string(b)
		break;

	}
	*/
}

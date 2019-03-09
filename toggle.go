package jct

import "encoding/json"

func Toggle(j []byte, from, to Case) ([]byte, error) {

	var data interface{}
	err := json.Unmarshal(j, &data)
	if err != nil {
		return nil, err
	}

	err = toggle(data, from, to)
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

func toggle(data interface{}, from, to Case) (err error) {

	switch typed := data.(type) {
	case map[string]interface{}:
		for k, v := range typed {
			fixed := fixKey(k, from, to)
			delete(typed, k)
			typed[fixed] = v
			err = toggle(v, from, to)
			if err != nil {
				return err
			}
		}
	case []interface{}:
		for _, v := range typed {
			err = toggle(v, from, to)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func fixKey(key string, from Case, to Case) string {
	return to.Join(from.Split(key))
}

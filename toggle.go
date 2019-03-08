package jct

import "encoding/json"

func Toggle(j json.RawMessage, from, to Case) (msg json.RawMessage, err error) {

	var data interface{}
	if err := json.Unmarshal([]byte(j), &data); err != nil {
		return nil, err
	}

	err = toggle(data, from, to)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(b), nil
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

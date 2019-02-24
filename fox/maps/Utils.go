package maps

//如果指定的key不存在于指定的Map中 则返回指定的默认值
func GetOrDefault(m map[interface{}]interface{}, key interface{}, defaultValue interface{}) interface{} {
	if val, ok := m[key]; !ok || val == nil {
		return defaultValue
	}
	return m[key]
}

func Clear(m map[interface{}]interface{}) {
	for k := range m {
		delete(m, k)
	}
}

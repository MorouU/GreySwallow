package utils

// Init
var (
	GLOBALMAP map[string]interface{} = make(map[string]interface{})
)

// Set GLOBALMAP
func Set(key string, value interface{}, force bool) bool {
	if _, ok := GLOBALMAP[key]; !ok && !force {
		return false
	}
	GLOBALMAP[key] = value
	return true
}

// Get GLOBALMAP
func Get(key string) interface{} {

	if v, ok := GLOBALMAP[key]; ok {
		return v
	}
	return nil
}

package dft

func String(value, defValue string) string {
	if value == "" {
		return defValue
	}
	return value
}
// wtf
//func Bool(value, defValue bool) bool {
//	if value == false {
//		return defValue
//	}
//	return value
//}

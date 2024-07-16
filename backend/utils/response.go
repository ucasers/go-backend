package utils

func ResponseData(obj interface{}, modelName string) interface{} {
	switch modelName {
	case "User":
		return structToMap(obj, "ID", "Password")
	case "Extension":
		return structToMap(obj)
	case "CipherPair":
		return structToMap(obj, "Owner", "OwnerID")
	}
	return nil
}

package main

// Arguments to format are: [1]: type name
const valueMethod = `func (i %[1]s) Value() (driver.Value, error) {
	return int32(i), nil
}
`

const scanMethod = `func (i *%[1]s) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var val int64
	var err error

	switch v := value.(type) {
	case int64:
		val = v
	case int32:
        val = int64(v)
	case []byte:
		val, err = strconv.ParseInt(string(v), 10, 32)
		if err != nil {
			return fmt.Errorf("cannot scan '%%v' into int32: %%w", value, err)
		}
	case string:
		val, err = strconv.ParseInt(v, 10, 32)
		if err != nil {
			return fmt.Errorf("cannot scan '%%v' into int32: %%w", value, err)
		}
	default:
		return fmt.Errorf("invalid value of %[1]s: %%[1]T(%%[1]v)", value)
	}

	if v:=%[1]s(int32(val)); v.IsA%[1]s() {
		*i = v
	}

	return fmt.Errorf("unknown %[1]s type")
}
`

func (g *Generator) addValueAndScanMethod(typeName string) {
	g.Printf("\n")
	g.Printf(valueMethod, typeName)
	g.Printf("\n\n")
	g.Printf(scanMethod, typeName)
}

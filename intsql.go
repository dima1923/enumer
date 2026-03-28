package main

// Arguments to format are: [1]: type name
const intValueMethod = `func (i %[1]s) Value() (driver.Value, error) {
	return int64(i), nil
}
`

const intScanMethod = `
var err%[1]sNilPtr = fmt.Errorf("nil pointer")

func (i *%[1]s) Scan(value interface{}) (err error) {
	if value == nil {
		*i = %[1]s(0)
		return
	}

	switch v := value.(type) {
	case int64:
		*i = %[1]s(v)
	case string:
		var val int
		// try parsing the integer value as a string
		val, err = strconv.Atoi(v)	
		if err == nil {
			*i = %[1]s(val)
		}
	case []byte:
		var val int
		val, err = strconv.Atoi(string(v))
		if err == nil {
			*i = %[1]s(val)
		}
	case %[1]s:
		*i = v
	case int:
		*i = %[1]s(v)
	case *%[1]s:
		if v == nil {
			return err%[1]sNilPtr
		}
		*i = *v
	case uint:
		*i = %[1]s(v)
	case uint64:
		*i = %[1]s(v)
	case *int:
		if v == nil {
			return err%[1]sNilPtr
		}
		*i = %[1]s(*v)
	case *int64:
		if v == nil {
			return err%[1]sNilPtr
		}
		*i = %[1]s(*v)
	case float64:
		*i = %[1]s(v)
	case *float64:
		if v == nil {
			return err%[1]sNilPtr
		}
		*i = %[1]s(*v)
	case *uint:
		if v == nil {
			return err%[1]sNilPtr
		}
		*i = %[1]s(*v)
	case *uint64:
		if v == nil {
			return err%[1]sNilPtr
		}
		*i = %[1]s(*v)
	case *string:
		if v == nil {
			return err%[1]sNilPtr
		}
		var val int
		// try parsing the integer value as a string
		val, err = strconv.Atoi(*v)
		if err == nil {
			*i, err = %[1]s(val), nil
		}
	}

	return
}
`

func (g *Generator) addIntValueAndScanMethod(typeName string) {
	g.Printf("\n")
	g.Printf(intValueMethod, typeName)
	g.Printf("\n\n")
	g.Printf(intScanMethod, typeName)
}

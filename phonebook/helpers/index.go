package helpers

var index map[string]int

func CreateIndex() error {

	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
	return nil

}

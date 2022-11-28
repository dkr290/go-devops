package appliances

type Microwave struct {
	typeName string
}

func (fr *Microwave) Start() {
	fr.typeName = "Microwave"
}

func (fr *Microwave) GetPurpose() string {
	return "I am a " + fr.typeName + " I cook food"
}

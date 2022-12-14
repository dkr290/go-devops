package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcidsGram float64
type SodiumMilliGram float64
type FruitsPercent float64
type FibreGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium              SodiumMilliGram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

func (e EnergyKJ) GetPoints(st ScoreType) int {

}

func (s SugarGram) GetPoints(st ScoreType) int {

}

func (sfa SaturatedFattyAcidsGram) GetPoints(st ScoreType) int {

}

func (s SodiumMilliGram) GetPoints(st ScoreType) int {

}

func (f FruitsPercent) GetPoints(st ScoreType) int {

}

func (f FibreGram) GetPoints(st ScoreType) int {

}

func (p ProteinGram) GetPoints(st ScoreType) int {

}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilliGram {
	return SodiumMilliGram(saltMg / 2.5)
}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water {

		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fibre.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints(st)
	}

	return NutritionalScore{

		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

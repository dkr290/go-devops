package main

import "fmt"

func main() {

	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilligram(),
		Fruits:              FruitsPercent(),
		Fiber:               FibreGram(),
		Protein:             ProteinGram(),
	}, Food)

	fmt.Printf("Nutritional Score:%d\n", ns.Value)
}

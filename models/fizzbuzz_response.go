package models

//La structure FizzBuzzResponse contient un champ Result qui est un tableau de chaînes.
//Ce champ sera utilisé pour stocker
//la séquence de nombres résultante avec les remplacements fizz-buzz correspondants.

type FizzBuzzResponse struct {
	Result []string `json:"result"`
}

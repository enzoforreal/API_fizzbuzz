package models

//La structure FizzBuzzResponse contient un champ Result qui est un tableau de chaînes.
//Ce champ sera utilisé pour stocker
//la séquence de nombres résultante et leurs remplacement par fizzbuzz

type FizzBuzzResponse struct {
	Result []string `json:"result"`
}

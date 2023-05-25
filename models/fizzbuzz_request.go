package models

//La structure FizzBuzzRequest définit les paramètres de requête pour le fizz-buzz,
//tels que Int1 (premier entier), Int2 (deuxième entier),
//Limit (limite de la séquence de nombres),
//Str1 (première chaîne de remplacement) et Str2 (deuxième chaîne de remplacement).

type FizzBuzzRequest struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

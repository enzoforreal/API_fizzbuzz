API FizzBuzz
Version: 1.0.0

L'API FizzBuzz permet de générer des séquences FizzBuzz et de collecter des statistiques sur les requêtes effectuées.

Utilisation
L'API FizzBuzz propose deux endpoints :

FizzBuzz: Génère une séquence FizzBuzz en utilisant les paramètres spécifiés.

URL : /fizzbuzz
Méthode : GET
Paramètres de requête :
int1 (entier) : Le premier entier pour le FizzBuzz.
int2 (entier) : Le deuxième entier pour le FizzBuzz.
limit (entier) : La limite de la séquence FizzBuzz.
str1 (chaîne) : La chaîne à afficher pour les multiples de int1.
str2 (chaîne) : La chaîne à afficher pour les multiples de int2.
Réponse : Un tableau JSON contenant la séquence FizzBuzz générée.
Statistics: Récupère les statistiques sur les requêtes FizzBuzz effectuées.

URL : /statistics
Méthode : GET
Réponse : Un objet JSON contenant les statistiques des requêtes FizzBuzz, comprenant la requête la plus utilisée et le décompte des requêtes effectuées.
Exemple d'utilisation
Exemple de requête FizzBuzz
URL : http://localhost:8000/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz

Réponse :

json
Copy code
{
"result": [
"1",
"2",
"fizz",
"4",
"buzz",
"fizz",
"7",
"8",
"fizz",
"buzz",
"11",
"fizz",
"13",
"14",
"fizzbuzz"
]
}
Exemple de requête Statistics
URL : http://localhost:8000/statistics

Réponse :

json
Copy code
{
"MostUsedRequest": {
"int1": 3,
"int2": 5,
"limit": 15,
"str1": "fizz",
"str2": "buzz"
},
"RequestCount": [
{
"Request": {
"int1": 3,
"int2": 5,
"limit": 15,
"str1": "fizz",
"str2": "buzz"
},
"Count": 10
},
{
"Request": {
"int1": 4,
"int2": 6,
"limit": 10,
"str1": "hello",
"str2": "world"
},
"Count": 1
}
]
}


Environnement de développement
Pour exécuter localement l'API FizzBuzz, vous devez avoir les prérequis suivants installés :
Go 1.20 

Roadmap
Dans les versions futures, je prévois d'ajouter les fonctionnalités suivantes :

Interface utilisateur frontend avec React pour utiliser facilement l'API FizzBuzz.
Graphiques de statistiques pour une meilleure visualisation des données.


Auteurs
 RENAUD

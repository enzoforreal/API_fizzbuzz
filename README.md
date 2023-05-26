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
Fonctionnalités du frontend
Le frontend de l'API FizzBuzz a été développé avec React et offre les fonctionnalités suivantes :

Interface utilisateur conviviale pour interagir avec l'API FizzBuzz.
Génération et affichage des séquences FizzBuzz à partir des paramètres spécifiés.
Affichage des statistiques des requêtes FizzBuzz, y compris la requête la plus utilisée et le décompte des requêtes.

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

 Démarrage  la partie front pour tester l'application visuellement :

Assurez-vous d'avoir Node.js installé sur votre machine.

Ouvrez une nouvelle fenêtre de terminal.

Accédez au répertoire racine de votre projet frontend (où se trouve le fichier package.json).

Exécutez la commande suivante pour installer les dépendances nécessaires :
npm install

Une fois l'installation terminée, exécutez la commande suivante pour démarrer le serveur de développement :
npm start

Le serveur de développement démarrera et vous fournira une URL (par exemple, http://localhost:3000).

Ouvrez votre navigateur et accédez à l'URL fournie.

Vous devriez voir l'application frontend FizzBuzz, avec la possibilité d'interagir avec l'API en générant des séquences FizzBuzz et en affichant les statistiques.

Vous pouvez utiliser les formulaires et les boutons de l'interface utilisateur pour effectuer des requêtes FizzBuzz et voir les résultats.

Assurez-vous également que le serveur backend de l'API FizzBuzz est en cours d'exécution pour que l'application frontend puisse communiquer avec l'API.
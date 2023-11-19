# GolangLearningExercice


Exercice 1 : Manipulation de Map en Go

Contexte: Vous travaillez sur une application de gestion de mots et de définitions en utilisant Go. Le code actuel utilise une map en mémoire pour stocker les entrées.

Tâches:
	
Créez une fonction main dans le fichier principal (main.go) pour réaliser les tâches suivantes:
	
Utilisez la méthode Add pour ajouter quelques mots et définitions à la map.
		
Utilisez la méthode Get pour afficher la définition d'un mot spécifique.
		
Utilisez la méthode Remove pour supprimer un mot de la map.
		
Appelez la méthode List pour obtenir la liste triée des mots et de leurs définitions.
	
Exécutez le programme et vérifiez si les opérations sur la map sont correctement implémentées.



Exercice 2 : Création d'un Package Dictionary en Go

Contexte: Vous poursuivez le développement de votre application de gestion de mots et de définitions. Pour rendre le code plus modulaire, vous allez créer un package dictionary contenant les méthodes Add, Get, Remove, et List.

Tâches:

Créez un nouveau fichier nommé dictionary.go dans le répertoire dictionary de votre projet.
	
Déplacez les méthodes Add, Get, Remove, et List dans le package dictionary.
	
Importez et utilisez ce package dans le fichier main.go pour effectuer les opérations sur la map.


Instructions:

Créez un fichier dictionary.go dans le répertoire dictionary.
	
Copiez les méthodes Add, Get, Remove, et List de main.go vers dictionary.go.
	
Modifiez le fichier main.go pour importer le package dictionary et utilisez les méthodes du package pour réaliser les opérations sur la map.
	
Exécutez le programme et assurez-vous que tout fonctionne comme prévu.



Exercice3 : Gestion de Données avec Fichier en Go

Contexte: Vous avez implémenté un dictionnaire en utilisant une map dans un exercice précédent. Maintenant, vous allez modifier votre implémentation pour stocker les données dans un fichier plutôt que dans une map.

Tâches:
	
Modifiez le package dictionary pour utiliser un fichier au lieu d'une map pour stocker les entrées du dictionnaire.
	
Utilisez les méthodes Add, Get, Remove, et List du package dictionary dans main.go.
	
Assurez-vous que les opérations sur les données fonctionnent correctement après ces modifications.


Instructions:
	
Modifiez le code dans dictionary.go pour utiliser un fichier (au format de votre choix) au lieu d'une map pour stocker les entrées du dictionnaire.
	
Adaptez les méthodes Add, Get, Remove, et List en conséquence.
	
Testez les opérations dans main.go pour garantir que tout fonctionne correctement.


Consignes supplémentaires:
	
Ajoutez et validez (git add, git commit) vos modifications pour chaque étape.
	
Poussez (git push) régulièrement vos modifications sur GitHub.
	
Assurez-vous que votre programme fonctionne correctement avec les données stockées dans un fichier.
# SingleMarket

SingleMarket est une application de gestion de marché permettant de gérer des produits, des clients, des commandes, et d'exporter des données sous forme de fichiers CSV et PDF.

## Fonctionnalités

- Ajouter, afficher, modifier et supprimer des produits
- Ajouter, afficher et modifier des clients
- Passer des commandes et générer des confirmations de commande au format PDF
- Envoyer des confirmations de commande par e-mail aux clients
- Exporter les produits, clients et commandes sous forme de fichiers CSV

## Prérequis

- [Go](https://golang.org/dl/) 1.16 ou supérieur
- [Docker](https://www.docker.com/get-started) (pour utiliser MailHog)
- Un serveur MySQL

## Installation

1. Clonez le dépôt :

   ```bash
   git clone https://github.com/samuelaudic/singlemarket.git
   cd singlemarket
   ```

2. Installez les dépendances :

   ```bash
   go mod download
   ```

3. Configurez la base de données MySQL :

   - Assurez-vous que MySQL est installé et en cours d'exécution.
   - Créez une base de données pour l'application.
   - Mettez à jour les informations de connexion à la base de données dans `database/database.go`.

4. Lancez MailHog avec Docker :

   ```bash
   docker-compose up -d
   ```

## Utilisation

1. Compilez et exécutez l'application :

   ```bash
   go run main.go
   ```

2. Suivez les instructions du menu pour ajouter des produits, des clients, passer des commandes, et exporter des données.

## Exportation des données

- Les produits peuvent être exportés sous forme de fichiers CSV dans le répertoire `./exports/csv`.
- Les confirmations de commande sont générées sous forme de fichiers PDF dans le répertoire `./exports/pdf`.
- Les commandes peuvent être exportées sous forme de fichiers CSV groupés par client, avec un tableau des commandes effectuées et la somme totale de leurs commandes.

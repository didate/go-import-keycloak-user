# Keycloak CSV User Importer

This project is a Go-based tool for importing users from a CSV file into Keycloak. It reads a list of users from a CSV file, creates them in a specified Keycloak realm, and sets their passwords. The tool uses the [GoCloak](https://github.com/Nerzal/gocloak) library to interact with the Keycloak API.

## Features

- Load user data from a CSV file.
- Create users in a Keycloak realm.
- Set passwords for the users in Keycloak.
- Supports user realm roles.

## Prerequisites

- Go 1.19 or higher
- Keycloak instance running
- Users CSV file (see format below)
- A `.env` file to store Keycloak configuration

## CSV File Format

The CSV file should contain the following columns in order: username,email,firstName,lastName,roles

### Example:

```
user1,test1@example.com,John,Doe,role1,role2
user2,test2@example.com,Jane,Doe,role2
```

## Setup

### 1. Clone the Repository

First, clone the repository to your local machine using Git:

```
git clone https://github.com/yourusername/keycloak-csv-importer.git
cd keycloak-csv-importer
```

### 2. Create a .env File
Create a .env file in the root of the project. This file will contain your Keycloak credentials and other configurations. The contents of the file should look like this:
```
KC_BASEURL=https://your-keycloak-instance.com
KC_REALM=your-realm
KC_USERNAME=your-admin-username
KC_PASSWORD=your-admin-password
```
Replace the values with your actual Keycloak setup.

### 3. Install Dependencies

Make sure you have Go installed (version 1.19 or higher). Then, install the project dependencies by running:
```
go mod tidy
```
This will download and install all required packages.

## Run the Program

To execute the program, simply run:
```
go run main.go
```

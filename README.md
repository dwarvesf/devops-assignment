# Devops Assignment

## Prerequisite
* docker
* docker-compose
* make

## Running Locally

To run the program locally, install Docker. After docker is installed run:
```
# we need to run 3rd party first (e.g mysql) and seed data before running backend 
$ make init

# run backend
$ make backend-up

# run frontend
$ make frontend-up
```
from the root directory. Then we can access the web on browser http://localhost:3000/
```
$ open http://localhost:3000/
```

## What the Application Does

The application display the latest workplace accident and keep counting of how many days since the last time we faced it

## Parts of Application

1. Mysql which contains table `accidents (id, content, date)`
2. Backend is written in Golang to get data from database and return to the frontend. Backend runs on port 1323 as default
3. Frontend is written in Node, using `express` to implement. Frontend runs on port 3000 as default

## Your Job

The purpose of this exercise is to demonstrate how to setup the infrastructure in a deployed AWS or GCP environment using code. You can use whatever tools you want to do this (Ansible, Puppet, Terraform, etc.). The system is setup to run locally using Docker as an example. You need to use Docker as part of the way the application is deployed.

A few caveats for how the application should be deployed:
* Frontend is publicly available
* Backend Service is a microservice that is called by Frontend but is not publically accessible
* MySQL database should not be publically accessible
* Frontend & Backend need to be accessed via `https` protocal
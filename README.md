# go-demo

## launch up App

- docker volume create go-demo-db
- docker-compose -f db.docker-compose.yml up
## controller 

- Managment of the REST interafce to business logic
- Route will be setup in controller/index.go

## service

- Business logic implementations
- Currently, it only has one serivce, once it grwoing up, will be more than one.

## repo

- Storage of the entity beans in the system

## dto

- represent of the request object
- represent of the response object

## test

- e2e
  - end to end test, simple put, integration test.
## reference 

- https://tom-collings.medium.com/controller-service-repository-16e29a4684e5
- https://www.coreycleary.me/what-is-the-difference-between-controllers-and-services-in-node-rest-apis
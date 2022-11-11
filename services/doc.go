/*
1.every service mentioned behaves like a microservice
2.Inter services calls are not allowed due to the hazard of circular dependencies
3.Every service inherits the services definitions- log, conf etc.. there has them common

Naming Conventions

1. don't export service, encapsulate the methods under services and export them

*/

package services

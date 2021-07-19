# WebService

A simple web service getting driver duration and distance calculation from OSMR API.
Make sure you have to Go version 1.14 or newer

## Code Architecture
It has three packages

* <b>Rest</b> package is reponsible for routing and running http server
* <b> Service </b> package has two services API and Route
  * <b>API</b> service is responsible for getting data from OSRM API
  * <b> Route </b> service is responsible getting data from api service and do some business logic and sends back to routes results to handlers 

# How to run?

```
make init
make build
./bin/server -port 3000

```

The application has already been deployed to Heroku!
<a href="https://polar-brushlands-62667.herokuapp.com/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219">Link to the application</a>

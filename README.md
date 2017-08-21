
## Salvation Army Private Api ##
_Simple, Awesome_


### Structure ###
* The API is composed of 3 services:

              - orders service
              - customer service
              - partner service

* Each of these services is self contained with:

             - logging and
             - instrumentation

* Each service has a test endpoint:
    - Orders service

          - /v1/add parameters {"a":1,"b":4}

    - Customer service

          - /v1/test parameters {"s":"Customer"}

    - Partner service

          - /v1/test parameters {"s":"Partner"}

    ** Each of these endpoints accepts a post request.

* There is also a metrics endpoint which shows the statistics of
the API at.

            - /metrics

    ** this endpoint accepts get requests

* To get the authorization token:

                  - /v1/auth parameters {"clientId": "mobile", "clientSecret": "m_secret"}

             ** set the token received to the Authorizatioln Headers of each subsequent request.



### Installing ###
* cd $GOPATH/src/
* mkdir sendyit
* cd sendyit
* mkdir private_api
* cd private_api
* git clone git@gitlab.com:sendy/go_private_api.git
* Run glide install
* Run  go run *.go
* docker build -t private_api:v1 .
* docker run -p 8000:8000 private_api:v1




### To do ###
* Write test cases for the services and endpoints
* Write structs for custom response objects
* implements services for the apis

            - school service
            - auth service
            - Performance service
            - infrastructure service
            - project service
            - activities service
            - upload service






* log in as user, then use the id to fetch the token - {"clientId":id}
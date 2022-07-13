# Say Who I Am

Web app for writing reviews of people.

## Installation:
```go get github.com/phanvanpeter/say-who-i-am```

## REST examples:
#### GET PEOPLE
``curl localhost:9090/people -v | jq``

#### GET PERSON (id is obligatory)
``curl localhost:9090/people/1 -v | jq``

#### POST PEOPLE
``curl localhost:9090/people -v -X POST -d '{"firstName": "Mike", "lastName":"Spirit", "email": "mike@spirit.com", "age": 54, "stars": 2.5}'``

#### PUT PERSON (id is obligatory)
``curl localhost:9090/people -v -X PUT -d '{"id": 3, "firstName": "Mike", "lastName":"Jackson", "email": "mike@jackson.com", "age": 53, "stars": 4.5}'``

#### DELETE PERSON (id is obligatory)
``curl localhost:9090/people/3 -v -X DELETE``

#### POST IMAGE (id is obligatory)
``curl -v localhost:9090/images/1/test.png -X POST --data-binary @coffee.png``

#### GET IMAGE (id is obligatory)
``curl -v localhost:9090/images/1/coffee.png``

#### GET REVIEWS
``curl localhost:9090/reviews -v | jq``

#### GET REVIEW (id is obligatory)
``curl localhost:9090/reviews/12 -v | jq``

## License
MIT licensed. See the LICENSE file for details.
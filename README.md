## Quick start

- `make docker-run-postgres`
- `make migrate` - Database migration
- `make server` to run the api at port 8080

## Endpoints

- Create a team 

```
curl -X POST -H 'Content-Type: application/json' --data '{"name": "Vultures" }' localhost:8080/teams | jq
```

- Get teams

``` 
curl -X GET localhost:8080/teams | jq

```

- Get a team

``` 
curl -X GET localhost:8080/teams/1 | jq

```

- Create a player

```
curl -X POST -H 'Content-Type: application/json' --data '{"name":"Yan","number":34}' localhost:8080/players | jq

```

- Get players

``` 
curl -X GET localhost:8080/players | jq

```

- Get a player

``` 
curl -X GET localhost:8080/players/1 | jq

```

- Add a player to a team

```
curl -X PUT -H 'Content-Type: application/json' --data '{"playerIds":[1]}' localhost:8080/teams/1 | jq

```


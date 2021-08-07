# REST API GOLANG
Simple REST APIs about Pokemon Build with Golang you can CRUD this APIs freely 😃.
# BUILD WITH
- MySQL
- Gin
- GORM v2

# HOW TO RUN PROJECT ?
- Import database pokemon.sql at folder config
- run this project using <b>go run main.go</b>

# GET
- [api/v1/pokemons](https://pokemon-api-go.herokuapp.com/api/v1/pokemons) (retrive all pokemons data)
- [api/v1/pokemons/:id](https://pokemon-api-go.herokuapp.com/api/v1/pokemons/1)  (retrive specific pokemon data)
- Response :

```json 
{
  "status": 200,
  "data": [
      {
          "id": 2,
          "name": "ivysaur",
          "height": 10,
          "weight": 130,
          "base_experience": 142,
          "types": [
              {
                  "id": 4,
                  "name": "poison",
                  "damage_type_id": 4
              },
              {
                  "id": 12,
                  "name": "grass",
                  "damage_type_id": 12
              }
          ]
      }
  ]
}
```

# POST
- [api/v1/pokemons](https://pokemon-api-go.herokuapp.com/api/v1/pokemons) (create new pokemon)
- Body :

```json 
{
  "name": "Felix",
  "height": 170,
  "weight": 70,
  "base_experience": 100
}
```

- Response :

```json 
{
  "status": 201,
  "message": "created",
  "data": {
      "id": 730,
      "name": "Felix",
      "height": 170,
      "weight": 70,
      "base_experience": 100,
      "types": null
  },
  "created_at": "Saturday, 31-Jul-21 20:32:44 +07"
}
```
# PATCH 
- [api/v1/pokemons](https://pokemon-api-go.herokuapp.com/api/v1/pokemons/730) (update pokemon)
- Body :

```json 
{
  "name": "Fernando",
  "height": 171,
  "weight": 71,
  "base_experience": 101
}
```

- Response :

```json 
{
  "status": 201,
  "message": "updated",
  "data": {
      "id": 730,
      "name": "Felix",
      "height": 170,
      "weight": 70,
      "base_experience": 100,
      "types": null
  },
  "updated_at": "Saturday, 31-Jul-21 20:32:44 +07"
}
```

# DELETE
- [api/v1/pokemons/:id](https://pokemon-api-go.herokuapp.com/api/v1/pokemons/700) (delete pokemon)
- Response :

```json 
{
  "status": 200,
  "message": "deleted",
  "deleted_at": "Saturday, 31-Jul-21 20:52:46 +07",
}
```
# REST API GOLANG
Simple REST APIs about Pokemon Build with Golang you can CRUD this APIs freely 😃.
# BUILD WITH
- MySQL
- Gin
- GORM

# HOW TO RUN PROJECT ?
- Import database pokemon.sql at folder config
- run project using <b>go run main.go</b>

# GET
- [api/v1/pokemons](https://pokemon-api-go.herokuapp.com/api/v1/pokemons) (retrive all pokemons data)
- [api/v1/pokemons/:id](https://pokemon-api-go.herokuapp.com/api/v1/pokemons/1)  (retrive specific pokemon data)

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

Response :

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
- TBA

# DELETE
- TBA
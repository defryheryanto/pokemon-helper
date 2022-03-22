# Pokemon Helper
Your pokemon helper, here you can get pokemon infos, and simulate your team.

# Features
## Get Pokemon Data ```/api/v1/pokemon/{pokemonName}```
Get pokemon data by name
Example:
Response:
```
{
    "name": "Dragonite",
    "base_status": {
        "hp": 91,
        "attack": 134,
        "defense": 95,
        "special_attack": 100,
        "special_defense": 100,
        "speed": 80,
        "total": 600
    },
    "types": [
        "Dragon",
        "Flying"
    ]
}
```
## Simulate Team ```/api/v1/simulate-team```
Return pokemon types which will be super_effective when attacked by the team pokemon. You can turn on the type suggestion by sending 'with_type_suggestion' true in the payload.
Example:
Payload
```
{
    "pokemons": [
        "Typhlosion",
        "gengar",
        "dragonite",
        "gyarados",
        "tyranitar"
    ],
    "with_type_suggestion": true,
    "with_pokemon_data": true
}
```
Response
```
{
    "covered_types": [
        "Grass",
        "Ice",
        "Bug",
        "Steel",
        "Psychic",
        "Ghost",
        "Grass",
        "Fairy",
        "Dragon",
        "Grass",
        "Fighting",
        "Bug",
        "Fire",
        "Ground",
        "Rock",
        "Grass",
        "Fighting",
        "Bug",
        "Fire",
        "Ice",
        "Flying",
        "Bug",
        "Psychic",
        "Ghost"
    ],
    "pokemons": [
        {
            "name": "Typhlosion",
            "base_status": {
                "hp": 78,
                "attack": 84,
                "defense": 78,
                "special_attack": 109,
                "special_defense": 85,
                "speed": 100,
                "total": 534
            },
            "types": [
                "Fire"
            ]
        },
        {
            "name": "Gengar",
            "base_status": {
                "hp": 60,
                "attack": 65,
                "defense": 60,
                "special_attack": 130,
                "special_defense": 75,
                "speed": 110,
                "total": 500
            },
            "types": [
                "Ghost",
                "Poison"
            ]
        },
        {
            "name": "Dragonite",
            "base_status": {
                "hp": 91,
                "attack": 134,
                "defense": 95,
                "special_attack": 100,
                "special_defense": 100,
                "speed": 80,
                "total": 600
            },
            "types": [
                "Dragon",
                "Flying"
            ]
        },
        {
            "name": "Gyarados",
            "base_status": {
                "hp": 95,
                "attack": 125,
                "defense": 79,
                "special_attack": 60,
                "special_defense": 100,
                "speed": 81,
                "total": 540
            },
            "types": [
                "Water",
                "Flying"
            ]
        },
        {
            "name": "Tyranitar",
            "base_status": {
                "hp": 100,
                "attack": 134,
                "defense": 110,
                "special_attack": 95,
                "special_defense": 100,
                "speed": 61,
                "total": 600
            },
            "types": [
                "Rock",
                "Dark"
            ]
        }
    ],
    "suggestion_types": [
        "Fighting",
        "Ground"
    ],
    "uncovered_types": [
        "Normal",
        "Electric",
        "Water",
        "Dark",
        "Poison"
    ]
}
```

## Types Suggestion ```/api/v1/type/suggestion```
Return types suggestion -> which type will be covered most other types by super effective attack<br>
Example:<br>
Payload
```
{
    "uncovered_types": [
        "Normal",
        "Electric",
        "Poison",
        "Dark",
        "Water"
    ],
    "suggestion_length": 15
}
```
Response
```
{
    "suggestion_types": [
        "Fighting",
        "Ground",
        "Psychic",
        "Grass",
        "Electric",
        "Bug",
        "Fairy"
    ]
}
```

## Note
The author will work on this repository on free time only :)

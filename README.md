# galactic-facts
An API to give random facts about space

## Routes
### /api/fact
Returns a JSON response of the form:
```
{
    "data": [
        "fact",
        "fact",
    ],
    "count": number of facts,
    "status": status code,
}
```
Use the parameter `count` to get specified number of facts.
## Instructions
- Copy `config.json.example` to `config.json` and edit as necessary 
- Run `go build` to build the application
- Run `./rediyeti-avalanche-adapter serve`

## Config options:
- `port`: what port to listen on
- `api_host`: location of avalanchego server
- `api_port`: port of avalanchego server

## Request
```json
{
    "id": string,
    "data": {
        "chain": string,
        "method": string,
        "params": Object,
    }
}
```
 
## Example 

```
curl -X POST -H "content-type:application/json" "http://localhost:8080/" --data '{ 
    "id": "abc", 
    "data": {
        "chain": "P", 
        "method": "platform.getCurrentValidators", 
        "params": {} 
    }
}'
```

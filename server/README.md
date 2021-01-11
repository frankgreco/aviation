```
$ grpcurl -d "{\"requested\": \"N_NUMBER=959\", \"existing\": [\"MAKE=BOEING\"]}" -plaintext 127.0.0.1:8082 types.SuggestionService.ListSuggestions | jq
$ curl -s "127.0.0.1:9092/v1/suggestions?requested=N_NUMBER%3D959&existing=MAKE%3DBOEING" | jq
{
  "type": "N_NUMBER",
  "suggestions": [
    "959AT",
    "959FD",
    "959AN",
    "959BP",
    "959WN",
    "959DN",
    "959NN"
  ],
  "size": 7
}
```

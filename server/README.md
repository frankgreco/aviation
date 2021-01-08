```
$ grpcurl -d "{\"requested\":{\"type\": \"N_NUMBER\", \"value\": \"959\"}, \"existing\": [{\"type\": \"MAKE\",  \"value\":  \"BOEING\"}]}" -plaintext 127.0.0.1:8082 types.SuggestionService.ListSuggestions
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
  ]
}
```
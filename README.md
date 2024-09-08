# vinyl_store

## Endpoints

### Create Album

```sh
 curl -H 'Content-Type: application/json' \
      -d '{"title":"Grace","artist":"Jeff Buckley","price":99.99}' \
      -X POST \
      https://xysxevsifj.execute-api.us-east-2.amazonaws.com/Prod/vinyl
```

### Get Album

```sh
 curl https://xysxevsifj.execute-api.us-east-2.amazonaws.com/Prod/vinyl/4a7f6d57-c324-4854-bf0a-f77926fa5e6c
```
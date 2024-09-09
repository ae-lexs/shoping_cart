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
 curl https://xysxevsifj.execute-api.us-east-2.amazonaws.com/Prod/vinyl/{id}
```

### Get Albums

```sh
 curl https://xysxevsifj.execute-api.us-east-2.amazonaws.com/Prod/vinyl
```
# Schlingel

## Development

```bash
# enter dev shell
nix develop

# apply code format
nix fmt

# running the application
nix run .
```

For API development, there is a [bruno](https://github.com/usebruno/bruno/) collection checked in which can be used.

## Building and running the container image

```bash
# build the container image
nix build .#containerimage

# load the container image into docker
docker load < result

# start the container
docker compose up -d
```

## Contributing

- Don't forget to apply code format: `nix fmt`
- Don't forget to extend or adjust the [bruno](https://github.com/usebruno/bruno/) collection if you did some changes to the API.

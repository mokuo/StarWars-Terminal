# Development

## Build

```zsh
cd cmd/starwars/
go build
./starwars
```

## Generate images

Install ImageMagick 7 (or 6).

```zsh
# ex) bb-8
composite -dissolve 40%x60% original_images/512x512_black.png original_images/starwars-bb-8.png cmd/starwars/images/bb-8.png
```

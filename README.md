# StarWars-Terminal

StarWars terminal themes inspired at [Pokemon-Terminal](https://github.com/LazoCoder/Pokemon-Terminal).

## Development

```
go run cmd/starwars/main.go

# or
go build -o starwars
./starwars
```

### Generate images

Install ImageMagick 7 (or 6).

```sh
# ex) bb-8
composite -dissolve 40%x60% original_images/512x512_black.png original_images/starwars-bb-8.png starwars/images/bb-8.png
```

# StarWars-Terminal

StarWars terminal themes inspired at [Pokemon-Terminal](https://github.com/LazoCoder/Pokemon-Terminal).

## Development

```
cd cmd/starwars/
go build
./starwars
```

### Generate images

Install ImageMagick 7 (or 6).

```sh
# ex) bb-8
composite -dissolve 40%x60% original_images/512x512_black.png original_images/starwars-bb-8.png cmd/starwars/images/bb-8.png
```

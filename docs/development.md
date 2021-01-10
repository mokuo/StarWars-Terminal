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

## Deploy

1. Install [lerna-changelog](https://github.com/lerna/lerna-changelog).
2. Get [personal access token](https://github.3.com/settings/tokens) for `GITHUB_TOKEN`.
3. Change `version` of version.go before deploy.
4. Deploy:

```zsh
cd devtools/deploy/
go build
GITHUB_TOKEN=xxx ./deploy

brew uninstall mokuo/starwars-terminal
brew untap mokuo/starwars-terminal
brew tap mokuo/starwars-terminal
brew install mokuo/starwars-terminal
```

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

## Deploy with GitHub Actions

Apply label to PR. (ref: [lerna-changelog](https://github.com/lerna/lerna-changelog))

Merge PR to `master` branch.

## Deploy from local machine

1. Install [lerna-changelog](https://github.com/lerna/lerna-changelog).
2. Get [personal access token](https://github.com/settings/tokens) ( `public_repo` scope) for `GITHUB_AUTH` .
    - go-git, go-github, lerna-changelog use `GITHUB_AUTH`
3. Change `version` of version.go before deploy.
4. Deploy:

```zsh
cd devtools/deploy
go build
GITHUB_AUTH=xxx ./deploy

# check
brew uninstall starwars-terminal
brew untap mokuo/starwars-terminal
brew tap mokuo/starwars-terminal
brew install starwars-terminal
```

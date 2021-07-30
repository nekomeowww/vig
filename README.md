<h1 align="center">vig</h1>
<p align="center">
  Vue inside Go, a complete boilerplate for Vue + Go project
</p>


---

## Features

1. vue3 as frontend framework
2. gin as backend framework
3. unit test supported
4. less, babel, eslint installed for vue3
5. vuex, vue-router, vue-i18n configured
6. fully functional build script for distribution

---

## Project preparation

### Download go mod
```
go mod download
go mod vendor
```

### Download frontend dependencies
```
yarn
```

### Init Git repo if not already
```
git init
```

### Use submodule for frontend
```shell
git rm -r --cached assets
cd assets
git init
git add *
git commit -S -sm "Init"
git branch -M main
git remote add origin git@[YOUR GIT SSH HOST]:[YOUR FRONTEND REPO NAME].git
cd ..
git submodule add git@[YOUR GIT SSH HOST]:[YOUR FRONTEND REPO NAME].git assets
git push -u origin main
```

### Push to your repo
```shell
git commit -S -sam "set frontend assets as submodule"
git push
```

### Build and distribute

### Compile by yourself

1. compile frontend asssets
```shell
cd assets
yarn build
cd ..
```

2. prepare frontend assets before compile backend
```shell
go get github.com/rakyll/statik
statik -src=assets/dist/  -include=*.html,*.js,*.json,*.css,*.png,*.svg,*.ico -f
```

3. compile backend as executable for your platform
```shell
export COMMIT=$(git rev-parse --short HEAD)
export VERSION=$(git describe --tags)
export STAGE=release

go build -a -o vig -ldflags " -X 'github.com/nekomeowww/vig/config.BackendVersion=$VERSION' -X 'github.com/nekomeowww/vig/config.LastCommit=$COMMIT' -X 'github.com/nekomeowww/vig/config.Stage=$STAGE'"
```

**NOTICE: Apple Silicon requires go 1.16 or higher to compile**

### Use script

1. build frontend assets only
```shell
./build.sh -a
```

2. build backend only
```shell
export STAGE=release
./build.sh -c
```

3. build frontend ans backend
```shell
export STAGE=release
./build.sh -b
```

4. build for all supported platform
```shell
export STAGE=release
./build.sh -r
```

**NOTICE: Please set your stage before running last command, or it will build a debug version of the package**


## Use Github Action as CI/CD
The workflow file is ready, you just need to adjust your project name in Line 182 in [release.yml](https://github.com/nekomeowww/vig/blob/main/.github/workflows/release.yml) where the **vig** presents.
```
    - name: Upload artifact
      uses: actions/upload-artifact@v1.0.0
      with:
        name: vig-${{ matrix.os }}
        path: release/
```

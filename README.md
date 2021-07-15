<h1 align="center">vig</h1>
<p align="center">
  Vue inside Go, a complete boilerplate for Vue + Go project
</p>


---

### Features

1. vue3 as frontend framework
2. gin as backend framework
3. unit test supported
4. less, babel, eslint installed for vue3
5. vuex, vue-router, vue-i18n configured
6. fully functional build script for distribution

---

### Project preparation

#### Download go mod
```
go mod download
go mod vendor
```

#### Download frontend dependencies
```
yarn
```

### Init Git repo if not already
```
git init
```

#### Use submodule for frontend
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

#### Push to your repo
```shell
git commit -S -sam "set frontend assets as submodule"
git push
```

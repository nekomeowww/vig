#!/bin/bash
echo "Checking..."
NAME=$(go list ./... | head -n 1 | rev | cut -d/ -f1 | rev)
REPO=$(cd $(dirname $0); pwd)
COMMIT_SHA=$(git rev-parse --short HEAD)
STAGE="$STAGE"
ASSETS="false"
BINARY="false"
RELEASE="false"
VERSION=$(git describe --tags)

if [ $? -ne 0 ]; then
  VERSION="1.0.0"
fi

if [[ $STAGE == "" ]]; then
  echo ""
  echo "You don't have any development stage set yet, default to 'debug'"
  echo "Please use 'export STAGE=release' to 'release' when build for production"
  echo ""
  STAGE="debug"
fi

buildAssets () {
  echo ""
  echo "Building assets..."
  echo ""

  cd $REPO
  rm -rf assets/dist
  rm -f statik/statik.go

  export CI=false

  cd $REPO/assets

  yarn install
  yarn run build

  GO_VERSION=$(go version)
  if [[ $GO_VERSION != *"1.16"* ]]; then
    echo "It's there!"
  fi

  if ! [ -x "$(command -v statik)" ]; then
    export CGO_ENABLED=0
    go get github.com/rakyll/statik
  fi

  cd $REPO
  statik -src=assets/dist/  -include=*.html,*.js,*.json,*.css,*.png,*.svg,*.ico,*.ttf -f
}

buildBinary () {
  echo ""
  echo "Building executable..."
  echo ""

  cd $REPO
  go build -a -o "release/${NAME}" -ldflags " -X 'github.com/nekomeowww/vig/config.BackendVersion=$VERSION' -X 'github.com/nekomeowww/vig/config.LastCommit=$COMMIT_SHA' -X 'github.com/nekomeowww/vig/config.Stage=$STAGE'"
}

vercomp () {
    if [[ $1 == $2 ]]
    then
        return 0
    fi
    local IFS=.
    local i ver1=($1) ver2=($2)
    # fill empty fields in ver1 with zeros
    for ((i=${#ver1[@]}; i<${#ver2[@]}; i++))
    do
        ver1[i]=0
    done
    for ((i=0; i<${#ver1[@]}; i++))
    do
        if [[ -z ${ver2[i]} ]]
        then
            # fill empty fields in ver2 with zeros
            ver2[i]=0
        fi
        if ((10#${ver1[i]} > 10#${ver2[i]}))
        then
            return 1
        fi
        if ((10#${ver1[i]} < 10#${ver2[i]}))
        then
            return 2
        fi
    done
    return 0
}

_build() {
    local osarch=$1
    IFS=/ read -r -a arr <<<"$osarch"
    os="${arr[0]}"
    arch="${arr[1]}"
    gcc="${arr[2]}"

    # Go build to build the binary.
    export GOOS=$os
    export GOARCH=$arch
    export CGO_ENABLED=0

    GOVERSION=`go version | { read _ _ v _; echo ${v#go}; }`
    SKIPARM64=false
    vercomp $GOVERSION "1.16"
    if [[ $? == 2 && $os == "darwin" && $arch == "arm64" ]]; then
      echo ""
      echo "darwin/arm64 requires ^1.16.0 of go, skipping..."
      echo ""
      return
    fi

    if [[ $os == "darwin" && $arch == "arm64" ]]; then
      export CC=""
      export CGO_ENABLED=1
      export SDKROOT=$(xcrun --sdk macosx --show-sdk-path)
    else
      export CC=$gcc
      export CGO_ENABLED=0
    fi

    echo ""
    echo "Building for $GOOS on $GOARCH using $CC with CGO_ENABLED=$CGO_ENABLED"
    echo ""

    if [ -n "$VERSION" ]; then
        out="release/${NAME}_${VERSION}_${os}_${arch}"
    else
        out="release/${NAME}_${COMMIT_SHA}_${os}_${arch}"
    fi

    go build -a -o "${out}" -ldflags " -X 'github.com/nekomeowww/vig/config.BackendVersion=$VERSION' -X 'github.com/nekomeowww/vig/config.LastCommit=$COMMIT_SHA' -X 'github.com/nekomeowww/vig/config.Stage=$STAGE'"

    if [ "$os" = "windows" ]; then
      mv $out release/${NAME}.exe
      zip -j -q "${out}.zip" "release/${NAME}.exe"
      rm -f "release/${NAME}.exe"
    else
      mv $out "release/${NAME}"
      tar -zcf "${out}.tar.gz" "release/${NAME}"
      rm -f "release/${NAME}"
    fi
}

release(){
  cd $REPO
  ## List of architectures and OS to test coss compilation.
  SUPPORTED_OSARCH="darwin/arm64/gcc"

  echo ""
  echo "Release builds for OS/Arch/CC: ${SUPPORTED_OSARCH}"
  echo "Building executable..."
  echo ""

  for each_osarch in ${SUPPORTED_OSARCH}; do
      _build "${each_osarch}"
  done
}

usage() {
  echo "Usage: $0 [-a] [-c] [-b] [-r]" 1>&2;
  exit 1;
}

while getopts "bacr" o; do
  case "${o}" in
    b)
      ASSETS="true"
      BINARY="true"
      ;;
    a)
      ASSETS="true"
      ;;
    c)
      BINARY="true"
      ;;
    r)
      ASSETS="true"
      RELEASE="true"
      ;;
    d)
      DEBUG="true"
      ;;
    *)
      usage
      ;;
  esac
done
shift $((OPTIND-1))

echo ""
echo "Building information..."
echo "Name:           $NAME"
echo "Repo:           $REPO"
echo "Build assets:   $ASSETS"
echo "Build binary:   $BINARY"
echo "Release:        $RELEASE"
echo "Version:        $VERSION"
echo "Stage:          $STAGE"
echo "Commit:         $COMMIT_SHA"
echo ""

if [ "$ASSETS" = "true" ]; then
  echo "building assets"
  buildAssets
fi

if [ "$BINARY" = "true" ]; then
  echo "building binary"
  buildBinary
fi

if [ "$RELEASE" = "true" ]; then
  echo "building release"
  release
fi

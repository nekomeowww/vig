#!/bin/bash
echo "Checking..."
NAME=$(go list ./... | head -n 1 | rev | cut -d/ -f1 | rev)
REPO=$(cd $(dirname $0); pwd)
COMMIT_SHA=$(git rev-parse --short HEAD)
VERSION=$(git describe --tags)
MODE="$STAGE"
ASSETS="false"
BINARY="false"
RELEASE="false"

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
  go build -a -o "release/${NAME}" -ldflags " -X 'github.com/nekomeowww/vig/config.BackendVersion=$VERSION' -X 'github.com/nekomeowww/vig/config.LastCommit=$COMMIT_SHA'"
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
    export CC=$gcc
    export CGO_ENABLED=1

    if [ -n "$VERSION" ]; then
        out="release/${NAME}_${VERSION}_${os}_${arch}"
    else
        out="release/${NAME}_${COMMIT_SHA}_${os}_${arch}"
    fi

    go build -a -o "${out}" -ldflags " -X 'github.com/nekomeowww/vig/config.BackendVersion=$VERSION' -X 'github.com/nekomeowww/vig/config.LastCommit=$COMMIT_SHA' -X 'github.com/nekomeowww/vig/config.LastCommit=$COMMIT_SHA'"

    if [ "$os" = "windows" ]; then
      mv $out release/${NAME}.exe
      zip -j -q "${out}.zip" "release/${NAME}.exe"
      rm -f "release/${NAME}.exe"
    else
      mv $out "release/${NAME}"
      tar -zcvf "${out}.tar.gz" -C "release ${NAME}"
      rm -f "release/${NAME}"
    fi
}

release(){
  cd $REPO
  ## List of architectures and OS to test coss compilation.
  SUPPORTED_OSARCH="linux/amd64/gcc linux/arm/arm-linux-gnueabihf-gcc windows/amd64/x86_64-w64-mingw32-gcc linux/arm64/aarch64-linux-gnu-gcc"

  echo ""
  echo "Release builds for OS/Arch/CC: ${SUPPORTED_OSARCH}"
  echo "Building executable..."
  echo ""

  for each_osarch in ${SUPPORTED_OSARCH}; do
      echo ""
      echo "Building for ${each_osarch}"
      echo ""
      _build "${each_osarch}"
  done
}

usage() {
  echo "Usage: $0 [-a] [-c] [-b] [-r]" 1>&2;
  exit 1;
}

while getopts "bacr:d" o; do
  echo ${o}
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
echo "MODE:           $MODE"
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

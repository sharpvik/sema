# Maintainer: Viktor A. Rozenko Voitenko <sharp.vik@gmail.com>
pkgname=sema
pkgver=1.1.1 # TODO: change in case of successful PR
pkgrel=14 # TODO: change in case of successful PR
pkgdesc="Semantic commit tool"
arch=(x86_64)
url="https://github.com/sharpvik/sema"
license=('MIT')
depends=()
makedepends=(git go)
provides=(sema)
conflicts=()
replaces=()
backup=()
options=()
source=("git+$url")
md5sums=('SKIP')

package() {
    mkdir -p $(go env GOPATH)/bin
    cd $pkgname
    go install
}

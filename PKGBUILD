# Maintainer: Thomas Kühnel <kuehnelth@gmail.com>

pkgname=secure-nyancatd
pkgver=r12.1010100
pkgrel=1
pkgdesc='Secure nyancat server'
url=https://github.com/kuehnelth/secure-nyancatd
arch=('i686' 'x86_64' 'armv7h' 'armv6h' 'aarch64')
license=('BSD')
depends=('glibc' 'nyancat')
makedepends=('git' 'go')
source=("$pkgname::git+http://github.com/kuehnelth/$pkgname")
md5sums=('SKIP')
sha1sums=('SKIP')
install=secure-nyancatd.install

pkgver() {
  cd "$srcdir/$pkgname"
  printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
  cd "$srcdir/$pkgname"
  export GOPATH="$srcdir"
  go get -u -v github.com/gliderlabs/ssh
  go get -u -v github.com/kr/pty

  go build -v
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm755 $pkgname "$pkgdir"/usr/bin/$pkgname
  install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
  install -Dm644 $pkgname.service "${pkgdir}"/usr/lib/systemd/system/$pkgname.service
}

# vim:set ts=2 sw=2 et:


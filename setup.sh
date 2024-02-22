#!/bin/bash
wget https://gmplib.org/download/gmp/gmp-6.3.0.tar.xz
if [[ $(sha256sum gmp-6.3.0.tar.xz) = \
	a3c2b80201b89e68616f4ad30bc66aee4927c3ce50e33929ca819d5c43538898* ]]
then
	tar -xvJf gmp-6.3.0.tar.xz
	pushd gmp-6.3.0
	sh configure --prefix=/usr/local/gmp --disable-static --enable-shared
	make
	make check
	sudo make install
	popd
else
	echo 'Cannot download GMP library from gmplib.org.'
	exit 1
fi

tail -f /dev/null
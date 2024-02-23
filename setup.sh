#!/bin/bash
if ! [[ -d "/usr/local/gmp" ]]; then
  wget https://github.com/zkweb3/GMP/releases/download/v6.3.0/gmp-6.3.0.tar.xz
  if [[ $(sha256sum gmp-6.3.0.tar.xz) = \
	  a3c2b80201b89e68616f4ad30bc66aee4927c3ce50e33929ca819d5c43538898* ]]
  then
	  tar -xvJf gmp-6.3.0.tar.xz
	  pushd gmp-6.3.0
	  sh configure --prefix=/usr/local/gmp
	  make
	  make check
	  make install
	  popd
  else
	  echo 'Cannot download GMP library from gmplib.org.'
	  exit 1
  fi
fi
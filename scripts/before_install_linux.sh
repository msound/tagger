#!/bin/bash
sudo apt-get update
sudo apt-get install -y build-essential cmake libpthread-stubs0-dev libssl-dev libssh-dev
BUILD_DIR=`pwd`
cd /tmp
rm -f v0.25.1.tar.gz
rm -rf libgit2-0.25.1/
wget https://github.com/libgit2/libgit2/archive/v0.25.1.tar.gz
tar -xzf v0.25.1.tar.gz
cd libgit2-0.25.1/
mkdir -p build && cd build
cmake -DTHREADSAFE=ON \
      -DBUILD_CLAR=OFF \
      -DBUILD_SHARED_LIBS=OFF \
      -DCMAKE_C_FLAGS=-fPIC \
      -DCMAKE_BUILD_TYPE="RelWithDebInfo" \
      ..
sudo cmake --build . --target install
cd $BUILD_DIR

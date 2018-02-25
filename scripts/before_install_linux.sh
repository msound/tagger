#!/bin/bash
# sudo apt-get update
# sudo apt-get install -y build-essential cmake
# BUILD_DIR=`pwd`
# cd /tmp
# rm -f v0.25.1.tar.gz
# rm -rf libgit2-0.25.1/
# wget https://github.com/libgit2/libgit2/archive/v0.25.1.tar.gz
# tar -xzf v0.25.1.tar.gz
# cd libgit2-0.25.1/
# mkdir build && cd build
# cmake ..
# sudo cmake --build . --target install
# cd $BUILD_DIR
cd vendor/gopkg.in/libgit2/git2go.v25/
git submodule update --init
make install-static
cd ../../../../

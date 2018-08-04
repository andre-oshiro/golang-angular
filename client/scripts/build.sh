#!/bin/sh -eux

DIST_DIR="./dist"
SPA_DIR="../server/src/SPA"
TARGET_FILES="./dist/client/*"

if [ -e $DIST_DIR ]; then
    rm -rf $DIST_DIR
     echo "Removed: $DIST_DIR"
fi

echo 'Do Ahead of Time compilation.'
ng build --aot --prod
[ $? -ne 0 ] && exit 1


if [ -e $SPA_DIR ]; then
    rm -rf $SPA_DIR
     echo "Removed: $SPA_DIR"
fi

mkdir -p $SPA_DIR
cp -r $TARGET_FILES $SPA_DIR
rm -rf $DIST_DIR
[ $? -ne 0 ] && exit 1

exit 0



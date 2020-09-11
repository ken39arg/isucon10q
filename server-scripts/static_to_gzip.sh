#!/usr/bin/env sh

for type in ico txt ttf woff eot svg jpg png gif js css html json appcach
do
    find . -type f -name "*.${type}" | xargs gzip -k
done

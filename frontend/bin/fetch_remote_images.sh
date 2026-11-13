#!/bin/bash -x

# STEPS
# 1. Get each markdown file
# 2. Find images that match https://blog.mbell.dev/...
# 3. download to `src/assets` directory
# 4. replace every occurence with local path

_files=$(fd -tf --extension=md)
_capture='https:\/\/blog\.mbell\.dev\/(\w|\d|_|\/)+\.\w+'
_static_dir='src/assets/img'

echo "$_files"

for file in $_files; do
  _images=$(grep -oE "${_capture}" "${file}" | uniq)

  for img in $_images; do
    filename="${img##*/}"
    curl -o "${_static_dir}/${filename}" "${img}"
    echo "$_files" | xargs -I % -n1 sed -i '' "s,${img},/src/assets/img/${filename}," '%'
  done
done

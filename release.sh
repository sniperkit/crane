#!/usr/bin/env bash

set -eux

version=$1

if [ -z "$version" ]; then
  echo "No version passed! Example usage: ./release.sh 1.0.0"
  exit 1
fi

echo "Running tests..."
crane cmd test

echo "Update version..."
old_version=$(grep -o "[0-9]*\.[0-9]*\.[0-9]*" pkg/core/version_basic.go)
sed -i.bak 's/Version = "'$old_version'"/Version = "'$version'"/' pkg/core/version_basic.go pkg/core/version_pro.go
rm pkg/core/version_basic.go.bak
rm pkg/core/version_pro.go.bak
sed -i.bak 's/VERSION="'$old_version'"/VERSION="'$version'"/' download.sh
rm download.sh.bak
sed -i.bak 's/'$old_version'/'$version'/' README.md
rm README.md.bak

echo "Mark version as released in changelog..."
today=$(date +'%Y-%m-%d')
sed -i.bak 's/Unreleased/Unreleased\
\
## '$version' ('$today')/' docs/CHANGELOG.md
rm docs/CHANGELOG.md.bak

echo "Update contributors..."
git contributors | awk '{for (i=2; i<NF; i++) printf $i " "; print $NF}' > docs/CONTRIBUTORS

echo "Build binaries..."
crane cmd build

echo "Update repository..."
git add pkg/core/version_basic.go download.sh README.md docs/CHANGELOG.md docs/CONTRIBUTORS
git commit -m "Bump version to ${version}"
git tag --sign --message="v$version" --force "v$version"
git tag --sign --message="latest" --force latest


echo "v$version tagged."
echo "Now, run 'git push origin master && git push --tags --force' and publish the release on GitHub."

#!/bin/bash

# Leer la última versión del archivo VERSION
previousVersion=$(cat VERSION 2>/dev/null)

# Incrementar la parte PATCH
IFS='.' read -r major minor patch <<< "$previousVersion"
patch=$((patch + 1))

# Crear la nueva versión
newVersion="$major.$minor.$patch"

# Escribir la nueva versión en el archivo VERSION
echo "$newVersion" > VERSION

# Commit y tag
git add VERSION
git commit -m "Increment version to $newVersion"
git tag "v$newVersion"
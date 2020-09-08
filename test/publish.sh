#!/usr/bin/env bash

# Variables
CHARTS_DIR="$1"
HELM_URL="$2"
HELM_PKG="docs"

# Build each chart separately
for i in $(ls -1 ${CHARTS_DIR})
do
    echo "building ${i} chart"
    helmv3 package -d ${HELM_PKG} ${CHARTS_DIR}/${i}
done

# Update index.yaml with new charts
helmv3 repo index ${HELM_PKG} --url ${HELM_URL}

# Commit and push packages
git config --global user.email "46964379+petro-matic@users.noreply.github.com"
git config --global user.name "Petro Matviichuk"
git add docs
git commit -m "publish ${GITHUB_SHA}"
git push origin master

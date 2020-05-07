#!/usr/bin/env bash
git push origin :0.0.1
git tag -d 0.0.1
git tag 0.0.1
git push origin 0.0.1

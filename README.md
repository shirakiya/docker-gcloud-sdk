# docker-gcloud-sdk

[![Check Version](https://github.com/shirakiya/docker-gcloud-sdk/actions/workflows/check_version.yml/badge.svg?branch=main)](https://github.com/shirakiya/docker-gcloud-sdk/actions/workflows/check_version.yml)
[![Publish image](https://github.com/shirakiya/docker-gcloud-sdk/actions/workflows/publish.yml/badge.svg)](https://github.com/shirakiya/docker-gcloud-sdk/actions/workflows/publish.yml)

This project builds and publishes the docker images that contains [gcloud-sdk](https://cloud.google.com/cli)
and [kubectl](https://kubernetes.io/ja/docs/reference/kubectl/overview/). The Docker Hub repository page is below.

https://hub.docker.com/r/shirakiya/gcloud-sdk

## Build workflow

The workflow to publish image runs in GitHub Actions with following steps.

1. Check gcloud-sdk version at **00:00(UTC) every day**.
2. (If new version found, ) create a pull-request.
3. When the pull-request is merged, it is automatically built and pushed to the Docker Hub repository.

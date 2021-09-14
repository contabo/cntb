#!/bin/bash

USER_ID="${OAUTH2_USER_ID:-3619aec6-2d61-4eb8-980e-b3178f8ba6ef}"
TEST_SUFFIX="-59A80D5D-1027-4D9F-94E4-1C87E8EEF8CE-integration"
IMAGE_DOWNLOAD_URL="https://dl-cdn.alpinelinux.org/alpine/v3.13/releases/s390x/alpine-standard-3.13.5-s390x.iso"
INSTANCE_ID=100

load_lib() {
  local name="$1"
  if [ -z "${BATS_INSTALLATION+x}" ]; then
    load "/usr/local/lib/${name}/load.bash";
  else
    load "${BATS_INSTALLATION}/${name}/load.bash";
  fi
}

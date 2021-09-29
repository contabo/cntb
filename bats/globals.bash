#!/bin/bash

USER_ID="${OAUTH2_USER_ID:-adcb2f8d-5750-4525-9be6-dd0f397937fa}"
TEST_SUFFIX="-59A80D5D-1027-4D9F-94E4-1C87E8EEF8CE-integration"
IMAGE_DOWNLOAD_URL="https://dl-cdn.alpinelinux.org/alpine/v3.13/releases/s390x/alpine-standard-3.13.5-s390x.iso"
INSTANCE_ID=100
REINSTALL_INSTANCE_ID=105
STANDARD_IMAGE_ID="69b52ee3-2fda-4f44-b8de-69e480d87c7d"
STANDARD_IMAGE_ID2="db1409d2-ed92-4f2f-978e-7b2fa4a1ec90"
INT_ENVIRONMENT="${INT_ENVIRONMENT:-dev}"

load_lib() {
  local name="$1"
  if [ -z "${BATS_INSTALLATION+x}" ]; then
    load "/usr/local/lib/${name}/load.bash";
  else
    load "${BATS_INSTALLATION}/${name}/load.bash";
  fi
}

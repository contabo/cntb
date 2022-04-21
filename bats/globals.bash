#!/bin/bash

CURRENTEPOCTIME=`date +%s`
CURRENTEPOCTIME_NANAO=`date +%s%N`
USER_ID="${OAUTH2_USER_ID:-3e97f2b0-eccd-497b-8516-79a6708a9cf4}"
TEST_SUFFIX="$CURRENTEPOCTIME-integration-cli"
TEST_BUCKET_NAME="test123-integration-cli"
IMAGE_DOWNLOAD_URL='https://dl-cdn.alpinelinux.org/alpine/v3.13/releases/s390x/alpine-standard-3.13.5-s390x.iso'
INSTANCE_ID=100
REINSTALL_INSTANCE_ID=105
STANDARD_IMAGE_ID="db1409d2-ed92-4f2f-978e-7b2fa4a1ec90"
INT_ENVIRONMENT="${INT_ENVIRONMENT:-dev}"

load_lib() {
  local name="$1"
  if [ -z "${BATS_INSTALLATION+x}" ]; then
    load "/usr/local/lib/${name}/load.bash";
  else
    load "${BATS_INSTALLATION}/${name}/load.bash";
  fi
}

poll_instance(){
  echo "polling entered" >&3
  for i in {0..50}; do
    local status=$(./cntb get instance "$1" -o json | jq -r '.[].status')

    if [ "$status" == 'running' ]; then
      return 0
    elif [ "$status" == 'stopped' ]; then
      return 0
    fi

    sleep 2
  done

  return 1
}

deleteSnapshotsIfExisting() {
  local instance_id="$1"
  existingSnapshot=$(./cntb get snapshots $instance_id -s 1 -o json | jq -r '.[].snapshotId')

  if [ -n "$existingSnapshot" ]; then
    ./cntb delete snapshot "$1" "$existingSnapshot"
  
    sleep 4

    deleteSnapshotsIfExisting "$1"
  else
    sleep 4
  fi
}

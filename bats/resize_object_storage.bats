#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load cleanup-object-storage.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
    ensureTestConfig
}

function teardown_file() {
    restore_config_files

    deleteObjectStorageIfExisting "EU"
}

@test 'resize object storage: ok' {
  deleteObjectStorageIfExisting "EU"
  
  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success
  objectStorageId="$output"

  sleep 10 # waiting for the subscription to be published and subscribed

  run ./cntb resize objectStorage "${objectStorageId}" --totalPurchasedSpaceTB 2
  assert_success

  run ./cntb resize objectStorage "${objectStorageId}" --scalingLimitTB 2
  assert_success

  run ./cntb resize objectStorage "${objectStorageId}" --totalPurchasedSpaceTB 3 --scalingLimitTB 3
  assert_success
}

@test 'resize object storage - only auto scaling: ok' {
  objectStorageId=$(getObjectStorageInRegion 'EU')

  run ./cntb resize objectStorage "${objectStorageId}" --scalingLimitTB 2
  assert_success

  run ./cntb resize objectStorage "${objectStorageId}" --scalingState "disabled"
  assert_success
}

@test 'resize object storage with invalid auto scaling: nok' {
  objectStorageId=$(getObjectStorageInRegion 'EU')
  
  run ./cntb resize objectStorage "${objectStorageId}" --totalPurchasedSpaceTB 1 --scalingLimitTB 1 --scalingState "enabled"
  assert_failure
  assert_output --partial 'Error while updating object storage: 400 - Bad Request totalPurchasedSpaceTB size limit should be bigger than the existing one.\n'
}

@test 'resize object storage: nok : non existing object storage' {
  run ./cntb resize objectStorage "23729f17-85d3-4f0c-b4b8-b8b90225d167" --totalPurchasedSpaceTB 40
  assert_failure
  assert_output --partial 'Error while updating object storage: 404 - Entry ObjectStorage not found by objectStorageId'
}

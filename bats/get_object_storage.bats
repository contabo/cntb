#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load cleanup-object-storage.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
    ensureTestConfig
    deleteCache
}

function teardown_file() {
    restore_config_files
}

@test 'get object storage: normal ok' {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success
  objectStorageId="$output"

  run ./cntb get objectStorage "${objectStorageId}"
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
}

@test 'get object storage: wide ok' {
  objectStorageId=$(getObjectStorageInRegion "EU")

  run ./cntb get objectStorage "${objectStorageId}" -o wide
  assert_success
  assert_output --partial 'OBJECTSTORAGEID'
  assert_output --partial 'DATACENTER'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial 'STATUS'
  assert_output --partial 'TOTALPURCHASEDSPACETB'
  assert_output --partial 'S3URL'
}

@test 'get object storage: nok :objectStorageId not UUID' {
  run ./cntb get objectStorage 100
  assert_failure
  assert_output --partial 'Error while retrieving object storage: 400 - [objectStorageId must be a UUID]'
}

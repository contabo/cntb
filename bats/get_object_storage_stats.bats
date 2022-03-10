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

@test 'get object storage stats : ok' {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success
  objectStorageId="$output"

  run ./cntb stats objectStorage "${objectStorageId}"
  assert_success
  assert_output --partial 'USEDSPACETB'
  assert_output --partial 'USEDSPACEPERCENTAGE'
  assert_output --partial 'NUMBEROFOBJECTS'

  run ./cntb stats objectStorage "${objectStorageId}" -o json
  assert_success
  assert_output --partial '"customerId"'
  assert_output --partial '"numberOfObjects"'
  assert_output --partial '"tenantId"'
  assert_output --partial '"usedSpacePercentage"'
  assert_output --partial '"usedSpaceTB"'
}

@test 'get object storage stats no objectStorageId : nok' {
  run ./cntb stats objectStorage
  assert_failure
}

@test 'get object storage stats wrong objectStorageId : nok' {
  run ./cntb stats objectStorage 3162351f-456b-423a-b127-ab77baaa4e48
  assert_failure
}
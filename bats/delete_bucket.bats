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

@test 'delete bucket : ok' {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success

  run ./cntb create bucket EU ${TEST_SUFFIX}
  assert_success

  run ./cntb delete EU bucket EU ${TEST_SUFFIX}
  assert_success
}

@test 'delete bucket : nok : not existing' {
  run ./cntb delete bucket EU a19b8645-79d2-4fbb-ac56-a927d69b8d2b
  assert_failure
  assert_output --partial 'The specified bucket does not exist.'

  deleteObjectStorageIfExisting "EU"
}

@test 'delete bucket : nok : missing argument' {
  run ./cntb delete bucket EU
  assert_failure
  assert_output --partial 'Please provide a region and a bucketName'

  run ./cntb delete bucket bucket123
  assert_failure
  assert_output --partial 'Please provide a region and a bucketName'
}

@test 'delete buckets : nok : Invalid region' {
  run ./cntb delete bucket USA "${TEST_SUFFIX}"
  assert_failure
  assert_output --partial 'No Object Storage could be found in this region.'
}

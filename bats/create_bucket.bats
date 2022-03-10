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

@test 'create buckets : ok : command line flags' {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success

  run ./cntb create bucket EU ${TEST_SUFFIX}
  assert_success

  run ./cntb delete bucket EU ${TEST_SUFFIX}
  assert_success
}

@test 'create buckets : ok' {
  run ./cntb create bucket EU ${TEST_SUFFIX}
  assert_success

  run ./cntb delete bucket EU ${TEST_SUFFIX}
  assert_success
}

@test 'create buckets : nok : Invalid name' {
  run ./cntb create bucket EU 'T..${TEST_SUFFIX}'
  assert_failure
  assert_output --partial 'Could not create bucket T..${TEST_SUFFIX}. Name may contain unaccepted characters.'
}

@test 'create buckets : nok : Invalid region' {
  run ./cntb create bucket DE 'T..${TEST_SUFFIX}'
  assert_failure
  assert_output --partial 'No Object Storage could be found in this region.'

  deleteObjectStorageIfExisting "EU"
}


@test 'create buckets : nok : missing argument' {
  run ./cntb create bucket  'T..${TEST_SUFFIX}'
  assert_failure
  assert_output --partial 'Please provide a region and a bucketName'

  run ./cntb create bucket  EU
  assert_failure
  assert_output --partial 'Please provide a region and a bucketName'
}


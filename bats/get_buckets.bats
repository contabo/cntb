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

@test 'get buckets : ok' {
  deleteObjectStorageIfExisting "EU"

  run ./cntb create objectStorage --region "EU" --totalPurchasedSpaceTB 1 --scalingState "enabled" --scalingLimitTB 1
  assert_success

  run ./cntb create bucket EU ${TEST_SUFFIX}
  assert_success

  run ./cntb get buckets
  assert_success
  assert_output --partial 'NAME'

  run ./cntb delete EU ${TEST_SUFFIX}
  assert_success
}

@test 'get buckets wide: ok' {

 run ./cntb create bucket EU ${TEST_SUFFIX}
 assert_success

  run ./cntb get buckets -o wide
  assert_success

  assert_output --partial 'NAME'
  assert_output --partial 'CREATIONDATE'

  run ./cntb delete bucket EU ${TEST_SUFFIX}
  assert_success
}

@test 'get buckets : ok : format json' {

  run ./cntb create bucket EU ${TEST_SUFFIX}
  assert_success

  run ./cntb get buckets -o json
  assert_success

  assert_output --partial '"creationDate"'
  assert_output --partial '"name"'

  run ./cntb delete bucket EU ${TEST_SUFFIX}
  assert_success
}

@test 'get buckets : ok : format yaml' {
  run ./cntb create bucket EU ${TEST_SUFFIX}
  assert_success

  run ./cntb get buckets -o json
  assert_success

  assert_output --partial '"creationDate":'
  assert_output --partial '"name":'

  run ./cntb delete bucket EU ${TEST_SUFFIX}
  assert_success
}

@test 'get buckets : ok : specify region' {

  run ./cntb create bucket EU ${TEST_SUFFIX}
  assert_success

  run ./cntb get buckets -r EU
  assert_success

  assert_output --partial 'NAME'

  run ./cntb delete  bucket EU ${TEST_SUFFIX}
  assert_success

  deleteObjectStorageIfExisting "EU"
}


@test 'get buckets : nok : specify invalid region' {
  run ./cntb get buckets -r EUO
  assert_failure
  assert_output --partial 'No Object Storage could be found in this region'
}

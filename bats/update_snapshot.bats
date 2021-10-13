#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
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

@test 'update snapshot: ok' {
  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_success
  snapshotId="$output"

  run ./cntb update snapshot ${INSTANCE_ID} "$snapshotId" --name="snapshot${TEST_SUFFIX}" --description='test snapshot1'
  assert_success

  run ./cntb get snapshot ${INSTANCE_ID} "$snapshotId"
  assert_success
  assert_output --partial 'SNAPSHOTID'
  assert_output --partial 'NAME'
  assert_output --partial 'DESCRIPTION'
  assert_output --partial 'CREATEDDATE'
  assert_output --partial "snapshot${TEST_SUFFIX}"
  assert_output --partial "test snapshot1"
  assert_output --partial "$snapshotId"

  # clean up
  run ./cntb delete snapshot ${INSTANCE_ID} "$snapshotId"
  assert_success
}

@test 'update snapshot without snapshotId: nok' {
  run ./cntb update snapshot ${INSTANCE_ID}
  assert_failure
}

@test 'update snapshot with invalid snapshotId: nok' {
  run ./cntb update snapshot ${INSTANCE_ID} abc
  assert_failure
}

@test 'update snapshot with wrong arguments: nok' {
  run ./cntb update snapshot ${INSTANCE_ID} 72f0a447-8040-4098-9cb1-2fa3093a18f7 --nameSnap='test' --desc='test snapshot'
  assert_failure
}

@test 'update snapshot with invalid instanceId: nok' {
  run ./cntb update snapshot abc --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_failure
}

@test 'update snapshot without instanceId: nok' {
  run ./cntb update snapshot --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_failure
}
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

@test 'rollback snapshot: ok' {
  poll_instance "${INSTANCE_ID}"

  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot1'
  assert_success
  snapshotId="$output"

  run ./cntb rollback snapshot ${INSTANCE_ID} "$snapshotId"
  assert_success

  # clean up
  run ./cntb delete snapshot ${INSTANCE_ID} "$snapshotId"
  assert_success
}

@test 'rollback older snapshpot: nok' {
  poll_instance "${INSTANCE_ID}"

  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot1'
  assert_success
  snapshotId1="$output"

  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot2'
  assert_success
  snapshotId2="$output"

  run ./cntb rollback snapshot ${INSTANCE_ID} "$snapshotId1"
  assert_failure
  assert_output --partial '422'
  assert_output --partial "Snapshot cannot be used to rollback instance as is not the latest version"

  # clean up
  run ./cntb delete snapshot ${INSTANCE_ID} "$snapshotId1"
  assert_success

  run ./cntb delete snapshot ${INSTANCE_ID} "$snapshotId2"
  assert_success
}

@test 'rollback snapshot without arguments: nok' {
  run ./cntb rollback snapshot
  assert_failure
}

@test 'rollback snapshot with invalid instanceId: nok' {
  run ./cntb rollback snapshot abc
  assert_failure
}

@test 'rollback snapshot without instanceId: nok' {
  run ./cntb rollback snapshot
  assert_failure
}

@test 'rollback snapshot with invalid snapshotId: nok' {
  run ./cntb rollback snapshot ${INSTANCE_ID} abc
  assert_failure
}

@test 'rollback snapshot without snapshotId: nok' {
  run ./cntb rollback snapshot ${INSTANCE_ID}
  assert_failure
}
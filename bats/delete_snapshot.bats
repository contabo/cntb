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

@test 'delete snapshot: ok' {
  deleteSnapshotsIfExisting ${INSTANCE_ID}

  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_success
  snapshotId="$output"

  run ./cntb delete snapshot ${INSTANCE_ID} "$snapshotId"
  assert_success

  run ./cntb get snapshot ${INSTANCE_ID} "$snapshotId"
  assert_failure
}

@test "delete not existing snapshot: nok" {
  run ./cntb delete snapshot ${INSTANCE_ID} 72f0a447-8040-4098-9cb1-2fa3093a18f7
  assert_failure
  assert_output --partial 'Error while deleting snapshot: 404 - Entry Snapshot not found by snapshotId'
}


@test 'delete snapshot without arguments: nok' {
  run ./cntb delete snapshot
  assert_failure
}

@test 'delete snapshot with invalid instanceId: nok' {
  run ./cntb delete snapshot abc
  assert_failure
}

@test 'delete snapshot without instanceId: nok' {
  run ./cntb delete snapshot
  assert_failure
}

@test 'delete snapshot with invalid snapshotId: nok' {
  run ./cntb delete snapshot ${INSTANCE_ID} abc
  assert_failure
}

@test 'delete snapshot without snapshotId: nok' {
  run ./cntb delete snapshot ${INSTANCE_ID}
  assert_failure
}
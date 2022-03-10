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

@test 'get snapshots history: ok' {
  deleteSnapshotsIfExisting ${INSTANCE_ID}

  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_success
  snapshotId="$output"

  run ./cntb history snapshots
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'SNAPSHOTID'
  assert_output --partial 'ACTION'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial "$snapshotId"
}

@test 'get snapshots history wide: ok' {
  deleteSnapshotsIfExisting ${INSTANCE_ID}

  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_success
  snapshotId="$output"

  run ./cntb history snapshots -o wide
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'INSTANCEID'
  assert_output --partial 'SNAPSHOTID'
  assert_output --partial 'ACTION'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'CHANGES'
  assert_output --partial "$snapshotId"
}

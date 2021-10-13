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

@test 'create snapshot: ok' {
  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_success
  snapshotId="$output"
  # clean up
  run ./cntb delete snapshot ${INSTANCE_ID} "$snapshotId"
  assert_success
}

@test 'create snapshot without description argument: ok' {
  run ./cntb create snapshot ${INSTANCE_ID} --name="snapshot${TEST_SUFFIX}"
  assert_success
  snapshotId="$output"
  # clean up
  run ./cntb delete snapshot ${INSTANCE_ID} "$snapshotId"
  assert_success
}

@test 'create snapshot without arguments: nok' {
  run ./cntb create snapshot ${INSTANCE_ID}
  assert_failure
}

@test 'create snapshot without name argument: nok' {
  run ./cntb create snapshot ${INSTANCE_ID} --description='test snapshot'
  assert_failure
}

@test 'create snapshot with invalid instanceId: nok' {
  run ./cntb create snapshot abc --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_failure
}

@test 'create snapshot without instanceId: nok' {
  run ./cntb create snapshot --name="snapshot${TEST_SUFFIX}" --description='test snapshot'
  assert_failure
}
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

@test 'get secrets history: ok' {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --value="Test12345!" --type="password"
  assert_success
  secretId="$output"

  run ./cntb history secrets
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'SECRETID'
  assert_output --partial 'ACTION'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial "$secretId"

  run ./cntb delete secret $secretId
  assert_success
}

@test 'get secrets history wide: ok' {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --value="Test12345!" --type="password" -o wide
  assert_success
  secretId="$output"

  run ./cntb history secrets -o wide
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'SECRETID'
  assert_output --partial 'ACTION'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'CHANGES'
  assert_output --partial "$secretId"

  run ./cntb delete secret $secretId
  assert_success
}

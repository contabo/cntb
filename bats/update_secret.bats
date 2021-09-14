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

@test 'update secret: ok' {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --value="Test12345!" --type="password"
  assert_success
  secretId=$(echo "$output" | sed -n 's/.*secretId\s\+\([0-9a-zA-Z-]\+\).*$/\1/p')

  run ./cntb update secret $secretId --name="secretUpdate${TEST_SUFFIX}"
  assert_success

  run ./cntb get secret $secretId
  assert_success
  assert_output --partial 'SECRETID'
  assert_output --partial 'NAME'
  assert_output --partial 'TYPE'
  assert_output --partial "secretUpdate${TEST_SUFFIX}"
  assert_output --partial "$secretId"

  run ./cntb delete secret $secretId
  assert_success
}

@test 'update secret without secretId: nok' {
  run ./cntb update secret
  assert_failure
}

@test 'update secret with invalid secretId: nok' {
  run ./cntb update secret abc
  assert_failure
}

@test 'update secret with wrong arguments: nok' {
  run ./cntb update secret abc 123 --name1="secret${TEST_SUFFIX}" --value1="Test12345!"
  assert_failure
}
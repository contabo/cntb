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

@test 'delete secret: ok' {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --value="Test12345!" --type="password"
  assert_success
  secretId=$(echo "$output" | sed -n 's/.*secretId\s\+\([0-9a-zA-Z-]\+\).*$/\1/p')

  run ./cntb delete secret $secretId
  assert_success

  run ./cntb get secret $secretId
  assert_failure
}

@test "delete not existing secret: ok" {
  run ./cntb delete secret 123
  assert_failure
  assert_output --partial 'Error'
  assert_output --partial '404,'
  assert_output --partial "Error while deleting secret:"
}


@test 'delete secret without arguments: nok' {
  run ./cntb delete secret
  assert_failure
}

@test 'delete secret with invalid secretId: nok' {
  run ./cntb delete secret abc
  assert_failure
}

@test 'delete secret without secretId: nok' {
  run ./cntb delete secret
  assert_failure
}
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

@test 'create secret: ok' {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --value="Test12345!" --type="password"
  assert_success
  secretId="$output"

  run ./cntb delete secret $secretId
  assert_success
}

@test 'create secret without arguments: nok' {
  run ./cntb create secret
  assert_failure
}

@test 'create secret without name argument: nok' {
  run ./cntb create secret --value="Test12345!" --type="password"
  assert_failure
}

@test 'create secret without value argument: nok' {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --type="password"
  assert_failure
}

@test 'create secret without type argument: nok' {
  run ./cntb create secret --name="secret${TEST_SUFFIX}" --value="Test12345!"
  assert_failure
}

@test 'create secret with wrong type argument: nok' {
  run ./cntb create secret  --name="secret${TEST_SUFFIX}" --value="Test12345!" --type="other"
  assert_failure
}
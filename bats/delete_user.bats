#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
    ensureTestConfig
}

function teardown_file() {
    restore_config_files
}

@test "delete user ok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
  assert_success
  roleId="$output"

  run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --roles="$roleId" --locale de
  assert_success
  userId="$output"

  run ./cntb delete user "${userId}"
  assert_success

  #cleanup
  run ./cntb delete role "${roleId}"
  assert_success

}

@test "delete user wrong id type: nok" {
    run ./cntb delete user abc
    assert_failure
}

@test "delete user not existing: nok" {
  run ./cntb delete user 0
  assert_failure
  assert_output --partial "Error while deleting user: 404 - Entry User not found by userId"
}

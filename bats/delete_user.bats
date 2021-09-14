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
  run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
  assert_success
  roleId="$output"

  run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
  assert_success
  userId="$output"

  run ./cntb delete user "${userId}"
  assert_success

  #cleanup
  run ./cntb delete role apiPermission "${roleId}"
  assert_success

}

@test "delete user wrong id type: nok" {
    run ./cntb delete user abc
    assert_failure
}

@test "delete user not existing: nok" {
  run ./cntb delete user 0
  assert_failure
  assert_output --partial "404 Not Found"
}

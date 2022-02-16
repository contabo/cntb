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

@test "delete role : ok" {
    run ./cntb create role --name "foo${TEST_SUFFIX}" --permissions '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb get role "$roleId"
    assert_success

    run ./cntb delete role "$roleId"
    assert_success

    run ./cntb get role "$roleId"
    assert_failure
}

@test "delete not existing role : nok" {
  run ./cntb delete role 0
  assert_failure
  assert_output --partial "Error while deleting role: 404 - Entry Role not found by roleId"
}

@test "delete without roleId : nok" {
    run ./cntb delete role
    assert_failure
}

@test "delete with role type and roleId : nok" {
  run ./cntb delete role apiPermission 2
  assert_failure
}
@test "delete with roleId and role type : nok" {
  run ./cntb delete role 2 apiPermission
  assert_failure
}

@test "user string instead of number for role id : nok" {
    run ./cntb delete role twelve
    assert_failure
}

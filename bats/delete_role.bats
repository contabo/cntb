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

@test "delete role ok:" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb delete role apiPermission "$roleId"
    assert_success


    run ./cntb get role apiPermission "$roleId"
    assert_failure
}

@test "delete not existing role" {
  run ./cntb delete role apiPermission 0
  assert_failure
  assert_output --partial "Error while deleting role: 404 - Entry Role not found by roleId"
}

@test "use wrong permission type: nok" {
  run ./cntb delete role akdslfj 0
  assert_failure
}

@test "use only 1 input: nok" {
    run ./cntb delete role apiPermission
    assert_failure
}

@test "user more than 2 inputs: nok" {
  run ./cntb delete role apiPermission 2 apiPermission
  assert_failure
}

@test "user string instead of number for role id" {
    run ./cntb delete role apiPermission test
    assert_failure
}

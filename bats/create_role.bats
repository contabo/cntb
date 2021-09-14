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


@test "create role api permission: ok" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    # clean up
    run ./cntb delete role apiPermission "${roleId}"
    assert_success
}


@test "create role wrong permission type: nok" {
    run ./cntb create role aa --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_failure
}

@test "create role wrong number of inputs: nok" {
    run ./cntb create role --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_failure

    run .run ./cntb create role apiPermission apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_failure
}


wrapperCreateApiPermissionFile() {
    echo "$1" | ./cntb create role apiPermission -f -
}

@test "create from file api permission: ok" {
    run wrapperCreateApiPermissionFile '{"name": "foo", "apiPermissions": [{"apiName": "/v1/users", "actions": ["READ", "CREATE"]}]}'
    assert_success
    roleId="$output"

    # clean up
    run ./cntb delete role apiPermission "$roleId"
    assert_success
}


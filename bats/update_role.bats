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


@test "update role: ok : set different action to api" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ"]}]'
    assert_success
    roleId="$output"

    run ./cntb update role apiPermission "$roleId" --name="foo${TEST_SUFFIX}" -a='[{"apiName": "/v1/users", "actions": ["READ", "CREATE"] }]'
    assert_success

    run ./cntb get role apiPermission "$roleId" -o json
    assert_success
    assert_output --partial "CREATE"

    # cleanup

    run ./cntb delete role apiPermission "$roleId"
    assert_success

}

@test "update role not available action: nok" {
        run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ"]}]'
    assert_success
    roleId="$output"

    run ./cntb update role apiPermission "$roleId" --name="foo${TEST_SUFFIX}" -a='[{"apiName": "/v1/users", "actions": ["READ", "UPDATE"] }]'
    assert_failure

    # cleanup

    run ./cntb delete role apiPermission "$roleId"
    assert_success
}

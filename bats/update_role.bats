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


@test "update role : ok" {
    name="foo${TEST_SUFFIX}"
    run ./cntb create role -n "$name" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb get role "$roleId"
    assert_success
    assert_output --partial "$name"

    run ./cntb update role "$roleId" --name="updated$name" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success

    run ./cntb get role "$roleId"
    assert_success
    assert_output --partial "updated$name" 

    # cleanup
    run ./cntb delete role "$roleId"
    assert_success
}


@test "update role without name nor permissions : nok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb update role "$roleId"
    assert_failure

    # cleanup
    run ./cntb delete role "$roleId"
    assert_success
}

@test "update role not available action : nok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb update role "$roleId" --name="foo${TEST_SUFFIX}" -p '[{"apiName": "/v1/users", "actions": ["READ", "UPDATE"] }]'
    assert_failure

    # cleanup
    run ./cntb delete role "$roleId"
    assert_success
}

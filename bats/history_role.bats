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


@test "get role history: normal ok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb history roles
    assert_success
    assert_output --partial 'ID'
    assert_output --partial 'ROLEID'
    assert_output --partial 'ACTION'
    assert_output --partial 'USERNAME'
    assert_output --partial 'TIMESTAMP'

    # clean up
    run ./cntb delete role "$roleId"
    assert_success
}

@test "get role wide: ok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"

    run ./cntb history roles -o wide
    assert_success
    assert_output --partial 'ID'
    assert_output --partial 'ROLEID'
    assert_output --partial 'ACTION'
    assert_output --partial 'USERNAME'
    assert_output --partial 'CHANGEDBY'
    assert_output --partial 'REQUESTID'
    assert_output --partial 'TRACEID'
    assert_output --partial 'TENANTID'
    assert_output --partial 'CUSTOMERID'
    assert_output --partial 'TIMESTAMP'
    assert_output --partial 'CHANGED'

    # clean up
    run ./cntb delete role "$roleId"
    assert_success
}
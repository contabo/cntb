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

@test "get user history: normal ok" {
     run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId="$output"

    run ./cntb history users
    assert_success
    assert_output --partial 'ID'
    assert_output --partial 'USERID'
    assert_output --partial 'ACTION'
    assert_output --partial 'USERNAME'
    assert_output --partial 'TIMESTAMP'

    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role apiPermission "$roleId"
    assert_success
}

@test "get user wide: ok" {

   run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId="$output"

    run ./cntb history users -o wide
    assert_success
    assert_output --partial 'ID'
    assert_output --partial 'USERID'
    assert_output --partial 'ACTION'
    assert_output --partial 'USERNAME'
    assert_output --partial 'CHANGEDBY'
    assert_output --partial 'REQUESTID'
    assert_output --partial 'TRACEID'
    assert_output --partial 'TENANTID'
    assert_output --partial 'CUSTOMERID'
    assert_output --partial 'TIMESTAMP'
    assert_output --partial 'CHANGED'

    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role apiPermission "$roleId"
    assert_success

}
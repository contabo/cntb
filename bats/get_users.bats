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

@test "get users : ok" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId="$output"

    run ./cntb create user --firstName="bar${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="bar.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId2="$output"

    run ./cntb get users
    assert_success
    assert_output --partial "USERID"
    assert_output --partial "FIRSTNAME"
    assert_output --partial "LASTNAME"
    assert_output --partial "EMAIL"
    assert_output --partial "ENABLED"
    assert_output --partial "$userId"
    assert_output --partial "$userId2"

    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete user "$userId2"
    assert_success
    run ./cntb delete role apiPermission "$roleId"
    assert_success
}

@test "get users filter email: ok" {

    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId="$output"

    run ./cntb create user --firstName="bar${TEST_SUFFIX}" --lastName="barbar${TEST_SUFFIX}" --email="bar.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId2="$output"

    run ./cntb get users --email="bar.bar${TEST_SUFFIX}@contabo.com"
    assert_success
    assert_output --partial "barbar${TEST_SUFFIX}"


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete user "$userId2"
    assert_success
    run ./cntb delete role apiPermission "$roleId"
    assert_success
}
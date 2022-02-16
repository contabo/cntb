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

@test "update user ok: set firstName" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --locale de --roles="$roleId"
    assert_success
    userId="$output"


    run ./cntb update user "$userId" --firstName="foo1${TEST_SUFFIX}"
    assert_success

    run ./cntb get user "$userId"
    assert_output --partial "foo1${TEST_SUFFIX}"


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role "$roleId"
    assert_success
}

@test "update user ok: set lastName" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --locale de --roles="$roleId"
    assert_success
    userId="$output"


    run ./cntb update user "$userId" --lastName="foo1${TEST_SUFFIX}"
    assert_success

    run ./cntb get user "$userId"
    assert_output --partial "foo1${TEST_SUFFIX}"


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role "$roleId"
    assert_success
}

@test "update user ok: set email" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --locale de --roles="$roleId"
    assert_success
    userId="$output"


    run ./cntb update user "$userId" --email="testuser2${TEST_SUFFIX}@contabo.com"
    assert_success

    run ./cntb get user "$userId"
    assert_output --partial "testuser2${TEST_SUFFIX}@contabo.com"


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role "$roleId"
    assert_success
}
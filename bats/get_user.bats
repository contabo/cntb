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

@test "get user multiple output: ok" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId="$output"


    run ./cntb get user "$userId"
    assert_success
    assert_output --partial "USERID"
    assert_output --partial "FIRSTNAME"
    assert_output --partial "LASTNAME"
    assert_output --partial "EMAIL"
    assert_output --partial "ENABLED"


    run ./cntb get user "$userId" -o wide
    assert_success
    assert_output --partial "USERID"
    assert_output --partial "FIRSTNAME"
    assert_output --partial "LASTNAME"
    assert_output --partial "EMAIL"
    assert_output --partial "ENABLED"
    assert_output --partial "TOTP"
    assert_output --partial "ADMIN"

    run ./cntb get user "$userId" -o json
    assert_success
    assert_output --partial '"userId"'
    assert_output --partial '"firstName"'
    assert_output --partial '"lastName"'
    assert_output --partial '"email"'
    assert_output --partial '"enabled"'
    assert_output --partial '"totp"'
    assert_output --partial '"admin"'
    assert_output --partial '"roles"'

    run ./cntb get user "$userId" -o yaml
    assert_success
    assert_output --partial 'userId:'
    assert_output --partial 'firstName:'
    assert_output --partial 'lastName:'
    assert_output --partial 'email:'
    assert_output --partial 'enabled:'
    assert_output --partial 'totp:'
    assert_output --partial 'admin:'
    assert_output --partial 'roles:'
    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role apiPermission "$roleId"
    assert_success

}

@test "get user no id: nok" {
    run ./cntb get user a
    assert_failure

    run ./cntb get user
    assert_failure
}
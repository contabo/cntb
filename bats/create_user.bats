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
@test "create user normal: ok" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId="$output"


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role apiPermission "$roleId"
    assert_success
}

@test "create user duplicate: nok" {
    run ./cntb create role apiPermission --name="foo${TEST_SUFFIX}" --apiPermission='[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_success
    userId="$output"

    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_failure


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role apiPermission "$roleId"
    assert_success
}


@test "create user missing e-mail: nok" {
    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_failure

}

@test "create user missing firstName: nok" {
    run ./cntb create user  --lastName="bar${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com"  --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_failure
}

@test "create user missing lastName: nok" {
    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --email="foo.bar${TEST_SUFFIX}@contabo.com"  --enabled=true --admin=true --accessAllResources=true --roles="$roleId"
    assert_failure
}

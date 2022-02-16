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
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --roles="$roleId" --locale de
    assert_success
    userId="$output"


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role "$roleId"
    assert_success
}

@test "create user duplicate: nok" {
    run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
    assert_success
    roleId="$output"


    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --roles="$roleId" --locale de
    assert_success
    userId="$output"

    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --roles="$roleId" --locale de
    assert_failure


    #clean up
    run ./cntb delete user "$userId"
    assert_success
    run ./cntb delete role "$roleId"
    assert_success
}


@test "create user missing e-mail: nok" {
    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --enabled=true --roles="$roleId" --locale de
    assert_failure

    export CNTB_FIRSTNAME="foo${TEST_SUFFIX}"
    export CNTB_LASTNAME="bar${TEST_SUFFIX}"
    export CNTB_ENABLED=true
    export CNTB_LOCALE=de

    run ./cntb create user
    assert_failure

    unset CNTB_FIRSTNAME CNTB_LASTNAME CNTB_ENABLED CNTB_LOCALE
}

@test "create user missing firstName: nok" {
    run ./cntb create user  --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com"  --enabled=true --roles="$roleId" --locale de
    assert_failure

    export CNTB_LASTNAME="bar${TEST_SUFFIX}"
    export CNTB_ENABLED=true
    export CNTB_LOCALE=de
    export CNTB_EMAIL="testuser${TEST_SUFFIX}@contabo.com"

    run ./cntb create user
    assert_failure

    unset CNTB_EMAIL CNTB_LASTNAME CNTB_ENABLED CNTB_LOCALE
}

@test "create user missing lastName: nok" {
    run ./cntb create user --firstName="foo${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com"  --enabled=true --roles="$roleId" --locale de
    assert_failure

    export CNTB_FIRSTNAME="foo${TEST_SUFFIX}"
    export CNTB_ENABLED=true
    export CNTB_LOCALE=de
    export CNTB_EMAIL="testuser${TEST_SUFFIX}@contabo.com"

    run ./cntb create user
    assert_failure

    unset CNTB_FIRSTNAME CNTB_EMAIL CNTB_ENABLED CNTB_LOCALE
}

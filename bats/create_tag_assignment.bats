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

@test 'create tag assignment: ok : arguments tagId resourceType and resourceId' {
    run ./cntb create tag --name "foo${TEST_SUFFIX}"
    tagid="$output"

    run ./cntb create tagAssignment $tagid instance de-98546@contabo.net
    assert_success

    # clean up
    run ./cntb delete tagAssignment $tagid instance de-98546@contabo.net
    assert_success
    run ./cntb delete tag "$tagid"
    assert_success
}

@test 'create tag assignment: nok : duplicate assignment' {
    run ./cntb create tag --name "foo${TEST_SUFFIX}"
    tagid="$output"

    run ./cntb create tagAssignment $tagid instance de-98546@contabo.net
    assert_success

    run ./cntb create tagAssignment $tagid instance de-98546@contabo.net
    assert_failure

    # clean up
    run ./cntb delete tagAssignment $tagid instance de-98546@contabo.net
    assert_success
    run ./cntb delete tag "$tagid"
    assert_success
}

@test 'create tag assignment: nok : only 2 arguments' {
    run ./cntb create tagAssignment 2 de-98546@contabo.net
    assert_failure
}

@test 'create tag assignment: nok : wrong tagId type' {
    run ./cntb create tagAssignment instance instance de-98546@contabo.net
}
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

    run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
    assert_success

    # clean up
    run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
    assert_success
    run ./cntb delete tag "$tagid"
    assert_success
}

@test 'create tag assignment: nok : duplicate assignment' {
    run ./cntb create tag --name "foo${TEST_SUFFIX}"
    tagid="$output"

    run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
    assert_success

    run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
    assert_failure

    # clean up
    run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
    assert_success
    run ./cntb delete tag "$tagid"
    assert_success
}

@test 'create tag assignment: nok : only 2 arguments' {
    run ./cntb create tagAssignment 2 cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
    assert_failure
}

@test 'create tag assignment: nok : wrong tagId type' {
    run ./cntb create tagAssignment instance instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
}
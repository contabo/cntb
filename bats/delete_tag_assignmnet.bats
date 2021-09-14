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


@test 'delete tag assignment: ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb get tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_failure

  #clean
  run ./cntb delete tag ${tagid}
  assert_success
}

@test 'delete tag assignment: non existing nok' {
    run ./cntb delete tagAssignment 0 instance testw23
    assert_failure
    assert_output --partial 'Error'
    assert_output --partial '404,'
    assert_output --partial 'Error while deleting tag assignment'

}

@test 'delete tag assignmnet: 1 input nok' {
    run ./cntb delete tagAssignment 0
    assert_failure
}
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

  run ./cntb create tagAssignment $tagid instance de-98546@contabo.net
  assert_success

  run ./cntb delete tagAssignment $tagid instance de-98546@contabo.net
  assert_success

  run ./cntb get tagAssignment $tagid instance de-98546@contabo.net
  assert_failure

  #clean
  run ./cntb delete tag ${tagid}
  assert_success
}

@test 'delete tag assignment: non existing nok' {
    run ./cntb delete tagAssignment 0 instance testw23
    assert_failure
    assert_output --partial 'Error while deleting tag assignment: 404 - Entry Tag not found by tagId'
}

@test 'delete tag assignmnet: 1 input nok' {
    run ./cntb delete tagAssignment 0
    assert_failure
}
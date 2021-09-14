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

@test 'update tag: ok  : name and color by flag and env' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb update tag "${tagid}" --color "#000000"
  assert_success
  run ./cntb get tag "$tagid"
  assert_success
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "#000000"


  run ./cntb update tag "${tagid}" --name "bar${TEST_SUFFIX}" --color "#FFFFFF"
  assert_success
  run ./cntb get tag "$tagid"
  assert_success
  assert_output --partial "bar${TEST_SUFFIX}"
  assert_output --partial "#FFFFFF"


  export CNTB_COLOR="#000000"
  run ./cntb update tag "${tagid}"
  assert_success
  unset CNTB_COLOR
  run ./cntb get tag "$tagid"
  assert_success
  assert_output --partial "bar${TEST_SUFFIX}"
  assert_output --partial "#000000"

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'update tag: nok : non existing  tag' {
  run ./cntb update tag 0 --name "foo"
  assert_failure
  assert_output --partial 'Error'
  assert_output --partial '404,'
  assert_output --partial "Entry Tag not found by tagId = 0"
}

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

@test 'delete tag: ok  : existing tag' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb delete tag "$tagid"
  assert_success

  run ./cntb get tag "$tagid"
  assert_failure
}

@test 'delete tag: nok  : non existing tag' {
  run ./cntb delete tag 0
  assert_failure
  assert_output --partial 'Error'
  assert_output --partial '404,'
  assert_output --partial "Entry Tag not found by tagId = 0"
}


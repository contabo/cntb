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

@test 'get tag history: ok  : normal ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb history tags
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'TAGID'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'get tag history: ok  : wide ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb history tags -o wide
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'TAGID'
  assert_output --partial 'ACTION'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'CHANGES'


  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

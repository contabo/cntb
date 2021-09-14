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

@test 'get tag: ok  : existing tag with various output formats' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb get tag "$tagid"
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_output --partial "${tagid}"
  assert_output --partial "#0A78C3"

  run ./cntb get tag "$tagid" -o wide
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_output --partial "${tagid}"
  assert_output --partial "#0A78C3"

  run ./cntb get tag "$tagid" -o json
  assert_success
  assert_output --partial '"color"'
  assert_output --partial '"tagId"'
  assert_output --partial '"name"'
  assert_output --partial ": ${tagid}"
  assert_output --partial '"#0A78C3"'

  run ./cntb get tag "$tagid" -o yaml
  assert_success
  assert_output --partial 'color:'
  assert_output --partial 'tagId:'
  assert_output --partial 'name:'
  assert_output --partial "${tagid}"
  assert_output --partial "#0A78C3"


  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'get tag: nok : non existing tag' {
  run ./cntb get tag 0
  assert_failure
  assert_output --partial 'Error'
  assert_output --partial '404,'
  assert_output --partial "Entry Tag not found by tagId = 0"
}

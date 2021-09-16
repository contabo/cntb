#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
    ensureTestConfig
    deleteCache

    # create_a_lot_of_tags
}

function teardown_file() {
    # delete_a_lot_of_tags

    restore_config_files
}

@test 'get tag assignment history: normal ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb create tagAssignment $tagid instance de-98546@contabo.net
  assert_success

  run ./cntb history tagAssignment --orderBy=tag_id:asc
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'RESOURCEID'
  assert_output --partial 'RESOURCETYPE'

    # clean up
  run ./cntb delete tagAssignment $tagid instance de-98546@contabo.net
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success

}




@test 'get tag assignment  history: ok  : wide ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb create tagAssignment $tagid instance de-98546@contabo.net
  assert_success

  run ./cntb history tagAssignment -s 4 -o wide
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'RESOURCEID'
  assert_output --partial 'RESOURCETYPE'
  assert_output --partial 'TAGID'
  assert_output --partial 'ACTION'
  assert_output --partial 'CHANGEDBY'
  assert_output --partial 'USERNAME'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'CHANGES'


  # clean up
  run ./cntb delete tagAssignment $tagid instance de-98546@contabo.net
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success
}

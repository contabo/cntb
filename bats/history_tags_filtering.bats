#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
    ensureTestConfig
    deleteCache

    create_a_lot_of_tags
}

function teardown_file() {

    delete_a_lot_of_tags

    restore_config_files
}




function create_a_lot_of_tags() {
  > /tmp/historytagsfiltering_tagid
  for no in {1..30}; do
    run ./cntb create tag --name "foo${no}${TEST_SUFFIX}"
    assert_success
    echo ${output} >> /tmp/historytagsfiltering_tagid
  done

}

function delete_a_lot_of_tags() {
  cat /tmp/historytagsfiltering_tagid | while read elem;do
    run ./cntb delete tag ${elem}
    assert_success
  done
}


@test 'history tags: ok  : pagination' {
  # up to 100
  run ./cntb history tags
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'ID'
  assert_output --partial 'ACTION'
  assert [ "${#lines[@]}" -ge "31" ]

  # request nonexisting page
  run ./cntb history tags --page 10000
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'ID'
  assert_output --partial 'ACTION'
  assert_equal "${#lines[@]}" "1"

  # page size = 5
  run ./cntb history tags --size 5
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'ID'
  assert_output --partial 'ACTION'
  assert_equal "${#lines[@]}" "6"

  run ./cntb history tags --size 5 --page 2
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'ID'
  assert_output --partial 'ACTION'
  assert_equal "${#lines[@]}" "6"

  # page size 10
  run ./cntb history tags --size 10
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'ID'
  assert_output --partial 'ACTION'
  assert_equal "${#lines[@]}" "11"

  run ./cntb history tags --size 10 --page 2
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'ID'
  assert_output --partial 'ACTION'
  assert_equal "${#lines[@]}" "11"
}


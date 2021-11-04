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

declare -a tagids

function teardown_file() {
    delete_a_lot_of_tags

    restore_config_files
}

function create_a_lot_of_tags() {
  > /tmp/gettagsfiltering_tagid
  for no in {1..30}; do
    run ./cntb create tag --name "bar${no}${TEST_SUFFIX}"
    assert_success
    echo ${output} >> /tmp/gettagsfiltering_tagid
  done
   export tagids=${tagids[@]}
}


function delete_a_lot_of_tags() {
  cat /tmp/gettagsfiltering_tagid | while read elem;do
    run ./cntb delete tag ${elem}
    assert_success
  done
}


@test 'get tags: ok  : pagination' {
  # up to 100
  run ./cntb get tags -t bar
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'

  # request nonexisting page
  run ./cntb get tags --page 100
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_equal "${#lines[@]}" "1"

  # page size = 5
  run ./cntb get tags --size 5
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_equal "${#lines[@]}" "6"

  run ./cntb get tags --size 5 --page 2 --tagName="bar" --orderBy "id:asc"
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_equal "${#lines[@]}" "6"
  assert_line --index 2 --partial "bar7"


  # page size 10
  run ./cntb get tags --size 10
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_equal "${#lines[@]}" "11"

  run ./cntb get tags --size 10 --page 2 --tagName="bar" --orderBy "id:asc"
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_equal "${#lines[@]}" "11"
  assert_line --index 2 --partial "bar12"
}

@test 'get tags: ok  : filtering' {
  # up to 100
  run ./cntb get tags --tagName=bar11
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'
  assert_line --index 1 --partial "bar11"
}

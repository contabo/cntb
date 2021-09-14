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

###
### arugments
###

@test 'create tag: ok  : arguments only name' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'create tag: ok  : arguments name and color' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}" --color "#000000"
  assert_success
  tagid="$output"

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'create tag: nok : arguments name invalid color' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}" --color "000000"
  assert_failure
}

@test 'create tag: nok : arguments duplicate names' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_failure

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}


###
### stdin
###

wrapperCreateTag() {
  echo "$1" | ./cntb create tag -f -
}
@test 'create tag: ok  : stdin only name' {
  run wrapperCreateTag "{ \"name\": \"foo${TEST_SUFFIX}\" }"
  assert_success
  tagid="$output"

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'create tag: ok  : stdin name and color' {
  run wrapperCreateTag "{ \"name\": \"foo${TEST_SUFFIX}\", \"color\": \"#000000\" }"
  assert_success
  tagid="$output"

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'create tag: nok : stdin name invalid color' {
  run wrapperCreateTag "{ \"name\": \"foo${TEST_SUFFIX}\", \"color\": \000000\" }"
  assert_failure
}

@test 'create tag: nok : stdin duplicate names' {
  run wrapperCreateTag "{ \"name\": \"foo${TEST_SUFFIX}\" }"
  assert_success
  tagid="$output"

  run wrapperCreateTag "{ \"name\": \"foo${TEST_SUFFIX}\" }"
  assert_failure

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

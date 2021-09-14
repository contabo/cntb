
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

@test 'get tags: normal ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb get tags
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'get tags: wide ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb get tags -o wide
  assert_success
  assert_output --partial 'COLOR'
  assert_output --partial 'TAGID'
  assert_output --partial 'NAME'

  # clean up
  run ./cntb delete tag "$tagid"
  assert_success
}

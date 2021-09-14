#!/usr/bin/env bats

load handling_conf_files.bash
load globals.bash
load_lib bats-support
load_lib bats-assert

function setup_file() {
    store_config_files
}

function teardown_file() {
    restore_config_files
}

@test 'cntb: ok  : returning help' {
  run ./cntb
  assert_success
  assert_output --partial 'cntb is the command-line interface (CLI)'

  run ./cntb --help
  assert_success
  assert_output --partial 'cntb is the command-line interface (CLI)'

  run ./cntb asdhfjkdhfkdsfjsadh
  assert_success
  assert_output --partial 'cntb is the command-line interface (CLI)'
}

@test 'cntb version: ok  : returning version' {
  run ./cntb version --help
  assert_success
  assert_output --partial 'Shows the current version of cntb CLI.'

  run ./cntb version
  assert_success
  assert_output --partial 'cntb'

  run ./cntb version -v
  assert_success
  assert_output --partial 'cntb'
  assert_output --partial 'commit:'
  assert_output --partial '(on '
}


@test 'cntb suggestions: ok  : returning alternatives' {
  run ./cntb crate
  assert_failure
  assert_output --partial 'Error: invalid argument "crate" for "cntb"'
  assert_output --partial "Did you mean this?"
  assert_line --regexp '^[[:space:]]create$'

  run ./cntb new
  assert_failure
  assert_output --partial 'Error: invalid argument "new" for "cntb"'
  assert_output --partial 'Did you mean this?'
  assert_line --regexp '^[[:space:]]create$'
}

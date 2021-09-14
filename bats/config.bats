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

@test 'config: ok  : returning help' {
  run ./cntb config
  assert_success
  assert_output --partial 'View and edit config files.'
  assert_output --partial '$HOME/.cntb.yml'
  assert_output --partial 'CNTB_'

  run ./cntb config --help
  assert_success
  assert_output --partial 'View and edit config files.'
  assert_output --partial '$HOME/.cntb.yml'
  assert_output --partial 'CNTB_'
}

@test 'config: ok  : viewing config' {
  run ./cntb config view --help
  assert_success
  assert_output --partial 'Shows the current configuration from various sources'

  run ./cntb config view
  assert_success
  assert_output --partial '--api'
  assert_output --partial '--config'
  assert_output --partial '--debug'
  assert_output --partial '--oauth2-'
}

@test 'config: ok  : setting debug level' {
  run ./cntb config set-debug-level --debug trace

  run ./cntb config view
  assert_success
  assert_output --partial '--debug trace'

  run ./cntb config set-debug-level --debug warn
  assert_success

  run ./cntb config view
  assert_success
  assert_output --partial '--debug warn'
}

@test 'config: nok : not setting debug level' {
  run ./cntb config set-debug-level trace
  assert_failure
  assert_output --partial 'Please provide debug level via --debug'
}



@test 'config: ok  : setting credentials' {
  run ./cntb config set-credentials --oauth2-client-secret a --oauth2-clientid b --oauth2-password c --oauth2-tokenurl d --oauth2-user e
  assert_success

  run ./cntb config view
  assert_success
  assert_output --partial '--oauth2-client-secret a'
  assert_output --partial '--oauth2-clientid b'
  assert_output --partial '--oauth2-password c'
  assert_output --partial '--oauth2-tokenurl d'
  assert_output --partial '--oauth2-user e'

  run ./cntb config set-credentials --oauth2-client-secret e --oauth2-clientid d --oauth2-password pass --oauth2-tokenurl b --oauth2-user a
  assert_success

  run ./cntb config view
  assert_success
  assert_output --partial '--oauth2-client-secret e'
  assert_output --partial '--oauth2-clientid d'
  assert_output --partial '--oauth2-password pass'
  assert_output --partial '--oauth2-tokenurl b'
  assert_output --partial '--oauth2-user a'
}



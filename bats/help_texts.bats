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

@test 'help texts to be displayed : ok' {
  run ./cntb
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb [flags]'

  run ./cntb config
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb config'

  run ./cntb create
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb create'

  run ./cntb delete
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb delete'

  run ./cntb edit
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb edit'

  run ./cntb generateSecret
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb generateSecret'

  run ./cntb help
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb [flags]'

  run ./cntb rollback
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb rollback'

  run ./cntb start
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb start'

  run ./cntb stop
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb stop'


  run ./cntb update
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb update'

  run ./cntb version
  assert_success
  assert_output --partial 'cntb'

  run ./cntb get
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb get'

  run ./cntb history
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb history'

  run ./cntb completion
  assert_success
  assert_output --partial 'Usage:'
  assert_output --partial 'cntb completion'
}


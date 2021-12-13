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

@test 'get images stats : ok ' {
  run ./cntb get images-stats
  assert_success
  assert_output --partial 'TOTALSIZEMB'
  assert_output --partial 'FREESIZEMB'
  assert_output --partial 'USEDSIZEMB'
  assert_output --partial 'CURRENTIMAGESCOUNT'
}

@test 'get images stats wide : ok ' {
  run ./cntb get images-stats -o wide
  assert_success
  assert_output --partial 'TOTALSIZEMB'
  assert_output --partial 'FREESIZEMB'
  assert_output --partial 'USEDSIZEMB'
  assert_output --partial 'CURRENTIMAGESCOUNT'
}

@test 'get images stats json : ok ' {
  run ./cntb get images-stats -o json
  assert_success
  assert_output --partial '"currentImagesCount"'
  assert_output --partial '"freeSizeMb"'
  assert_output --partial '"totalSizeMb"'
  assert_output --partial '"usedSizeMb"'
}
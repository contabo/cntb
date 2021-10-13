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

@test 'delete image : ok' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id="$output"

  run ./cntb delete image $image_id
  assert_success
}

@test 'delete image : nok : not existing' {
  run ./cntb delete image a19b8645-79d2-4fbb-ac56-a927d69b8d2b
  assert_failure
  assert_output --partial 'Error'
  assert_output --partial '404'
}
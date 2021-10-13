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

@test 'get images : ok' {
  run ./cntb get images
  assert_success
  assert_output --partial 'IMAGEID'
  assert_output --partial 'NAME'
  assert_output --partial 'SIZEMB'
  assert_output --partial 'OSTYPE'
  assert_output --partial 'VERSION'
  assert_output --partial 'STATUS'
}

@test 'get images : ok : format wide' {
  run ./cntb get images -o wide
  assert_success
  assert_output --partial 'IMAGEID'
  assert_output --partial 'NAME'
  assert_output --partial 'SIZEMB'
  assert_output --partial 'OSTYPE'
  assert_output --partial 'TAGS'
  assert_output --partial 'VERSION'
  assert_output --partial 'URL'
  assert_output --partial 'DESCRIPTION'
  assert_output --partial 'STATUS'
  assert_output --partial 'ERRORMESSAGE'
  assert_output --partial 'STANDARDIMAGE'
}

@test 'get images : ok : format json' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id="$output"

  run ./cntb get images -o json
  assert_success
  assert_output --partial '"imageId"'
  assert_output --partial '"name"'
  assert_output --partial '"sizeMb"'
  assert_output --partial '"osType"'
  assert_output --partial '"tags"'
  assert_output --partial '"version"'
  assert_output --partial '"url"'
  assert_output --partial '"description"'
  assert_output --partial '"status"'
  assert_output --partial '"errorMessage"'
  assert_output --partial '"standardImage"'

  run ./cntb delete image "$image_id"
  assert_success
}

@test 'get images : ok : filter by name' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id="$output"

  run ./cntb get images --name "image${TEST_SUFFIX}"
  assert_success
  assert_output --partial 'IMAGEID'
  assert_output --partial 'NAME'
  assert_output --partial 'SIZEMB'
  assert_output --partial 'OSTYPE'
  assert_output --partial 'VERSION'
  assert_output --partial "image${TEST_SUFFIX}"
  assert_output --partial "$image_id"

  run ./cntb delete image "$image_id"
  assert_success
}
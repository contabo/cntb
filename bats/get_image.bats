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

@test 'get image : ok ' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id="$output"

  run ./cntb get image "$image_id"
  assert_success
  assert_output --partial 'IMAGEID'
  assert_output --partial 'NAME'
  assert_output --partial 'SIZEMB'
  assert_output --partial 'UPLOADEDSIZEMB'
  assert_output --partial 'OSTYPE'
  assert_output --partial 'VERSION'
  assert_output --partial 'STATUS'
  assert_output --partial "${image_id}"
  assert_output --partial "image${TEST_SUFFIX}"

  run ./cntb delete image "$image_id"
  assert_success
}

@test 'get image : ok : format wide' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id="$output"

  run ./cntb get image "$image_id" -o wide
  assert_success
  assert_output --partial 'IMAGEID'
  assert_output --partial 'NAME'
  assert_output --partial 'SIZEMB'
  assert_output --partial 'UPLOADEDSIZEMB'
  assert_output --partial 'OSTYPE'
  assert_output --partial 'VERSION'
  assert_output --partial 'STATUS'
  assert_output --partial 'STANDARDIMAGE'
  assert_output --partial 'URL'
  assert_output --partial 'DESCRIPTION'
  assert_output --partial 'ERRORMESSAGE'
  assert_output --partial "$image_id"
  assert_output --partial "image${TEST_SUFFIX}"

  run ./cntb delete image "$image_id"
  assert_success
}

@test 'get image : ok : format json' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id="$output"

  run ./cntb get image "$image_id" -o json
  assert_success
  assert_output --partial '"imageId"'
  assert_output --partial '"name"'
  assert_output --partial '"sizeMb"'
  assert_output --partial '"uploadedSizeMb"'
  assert_output --partial '"osType"'
  assert_output --partial '"version"'
  assert_output --partial '"status"'
  assert_output --partial '"standardImage"'
  assert_output --partial '"url"'
  assert_output --partial '"description"'
  assert_output --partial '"errorMessage"'
  assert_output --partial "$image_id"
  assert_output --partial "image${TEST_SUFFIX}"

  run ./cntb delete image "$image_id"
  assert_success
}

@test 'get image : nok : missing imageId' {
  run ./cntb get image
  assert_failure
}

@test 'get image : nok : imageId not uuidV4' {
  run ./cntb get image thisisnotauuidV4
  assert_failure
}
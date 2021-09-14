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

@test 'create image : ok : duplicate : nok' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id=$(echo "$output" | sed -n 's/.*imageId\s\+\([0-9a-zA-Z-]\+\).*$/\1/p')

  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_failure

  run ./cntb delete image "$image_id"
  assert_success
}

@test 'create image : nok : missing name' {
  run ./cntb create image --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_failure
}

@test 'create image : nok : missing description' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_failure
}

@test 'create image : nok : missing url' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --osType Linux --version 20.04
  assert_failure
}

@test 'create image : nok : missing osType' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --version 20.04
  assert_failure
}

@test 'create image : nok : missing version' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux
  assert_failure
}
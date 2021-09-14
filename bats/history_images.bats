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

@test 'get image history : ok' {
  run ./cntb history images
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'IMAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
}

@test 'get image history : ok : format wide' {
  run ./cntb history images -o wide
  assert_success
  assert_output --partial 'ID'
  assert_output --partial 'IMAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'USERNAME'
  assert_output --partial 'CHANGEDBY'
  assert_output --partial 'REQUESTID'
  assert_output --partial 'TRACEID'
  assert_output --partial 'CHANGES'
}

@test 'get image history : ok : format json' {
  run ./cntb history images -o json
  assert_success
  assert_output --partial '"id"'
  assert_output --partial '"imageId"'
  assert_output --partial '"action"'
  assert_output --partial '"timestamp"'
  assert_output --partial '"username"'
  assert_output --partial '"changedBy"'
  assert_output --partial '"requestId"'
  assert_output --partial '"traceId"'
  assert_output --partial '"changes"'
}

@test 'get image history : ok : filter by imageId' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id=$(echo "$output" | sed -n 's/.*imageId\s\+\([0-9a-zA-Z-]\+\).*$/\1/p')

  run ./cntb delete image "$image_id"
  assert_success

  run ./cntb history images --imageId "$image_id"
  assert_success
  # header
  assert_output --partial 'ID'
  assert_output --partial 'IMAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  # content
  assert_output --partial 'CREATED'
  assert_output --partial 'DELETED'
}

@test 'get image history : ok : filter by changedBy' {
  run ./cntb create image --name "image${TEST_SUFFIX}" --description 'description' --url "${IMAGE_DOWNLOAD_URL}" --osType Linux --version 20.04
  assert_success

  image_id=$(echo "$output" | sed -n 's/.*imageId\s\+\([0-9a-zA-Z-]\+\).*$/\1/p')

  run ./cntb delete image "$image_id"
  assert_success

  run ./cntb history images --changedBy "${USER_ID}" -o wide
  assert_success
  # header
  assert_output --partial 'ID'
  assert_output --partial 'IMAGEID'
  assert_output --partial 'ACTION'
  assert_output --partial 'TIMESTAMP'
  assert_output --partial 'CHANGEDBY'
  # content
  assert_output --partial 'CREATED'
  assert_output --partial 'DELETED'
  assert_output --partial "${USER_ID}"
}

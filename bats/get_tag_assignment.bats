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

@test 'get tag assignment: existing tag assignment ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb get tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'RESOURCENAME'
  assert_output --partial 'TAGNAME'
  assert_output --partial 'RESOURCETYPE'
  assert_output --partial 'RESOURCEID'
  assert_output --partial "$tag"
  assert_output --partial "test_vps"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"
  assert_output --partial "cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce"


  run ./cntb get tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'RESOURCENAME'
  assert_output --partial 'TAGNAME'
  assert_output --partial 'RESOURCETYPE'
  assert_output --partial 'RESOURCEID'
  assert_output --partial "$tag"
  assert_output --partial "test_vps"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"
  assert_output --partial "cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce"

  run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'get tag assignment: format wide ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb get tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  assert_output --partial 'TAGID'
  assert_output --partial 'RESOURCENAME'
  assert_output --partial 'TAGNAME'
  assert_output --partial 'RESOURCETYPE'
  assert_output --partial 'RESOURCEID'
  assert_output --partial "$tag"
  assert_output --partial "test_vps"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"
  assert_output --partial "cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce"

  run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success

}

@test 'get tag assignments: format json ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

  run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb get tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce -o json
  assert_success
  assert_output --partial '"resourceId"'
  assert_output --partial '"resourceName"'
  assert_output --partial '"resourceType"'
  assert_output --partial '"tagId"'
  assert_output --partial '"tagName"'
  assert_output --partial '"cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce"'
  assert_output --partial "test_vps"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"


  run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success

}

@test 'get tag assignmnet: 1 input nok' {
  run ./cntb get tagAssignment 1
  assert_failure
}


@test 'get tag assignmnet: 2 inputs nok' {
  run ./cntb get tagAssignment 1 instance
  assert_failure
}


@test 'get tag assignmnet: no number tagId nok' {
  run ./cntb get tagAssignment string instance customer-admin-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce-integration
  assert_failure
}
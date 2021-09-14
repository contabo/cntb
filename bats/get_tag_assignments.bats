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

@test 'get tag assignments: normal ok' {
  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"

   run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb get tagAssignments $tagid
  assert_success
  assert_output --partial "TAGID"
  assert_output --partial "TAGNAME"
  assert_output --partial "RESOURCETYPE"
  assert_output --partial "RESOURCEID"
  assert_output --partial "RESOURCENAME"

  run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'get tag assignments: multiple tags ok' {

  run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"


  run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb create tagAssignment $tagid instance cli-tool2-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb get tagAssignments $tagid
  assert_success
  assert_output --partial "TAGID"
  assert_output --partial "TAGNAME"
  assert_output --partial "RESOURCETYPE"
  assert_output --partial "RESOURCEID"
  assert_output --partial "RESOURCENAME"
  assert_output --partial "$tagid"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"
  assert_output --partial "cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce"
  assert_output --partial "test_vps"
  assert_output --partial "$tagid"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"
  assert_output --partial "cli-tool2-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce"
  assert_output --partial "test_vps"

  run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tagAssignment $tagid instance cli-tool2-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'get tag assignments: wide ok' {
      run ./cntb create tag --name "foo${TEST_SUFFIX}"
  assert_success
  tagid="$output"


  run ./cntb create tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb create tagAssignment $tagid instance cli-tool2-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success

  run ./cntb get tagAssignments $tagid -o wide
  assert_success
  assert_output --partial "TAGID"
  assert_output --partial "TAGNAME"
  assert_output --partial "RESOURCETYPE"
  assert_output --partial "RESOURCEID"
  assert_output --partial "RESOURCENAME"
  assert_output --partial "$tagid"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"
  assert_output --partial "test_vps"
  assert_output --partial "$tagid"
  assert_output --partial "foo${TEST_SUFFIX}"
  assert_output --partial "instance"
  assert_output --partial "cli-tool2-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce"
  assert_output --partial "test_vps"

  run ./cntb delete tagAssignment $tagid instance cli-tool-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tagAssignment $tagid instance cli-tool2-integration-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce
  assert_success
  run ./cntb delete tag "$tagid"
  assert_success
}

@test 'get tag assignments: wrong tagid nok' {
    run ./cntb get tagAssignments test instance customer-admin-59a80d5d-1027-4d9f-94e4-1c87e8eef8ce-integration
    assert_failure
}

@test 'get tag assignments: 2 inputs nok' {
    run ./cntb get tagAssignments 1 instance
    assert_failure
}

@test 'get tag assignments: no inputs nok' {
    run ./cntb get tagAssignments
    assert_failure
}
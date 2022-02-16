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

###
### arguments
###

@test 'resend email verification : ok' {
  run ./cntb create role -n "foo${TEST_SUFFIX}" -p '[{"apiName" : "/v1/users", "actions": ["READ", "CREATE"]}]'
  assert_success
  roleId="$output"


  run ./cntb create user --firstName="foo${TEST_SUFFIX}" --lastName="bar${TEST_SUFFIX}" --email="testuser${TEST_SUFFIX}@contabo.com" --enabled=true --locale de --roles="$roleId"
  assert_success
  userId="$output"

  run ./cntb resendEmailVerification user "$userId"
  assert_success

  #clean up
  run ./cntb delete user "$userId"
  assert_success
  run ./cntb delete role "$roleId"
  assert_success
  
}

@test 'resend email verification invalid user id : nok' {
  run ./cntb resendEmailVerification user 32163657126573126535612
  assert_failure
}

@test 'resend email verification wrong user id : nok' {
  run ./cntb resendEmailVerification user 78953218-1ba4-4b9b-ae32-1c5ebf037ac0
  assert_failure
}

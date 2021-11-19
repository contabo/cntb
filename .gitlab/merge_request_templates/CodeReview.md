## Code Review

Take code review into consideration while doing story point estimations

* Multiple people doing code review (more than 1)
* Security Code review
  - [ ] No Passwords/Sensitive data in logs or in audit database
  - [ ] No customer can access data of another customer by mistake
  - [ ] No sensitive information in exceptions
* Functionality Code review
  - [ ] Does the code do what is required from the ticket
  - [ ] Input/Output validation
  - [ ] Error Handling
  - [ ] Deploying only on single environment
  - [ ] Do the changes cause something else to break??
  - [ ] Handle 3rd party software (error handling and don't trust that it will always work)
  - [ ] Are all translations available in frontend code
* Test Code review
  - [ ] Do tests make sense
  - [ ] Are there corner cases missing?
  - [ ]  Are there tests at all
  - [ ] Are tests covering acceptance Criteria in ticket
  - [ ]  Are there negative tests
  - [ ]  No customer can access data of another customer by mistake0

## What has been done
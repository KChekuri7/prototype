---
description: Implements API Endpoints and Unit Tests based on approved User Story and Design.
---

1. **Schema Verification & Design**:
   - **Check**: Does the required MongoDB collection exist in the **Local MongoDB**?
   - **If Missing**:
     - **Action**: Carefully design the schema (BSON structure) based on the User Story and data requirements.

2. **Schema Approval (Mandatory)**:
   - **Action**: Present the proposed or verified schema to the user for explicit approval.
   - **Action**: Explicitly ask: "Does the proposed database schema meet your requirements? Please approve to proceed to implementation, or provide feedback for revisions."
   - **STOP** and wait for user input.
   - **If Changes Requested**:
     - Analyze the user's feedback.
     - Modify the schema design to incorporate requested changes.
     - **Loop Back**: Step 1.

3. **Implement API Endpoints**:
   - Input: The approved User Story (`user_stories/[feature_name_snake_case].md`) and Design Specification (`designs/[feature_name_snake_case]_ux.md`).
   - **Action**: Use the `Golang API Creator` skill (`.agent/skills/golang_api_creator/SKILL.md`) and strictly follow rules in `.agent/rules/golang_api.md`.
   - **Action**: Connect to **Local MongoDB** for all data operations.
   - **Action**: Prioritize using the `go-util` framework for database, errors, and utilities.
   - **Action**: Implement models, services, and controllers in the `api/` directory.

4. **Review & Approval (API Implementation)**:
   - Present the implemented API code to the user (models, services, and controllers).
   - Explicitly ask: "Does the implemented API meet your requirements? Please approve to proceed to Testing, or provide feedback for revisions."
   - **STOP** and wait for user input.

5. **Handle Feedback (API Implementation)**:
   - **If Approved**: 
     - Proceed to Step 6.
   - **If Changes Requested**:
     - Analyze the user's feedback.
     - Modify the API code to incorporate requested changes.
     - **Action**: Update relevant `.go` files in the `api/` directory.
     - **Loop Back**: Step 4.

6. **Write API Test Cases**:
   - **Action**: Create test files (`{feature}_test.go`) co-located with the service implementation in `src/services/{feature}/`.
   - **Requirement**: Implement comprehensive test cases covering happy paths and edge cases.
   - **Action**: Interact with **Local MongoDB** to `create`, `update`, and `get` real test data during the tests to verify persistence.

7. **Run & Verify Tests**:
   - **Action**: Use the `run_command` tool to execute `go test -v ./src/services/...` for the modified feature.
   - **Verification**: 
     - **Strictly validate** that the endpoint/service returns a **SUCCESS** response.
     - Ensure data integrity in the Local MongoDB.
   - **Resolution**: If **ANY** errors occur (test failures, panic, data mismatch):
     - **Action**: Immediately analyze the error output.
     - **Action**: Fix the code/tests.
     - **Loop**: Re-run tests until 100% success is achieved.

8. **Postman Validation & Collection**:
   - **Action**: Start the API server and validate all endpoints using Postman (or equivalent HTTP client).
   - **Requirement**: Verify that every request receives a **200 OK** response.
   - **Action**: Create a Postman collection JSON file (`postman_collection.json`) in the project root containing all successfully validated requests.
   - **Resolution**: If any endpoint returns an error or unexpected status code:
     - **Action**: Fix the errors in the API code.
     - **Action**: Re-run server and validate again until all requests receive a 200 response and are added to the collection.
   - **Finalization**: Once all endpoints are validated with a 200 response and the collection is created, Workflow Complete.
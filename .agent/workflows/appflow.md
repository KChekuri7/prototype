---
description: Generates a User Story, UX Design Specification, and implements both UI Screen and API for a feature, with approval loops at each stage.
---

1. **User Story Phase**:
   - **Action**: Execute the `.agent/workflows/user_story_flow.md` workflow.
   - **Goal**: Draft, review, and approve a User Story for the feature.

2. **UX Design Phase**:
   - **Action**: Execute the `.agent/workflows/ux_design_flow.md` workflow.
   - **Goal**: Generate and approve UX Design Specifications and Visual Mockups based on the User Story.

3. **UI Implementation Phase**:
   - **Action**: Execute the `.agent/workflows/ui_screen_implementation_flow.md` workflow.
   - **Goal**: Implement the UI screens using React + Vite based on the approved design.

4. **API Implementation Phase**:
   - **Action**: Execute the `.agent/workflows/api_flow.md` workflow.
   - **Goal**: Implement the API endpoints using Golang, including Unit Tests.

5. **Integration Phase**:
   - **Action**: Execute the `.agent/workflows/integration_flow.md` workflow.
   - **Goal**: Connect the UI and API, ensuring data flow is correct.

6. **Testing & Validation Phase**:
   - **Action**: Execute the `.agent/workflows/testing_flow.md` workflow.
   - **Goal**: Perform End-to-End testing and validate the feature against the User Story.

## Skipping Steps / Resuming
**You can start this workflow at any specific phase.**
If the user requests to start at a later step (e.g., "Start at Step 3"):
1. **Context Check**: Verify that the required outputs/artifacts from the *previous* steps are:
   - Available in the current workspace (e.g., inside `doc/`, `brain/`, or source files).
   - OR provided explicitly in the user's prompt.
2. **Missing Context**: If the required context is missing, do **not** proceed. Stop and ask the user to provide the missing artifacts (e.g., "Please provides the User Story and UX Design before I can start implementation").
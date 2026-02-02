---
description: Performs End-to-End testing and QA validation.
---

1. **Testing & QA Validation**:
   - **Action**: Start the frontend (`npm run dev`) and backend (`go run main.go`) servers locally.
   - **Action**: Use the `browser_subagent` or `read_browser_page` tools to perform end-to-end testing.
   - **Verification**: Manually verify every scenario defined in the **Acceptance Criteria** of the approved User Story.

2. **Final Analysis & Refinement**:
   - **Action**: Contrast the actual browser behavior with the User Story and UX Design.
   - **Requirement**: If any functionality is missing or behavior is incorrect, address the missing parts.
   - **Completion**: Once the browser validation confirms the feature is 100% complete and matches the User Story, the feature development is finished.

---
description: Implements UI Screen based on approved UX Design.
---

1. **Implement UI Screen**:
   - Input: The approved UX Design Specification from `designs/[feature_name_snake_case]_ux.md`.
   - **Action**: Use the `UI Implementer` skill (`.agent/skills/ui_implementer/SKILL.md`) to implement the screen in the React + Vite application.
   - **Action**: Create/Update relevant `.tsx` and `.scss` files.

2. **Review & Approval (Implemented UI)**:
   - Present the implemented UI to the user (screenshots, recordings, or local dev URL).
   - Explicitly ask: "Does the implemented UI screen meet your expectations? Please approve to proceed to API Implementation, or provide feedback for refinements."
   - **STOP** and wait for user input.

3. **Handle Feedback (Implemented UI)**:
    - **If Approved**: 
      - Workflow Complete.
    - **If Changes Requested**:
      - Analyze the user's feedback.
      - Modify the code to incorporate requested changes.
      - **Action**: Update relevant `.tsx` and `.scss` files.
      - **Loop Back**: Go to Step 2.

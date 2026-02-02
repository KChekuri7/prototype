---
description: Generates a User Story for a feature, with approval loops.
---

1. **Input Collection**: Request the user to provide the feature description or core requirement if not already present in the context.

2. **Generate User Story**:
   - Use the `User Story Creator` skill (`.agent/skills/user_story_creator/SKILL.md`) to draft a detailed User Story based on the input.
   - **Action**: Save the generated content to a new file in `user_stories/`.
     - Filename format: `user_stories/[feature_name_snake_case].md` (e.g., `user_stories/login_auth.md`).

3. **Review & Approval (User Story)**:
   - Present the generated User Story to the user.
   - Explicitly ask: "Does this User Story meet your requirements? Please approve to proceed to UX Design, or provide feedback for revisions."
   - **STOP** and wait for user input.

4. **Handle Feedback (User Story)**:
   - **If Approved**: 
     - Workflow Complete.
   - **If Changes Requested**: 
     - Analyze the user's feedback.
     - Modify the User Story to incorporate the requested changes.
     - **Action**: Update the file `user_stories/[feature_name_snake_case].md` with the revised content.
     - **Loop Back**: Go to Step 3.

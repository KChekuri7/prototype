---
description: Generates UX Design Specification and visual mockup based on approved User Story.
---

1. **Generate UX Design**:
   - Input: The approved User Story from `user_stories/[feature_name_snake_case].md`.
   - **Action**: Use the `UX Designer` skill (`.agent/skills/ux_designer/SKILL.md`) to generate a Design Specification and a visual mockup.
   - **Action**: Save the generated image and Design Specification in the `designs/` directory.

2. **Review & Approval (UX Design)**:
   - Present the generated UX Design Specification and visual mockup to the user.
   - Explicitly ask: "Does this UX Design meet your requirements? Please approve to proceed to UI Screen Implementation, or provide feedback for revisions."
   - **STOP** and wait for user input.

3. **Handle Feedback (UX Design)**:
   - **If Approved**: 
     - Workflow Complete.
   - **If Changes Requested**:
     - Analyze the user's feedback.
     - Modify the UX Design Specification and/or regenerate the visual mockup.
     - **Action**: Update the files in the `designs/` directory with revised content.
     - **Loop Back**: Go to Step 2.

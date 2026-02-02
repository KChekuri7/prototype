---
description: Integrates UI and API, resolving connectivity issues.
---

1. **UI & API Integration**:
   - **Action**: Configure frontend proxy (if needed) to route API requests to backend.
   - **Action**: Use the `postman_collection.json` generated in the API flow as a reference for request methods, URLs, and payloads.
   - **Verification**: Ensure that the UI sends requests and handles responses exactly as defined in the Postman collection.
   - **Action**: Fix any CORS, payload format (JSON vs Headers), or routing issues.
   - **Verification**: Successfully trigger the API from the UI and receive a **200 OK** response (or the response defined in the collection).
   - **Requirement**: If integration fails or data doesn't match the collection, fix the UI code or return to API implementation flows if the collection itself needs updates.
   - **Finalization**: Once UI and API are perfectly synchronized according to the Postman collection, Integration Complete.

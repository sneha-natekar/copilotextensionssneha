**Copilot Skillsets Cheat Sheet**

**High-Level Summary of Components Needed to Build a GitHub Copilot
Extension Using the Skillsets Approach**

A **GitHub Copilot Extension** using the **skillsets** approach consists
of the following key components:

**1. Skillset Definition (API Endpoints)**

-   A **skillset** is a collection of up to **5 API endpoints** that
    Copilot can call.

-   Each endpoint provides structured responses that help Copilot
    **retrieve external data** or **perform specific actions**.

-   These APIs must return **machine-readable structured data (JSON)**
    that the LLM (Large Language Model) can interpret.

**2. Backend API Service (Skillset Server)**

-   Hosts the API endpoints defined in the skillset.

-   Each API should be in a structured format to the Copilot request.
    (api response + context that we get back)

**3. Natural Language Processing (How Copilot Interprets User Input)**

-   **Copilot itself** handles **natural language understanding**.

-   When a developer makes a request in plain language, Copilot:

    1.  **Breaks it down** into an intent (what the user is asking for).

    2.  **Matches it to a skillset endpoint** (if applicable).

    3.  **Calls the appropriate API endpoint** with structured
        parameters.

    4.  **Uses the API response** to generate a natural language
        completion.

**Is There a Dedicated NLP Endpoint?**

-   No **separate** NLP endpoint is needed.

-   Copilot **directly interprets** user input and maps it to available
    skillset endpoints.

-   The skillset endpoints are **passive data providers**---Copilot does
    the language processing itself.

**4. Response Formatting (Ensuring Copilot Can Use API Data)**

-   **API responses can be structured (for ex. JSON)** to be useful.

-   The LLM reads this response (JSON) and converts it into **natural
    language responses**.

**Example JSON Response from /random-user Endpoint (mermaid chart
candidate/sequence diagram):**

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "role": "Software Engineer"
}

```

-   Copilot uses this structured response to **generate a relevant
    output**, like:

    -   \"Here's a test user: John Doe (Software Engineer) -
        john.doe@example.com\"

**\
Summary of How It Works Together (combine with above #4)**

1.  **User Types a Request in Copilot**

    -   Example: *\"Generate a fake user for testing.\"*

2.  **Copilot Parses the Request**

    -   Recognizes it maps to /random-user endpoint.

3.  **API Call to Skillset Endpoint**

    -   Requests a new user profile.

4.  **Skillset API Responds with Structured Data**

    -   JSON object with user details.

5.  **Copilot Uses the Response to Generate Output**

    -   \"Here's a test user: John Doe, Software Engineer, email:
        john.doe@example.com.\"

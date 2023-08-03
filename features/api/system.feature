Feature: System
  Scenario: Make a request to "/livez" get health status
    When I make a GET request to "/livez"
    Then the response status code should be 200
    And the response body should match JSON:
      """
      {
        "ok": true
      }
      """

  Scenario: Make a request to "/readz" get health status
    When I make a GET request to "/readz"
    Then the response status code should be 200
    And the response body should match JSON:
      """
      {
        "ok": true
      }
      """

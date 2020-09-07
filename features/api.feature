Feature: get speed information
  In order to know what is the max speed allowed for monitored vehicles
  As an API user
  I need to be able to request max speed allowed

  Scenario: does not allow POST method
    When I send "POST" request to "/max-speed-allowed"
    Then the response code should be 405
    And the response should match json:
      """
      {
        "error": "Method not allowed"
      }
      """

  Scenario: should get max speed allowed
    When I send "GET" request to "/max-speed-allowed"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "speed": "90.0"
      }
      """

  Scenario: should get last speed record
    When I send "GET" request to "/{id}/last-speed"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "vehicle": "100"
        "speed": "85.0"
      }
      """

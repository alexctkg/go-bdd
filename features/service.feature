Feature: detect abrupt braking
  In order to identify events of abrupt braking
  As a event detector service
  I need to be able to analyse data points windows

  Scenario: Insufficient data points
    Given There are data points:
      | time       | imei            | speed | brakePedalStatus |
      | 1592218800 | 123456789000321 | 78    | 1                |

    When I process points
    Then the response message should be 'Insufficient data points'


  Scenario: Short Time Window
    Given There are data points:
      | time       | imei            | speed | brakePedalStatus |
      | 1592218800 | 123456789000321 | 78    | 1                |
      | 1592218920 | 123456789000321 | 80    | 1                |

    When I process points
    Then the response should contain no event


  Scenario: Normal Brake Occured
    Given There are data points:
      | time       | imei            | speed | brakePedalStatus |
      | 1592218800 | 123456789000321 | 78    | 1                |
      | 1592218815 | 123456789000321 | 70    | 1                |

    When I process points
    Then the response should contain 1 brake event


  Scenario: Abrupt Brake Occured
    Given There are data points:
      | time       | imei            | speed | brakePedalStatus |
      | 1592218800 | 123456789000321 | 85    | 1                |
      | 1592218815 | 123456789000321 | 78    | 1                |

    When I process points
    Then the response should contain 1 abrupt brake event

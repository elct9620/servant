@docker
Feature: Installer
  Scenario: Install servant controller to docker swarm
    Given a docker client
    When I run the installer
    Then I can see the following networks:
      | Name    |
      | servant |
    And I can see the follwoing services:
      | Name               |
      | servant-controller |

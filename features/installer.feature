@docker
Feature: Installer
  Scenario: Install servant controller to docker swarm
    When I run "docker servant install"
    Then the network should contains "servant"
    And the service should contains "servant-controller"

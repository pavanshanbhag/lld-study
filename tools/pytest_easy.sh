#!/usr/bin/env bash
# Run pytest for one 01-easy Python topic with isolated import path.
run_topic_pytest() {
  local topic="$1"
  local pydir="01-easy/${topic}/python"
  if [[ -d "${pydir}/tests" ]]; then
    PYTHONPATH="${pydir}" uv run pytest "${pydir}/tests" -q
  fi
}

run_all_easy_pytests() {
  run_topic_pytest parking-lot
  run_topic_pytest vending-machine
  run_topic_pytest logging-framework
  run_topic_pytest stack-overflow
  run_topic_pytest traffic-signal
  run_topic_pytest coffee-vending-machine
  run_topic_pytest task-management-system
}

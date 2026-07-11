#!/usr/bin/env bash
# Run pytest for one topic with an isolated PYTHONPATH (avoids cross-topic import clashes).

run_topic_pytest() {
  local pydir="$1"
  if [[ -d "${pydir}/tests" ]]; then
    PYTHONPATH="${pydir}" uv run pytest "${pydir}/tests" -q
  fi
}

run_all_easy_pytests() {
  run_topic_pytest "01-easy/parking-lot/python"
  run_topic_pytest "01-easy/vending-machine/python"
  run_topic_pytest "01-easy/logging-framework/python"
  run_topic_pytest "01-easy/stack-overflow/python"
  run_topic_pytest "01-easy/traffic-signal/python"
  run_topic_pytest "01-easy/coffee-vending-machine/python"
  run_topic_pytest "01-easy/task-management-system/python"
}

run_all_medium_pytests() {
  run_topic_pytest "02-medium/lru-cache/python"
  run_topic_pytest "02-medium/tic-tac-toe/python"
  run_topic_pytest "02-medium/pub-sub-system/python"
  run_topic_pytest "02-medium/atm/python"
  run_topic_pytest "02-medium/library-management-system/python"
  run_topic_pytest "02-medium/airline-management-system/python"
  run_topic_pytest "02-medium/car-rental-system/python"
  run_topic_pytest "02-medium/concert-ticket-booking-system/python"
  run_topic_pytest "02-medium/digital-wallet-service/python"
  run_topic_pytest "02-medium/elevator-system/python"
  run_topic_pytest "02-medium/hotel-management-system/python"
  run_topic_pytest "02-medium/linkedin/python"
  run_topic_pytest "02-medium/online-auction-system/python"
  run_topic_pytest "02-medium/restaurant-management-system/python"
  run_topic_pytest "02-medium/social-networking-service/python"
}

run_all_hard_pytests() {
  run_topic_pytest "03-hard/chess-game/python"
  run_topic_pytest "03-hard/course-registration-system/python"
  run_topic_pytest "03-hard/cricinfo/python"
  run_topic_pytest "03-hard/food-delivery-service/python"
  run_topic_pytest "03-hard/movie-ticket-booking-system/python"
  run_topic_pytest "03-hard/music-streaming-service/python"
  run_topic_pytest "03-hard/online-shopping-service/python"
  run_topic_pytest "03-hard/online-stock-brokerage-system/python"
  run_topic_pytest "03-hard/ride-sharing-service/python"
  run_topic_pytest "03-hard/snake-and-ladder/python"
  run_topic_pytest "03-hard/splitwise/python"
}

run_all_pytests() {
  run_all_easy_pytests
  run_all_medium_pytests
  run_all_hard_pytests
}

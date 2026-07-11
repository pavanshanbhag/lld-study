#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"
# shellcheck source=pytest_topics.sh
source "$ROOT/tools/pytest_topics.sh"

EASY_GO=(
  ./01-easy/parking-lot/golang/
  ./01-easy/vending-machine/golang/
  ./01-easy/logging-framework/golang/
  ./01-easy/stack-overflow/golang/
  ./01-easy/traffic-signal/golang/
  ./01-easy/coffee-vending-machine/golang/
  ./01-easy/task-management-system/golang/
)

MEDIUM_GO=(
  ./02-medium/lru-cache/golang/
  ./02-medium/tic-tac-toe/golang/
  ./02-medium/pub-sub-system/golang/
  ./02-medium/atm/golang/
  ./02-medium/library-management-system/golang/
  ./02-medium/airline-management-system/golang/
  ./02-medium/car-rental-system/golang/
  ./02-medium/concert-ticket-booking-system/golang/
  ./02-medium/digital-wallet-service/golang/
  ./02-medium/elevator-system/golang/
  ./02-medium/hotel-management-system/golang/
  ./02-medium/linkedin/golang/
  ./02-medium/online-auction-system/golang/
  ./02-medium/restaurant-management-system/golang/
  ./02-medium/social-networking-service/golang/
)

HARD_GO=(
  ./03-hard/chess-game/golang/
  ./03-hard/course-registration-system/golang/
  ./03-hard/cricinfo/golang/
  ./03-hard/food-delivery-service/golang/
  ./03-hard/movie-ticket-booking-system/golang/
  ./03-hard/music-streaming-service/golang/
  ./03-hard/online-shopping-service/golang/
  ./03-hard/online-stock-brokerage-system/golang/
  ./03-hard/ride-sharing-service/golang/
  ./03-hard/snake-and-ladder/golang/
  ./03-hard/splitwise/golang/
)

PATTERN_DEMOS=(
  "00-foundations/design-patterns/creational/singleton/golang"
  "00-foundations/design-patterns/creational/factory/golang"
  "00-foundations/design-patterns/creational/builder/golang"
  "00-foundations/design-patterns/creational/prototype/golang"
  "00-foundations/design-patterns/structural/adapter/golang"
  "00-foundations/design-patterns/structural/bridge/golang"
  "00-foundations/design-patterns/structural/composite/golang"
  "00-foundations/design-patterns/structural/decorator/golang"
  "00-foundations/design-patterns/structural/facade/golang"
  "00-foundations/design-patterns/structural/flyweight/golang"
  "00-foundations/design-patterns/structural/proxy/golang"
  "00-foundations/design-patterns/behavioral/chainofresponsibility/golang"
  "00-foundations/design-patterns/behavioral/iterator/golang"
  "00-foundations/design-patterns/behavioral/mediator/golang"
  "00-foundations/design-patterns/behavioral/memento/golang"
  "00-foundations/design-patterns/behavioral/observer/golang"
  "00-foundations/design-patterns/behavioral/state/golang"
  "00-foundations/design-patterns/behavioral/strategy/golang"
  "00-foundations/design-patterns/behavioral/templatemethod/golang"
)

run_go_tests() {
  local -a pkgs=("$@")
  for pkg in "${pkgs[@]}"; do
    go test -race "$pkg"
  done
}

run_pattern_demos() {
  for dir in "${PATTERN_DEMOS[@]}"; do
    (cd "$dir" && go run ./cmd/demo/)
  done
}

case "${1:-all}" in
  sync)
    uv sync
    ;;
  parking)
    go test -race ./01-easy/parking-lot/golang/
    uv run ruff check 01-easy/parking-lot/python
    uv run mypy 01-easy/parking-lot/python
    PYTHONPATH=01-easy/parking-lot/python uv run pytest 01-easy/parking-lot/python/tests -q
    ;;
  easy)
    run_go_tests "${EASY_GO[@]}"
    run_all_easy_pytests
    ;;
  easy-py)
    run_all_easy_pytests
    ;;
  medium)
    run_go_tests "${MEDIUM_GO[@]}"
    run_all_medium_pytests
    ;;
  medium-py)
    run_all_medium_pytests
    ;;
  hard)
    run_go_tests "${HARD_GO[@]}"
    run_all_hard_pytests
    ;;
  hard-py)
    run_all_hard_pytests
    ;;
  lru)
    go test -race ./02-medium/lru-cache/golang/
    run_topic_pytest "02-medium/lru-cache/python"
    ;;
  patterns)
    run_pattern_demos
    ;;
  test-go)
    run_go_tests "${EASY_GO[@]}" "${MEDIUM_GO[@]}" "${HARD_GO[@]}"
    ;;
  test-py)
    run_all_pytests
    ;;
  lint-py)
    uv run ruff check 01-easy/parking-lot/python
    uv run mypy 01-easy/parking-lot/python
    ;;
  test)
    run_go_tests "${EASY_GO[@]}" "${MEDIUM_GO[@]}" "${HARD_GO[@]}"
    run_all_pytests
    ;;
  all)
    run_go_tests "${EASY_GO[@]}" "${MEDIUM_GO[@]}" "${HARD_GO[@]}"
    uv run ruff check 01-easy/parking-lot/python
    uv run mypy 01-easy/parking-lot/python
    run_all_pytests
    run_pattern_demos
    go build ./...
    ;;
  *)
    echo "usage: $0 {sync|parking|easy|easy-py|medium|medium-py|hard|hard-py|lru|patterns|test-go|test-py|lint-py|test|all}"
    exit 1
    ;;
esac

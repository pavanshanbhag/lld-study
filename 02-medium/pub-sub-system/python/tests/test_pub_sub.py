from message import Message
from pub_sub_service import PubSubService
from subscriber import NewsSubscriber, Subscriber


class RecordingSubscriber(Subscriber):
    def __init__(self, subscriber_id: str) -> None:
        self.id = subscriber_id
        self.messages: list[Message] = []

    def get_id(self) -> str:
        return self.id

    def on_message(self, message: Message) -> None:
        self.messages.append(message)


def test_publish_delivers_to_subscriber() -> None:
    service = PubSubService()
    service.create_topic("SPORTS")
    fan = RecordingSubscriber("Fan1")
    service.subscribe("SPORTS", fan)
    service.publish("SPORTS", Message("Team wins!"))
    service.shutdown()  # waits for async delivery

    assert len(fan.messages) == 1
    assert fan.messages[0].get_payload() == "Team wins!"


def test_unsubscribe_stops_delivery() -> None:
    service = PubSubService()
    service.create_topic("NEWS")
    fan = RecordingSubscriber("Fan2")
    service.subscribe("NEWS", fan)
    service.unsubscribe("NEWS", fan)
    service.publish("NEWS", Message("ignored"))
    service.shutdown()
    assert fan.messages == []


def test_news_subscriber_constructs() -> None:
    sub = NewsSubscriber("n1")
    assert sub.get_id() == "n1"

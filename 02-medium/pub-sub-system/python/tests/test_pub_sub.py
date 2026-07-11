from message import Message
from pub_sub_service import PubSubService
from subscriber import NewsSubscriber


def test_publish_delivers_to_subscriber() -> None:
    service = PubSubService()
    service.create_topic("SPORTS")
    fan = NewsSubscriber("Fan1")
    service.subscribe("SPORTS", fan)
    service.publish("SPORTS", Message("Team wins!"))
    service.shutdown()

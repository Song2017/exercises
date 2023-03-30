from kafka import KafkaProducer
import json

_KAFKA_BROKER_URL = "shared-bootstrap.infra.samarkand.io:9094"

producer = KafkaProducer(
    bootstrap_servers=_KAFKA_BROKER_URL,
    value_serializer=lambda value: json.dumps(value).encode(),
    acks=1,
)

result = producer.send(
    "PILOT_NEW_WAYBILL_TOPIC", {
        "carrier": "samarkand.haiku.prod",
        "order_ref": "SO3205212",
        "pay_id": "2303281751500001270960",
        "payment_pay_id": "2303281751570005390797",
        "seller_order_ref": "E202303281751480397000851210HZ",
        "tracking_reference": "194957496398",
    }, b"194957496398")
print(result)

import pdb

pdb.set_trace()
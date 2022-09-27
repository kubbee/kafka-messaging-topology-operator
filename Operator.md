> operator-sdk init --domain kubbee.tech --repo github.com/kubbee/kafka-messaging-topology-operator

> go mod tidy

> operator-sdk create api --group messages --version v1beta1 --kind KafkaConsumer --resource --controller

> operator-sdk create api --group messages --version v1beta1 --kind KafkaProducer --resource --controller

> operator-sdk create api --group messages --version v1beta1 --kind KafkaACL --resource --controller

> operator-sdk create api --group messages --version v1beta1 --kind KafkaReference --resource --controller
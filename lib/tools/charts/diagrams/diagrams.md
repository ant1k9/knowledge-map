[diagrams](https://github.com/mingrammer/diagrams)

*Description*: Diagram as Code for prototyping cloud system architectures

*Labels*: #Python #Diagrams

*Docs*: https://diagrams.mingrammer.com/docs/getting-started/installation

*Examples*:

```python
from diagrams import Cluster, Diagram
from diagrams.aws.compute import ECS
from diagrams.aws.database import RDS
from diagrams.aws.network import Route53

with Diagram("Simple Web Service with DB Cluster", show=False):
    dns = Route53("dns")
    web = ECS("service")

    with Cluster("DB Cluster"):
        db_main = RDS("main")
        db_main - [RDS("replica1"),
                     RDS("replica2")]

    dns >> web >> db_main
```

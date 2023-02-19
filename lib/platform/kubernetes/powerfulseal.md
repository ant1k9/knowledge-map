[powerfulseal](https://github.com/powerfulseal/powerfulseal)

*Description*: A powerful testing tool for Kubernetes clusters.

*Labels*: #Python #Kubernetes #ChaosEngineering

*Links*:
  - https://www.bloomberg.com/company/stories/powerfulseal-testing-tool-kubernetes-clusters/

*Examples*:

```bash
$ echo 'scenarios:
- name: Kill one pod in my namespace, make sure the service responds
  steps:
  # kill a pod from myapp namespace
  - podAction:
      matches:
        - namespace: myapp
      filters:
        - randomSample:
            size: 1
      actions:
        - kill:
            probability: 0.75
  # check my service continues working
  - probeHTTP:
      target:
        service:
          name: my-service
          namespace: myapp
      endpoint: /healthz' > policy.yaml
$ powerfulseal autonomous --policy-file ./policy.yaml
```

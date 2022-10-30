[pyrra](https://github.com/pyrra-dev/pyrra)

*Description*: Making SLOs with Prometheus manageable, accessible, and easy to use for everyone!

*Labels*: #Go #Monitoring #SLO #Prometheus

*Links*:
  - https://demo.pyrra.dev

*Examples*:

```yaml
apiVersion: pyrra.dev/v1alpha1
kind: ServiceLevelObjective
metadata:
  name: pyrra-api-errors
  namespace: monitoring
  labels:
    prometheus: k8s
    role: alert-rules
    pyrra.dev/team: operations # Any labels prefixed with 'pyrra.dev/' will be propagated as Prometheus labels, while stripping the prefix.
spec:
  target: "99"
  window: 2w
  description: Pyrra's API requests and response errors over time grouped by route.
  indicator:
    ratio:
      errors:
        metric: http_requests_total{job="pyrra",code=~"5.."}
      total:
        metric: http_requests_total{job="pyrra"}
      grouping:
        - route 
```

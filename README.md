2025/02/27 00:19:51 deployElasticsearch called
2025/02/27 00:19:51 deploying elasticsearch-operator
2025/02/27 00:19:52 <---------------------------------------------------------->
2025/02/27 00:19:52 NAME: elastic-operator
2025/02/27 00:19:52 LAST DEPLOYED: 2025-02-26 22:55:23.088465 +0530 IST
2025/02/27 00:19:52 NAMESPACE: elastic-system
2025/02/27 00:19:52 Chart: eck-operator
2025/02/27 00:19:52 Chart VERSION: 2.16.1
2025/02/27 00:19:52 Home: https://github.com/elastic/cloud-on-k8s
2025/02/27 00:19:52 STATUS: deployed
2025/02/27 00:19:52 REVISION: 12
2025/02/27 00:19:52 <---------------------------------------------------------->
2025/02/27 00:19:55 preparing upgrade for elastic-operator
2025/02/27 00:19:56 performing update for elastic-operator
2025/02/27 00:19:56 creating upgraded release for elastic-operator
2025/02/27 00:19:56 checking 20 resources for changes
2025/02/27 00:19:56 Looks like there are no changes for ServiceAccount "elastic-operator"
2025/02/27 00:19:57 Looks like there are no changes for Secret "elastic-operator-webhook-cert"
2025/02/27 00:19:57 Looks like there are no changes for ConfigMap "elastic-operator"
2025/02/27 00:19:57 Looks like there are no changes for CustomResourceDefinition "agents.agent.k8s.elastic.co"
2025/02/27 00:19:58 Looks like there are no changes for CustomResourceDefinition "apmservers.apm.k8s.elastic.co"
2025/02/27 00:19:58 Looks like there are no changes for CustomResourceDefinition "beats.beat.k8s.elastic.co"
2025/02/27 00:19:58 Looks like there are no changes for CustomResourceDefinition "elasticmapsservers.maps.k8s.elastic.co"
2025/02/27 00:19:59 Looks like there are no changes for CustomResourceDefinition "elasticsearchautoscalers.autoscaling.k8s.elastic.co"
2025/02/27 00:20:03 Looks like there are no changes for CustomResourceDefinition "elasticsearches.elasticsearch.k8s.elastic.co"
2025/02/27 00:20:04 Looks like there are no changes for CustomResourceDefinition "enterprisesearches.enterprisesearch.k8s.elastic.co"
2025/02/27 00:20:04 Looks like there are no changes for CustomResourceDefinition "kibanas.kibana.k8s.elastic.co"
2025/02/27 00:20:05 Looks like there are no changes for CustomResourceDefinition "logstashes.logstash.k8s.elastic.co"
2025/02/27 00:20:05 Looks like there are no changes for CustomResourceDefinition "stackconfigpolicies.stackconfigpolicy.k8s.elastic.co"
2025/02/27 00:20:05 Looks like there are no changes for ClusterRole "elastic-operator"
2025/02/27 00:20:06 Looks like there are no changes for ClusterRole "elastic-operator-view"
2025/02/27 00:20:06 Looks like there are no changes for ClusterRole "elastic-operator-edit"
2025/02/27 00:20:06 Looks like there are no changes for ClusterRoleBinding "elastic-operator"
2025/02/27 00:20:06 Looks like there are no changes for Service "elastic-operator-webhook"
2025/02/27 00:20:06 Patch StatefulSet "elastic-operator" in namespace elastic-system
2025/02/27 00:20:07 Patch ValidatingWebhookConfiguration "elastic-operator.elastic-system.k8s.elastic.co" in namespace 
2025/02/27 00:20:07 waiting for release elastic-operator resources (created: 0 updated: 20  deleted: 0)
2025/02/27 00:20:07 beginning wait for 20 resources with timeout of 10m0s
2025/02/27 00:20:09 StatefulSet is ready: elastic-system/elastic-operator. 1 out of 1 expected pods are ready
2025/02/27 00:20:09 updating status for upgraded release for elastic-operator
2025/02/27 00:20:10 <---------------------------------------------------------->
2025/02/27 00:20:10 NAME: elastic-operator
2025/02/27 00:20:10 LAST DEPLOYED: 2025-02-27 00:19:55.917717 +0530 IST m=+4.734903001
2025/02/27 00:20:10 NAMESPACE: elastic-system
2025/02/27 00:20:10 Chart: eck-operator
2025/02/27 00:20:10 Chart VERSION: 2.16.1
2025/02/27 00:20:10 Home: https://github.com/elastic/cloud-on-k8s
2025/02/27 00:20:10 STATUS: deployed
2025/02/27 00:20:10 REVISION: 13
2025/02/27 00:20:10 <---------------------------------------------------------->
2025/02/27 00:20:10 deploying elasticsearch
2025/02/27 00:20:10 Applied elasticsearch yaml
2025/02/27 00:20:10 Fetching elastic user password
2025/02/27 00:20:10 Waiting for secret to be available
2025/02/27 00:20:10 Secret "es-cluster-es-elastic-user" in namespace "default" is now present!
2025/02/27 00:20:10 Elasticsearch URL:  https://elastic:I71GHxg0W3aj154G34rTd7GV@es-cluster-es-http:9200
2025/02/27 00:20:10 Deploying Prometheus ElasticSearch Exporter
2025/02/27 00:20:10 Reading rules file for alertmanager rules
2025/02/27 00:20:10 <---------------------------------------------------------->
2025/02/27 00:20:10 NAME: elastic-exporter
2025/02/27 00:20:10 LAST DEPLOYED: 2025-02-26 22:55:35.283065 +0530 IST
2025/02/27 00:20:10 NAMESPACE: default
2025/02/27 00:20:10 Chart: prometheus-elasticsearch-exporter
2025/02/27 00:20:10 Chart VERSION: 6.6.1
2025/02/27 00:20:10 Home: https://github.com/prometheus-community/elasticsearch_exporter
2025/02/27 00:20:10 STATUS: deployed
2025/02/27 00:20:10 REVISION: 11
2025/02/27 00:20:10 <---------------------------------------------------------->
2025/02/27 00:20:15 preparing upgrade for elastic-exporter
2025/02/27 00:20:16 performing update for elastic-exporter
2025/02/27 00:20:16 creating upgraded release for elastic-exporter
2025/02/27 00:20:16 checking 3 resources for changes
2025/02/27 00:20:16 Looks like there are no changes for Service "elastic-exporter-prometheus-elasticsearch-exporter"
2025/02/27 00:20:16 Patch Deployment "elastic-exporter-prometheus-elasticsearch-exporter" in namespace default
2025/02/27 00:20:16 Patch ServiceMonitor "elastic-exporter-prometheus-elasticsearch-exporter" in namespace default
2025/02/27 00:20:16 waiting for release elastic-exporter resources (created: 0 updated: 3  deleted: 0)
2025/02/27 00:20:16 beginning wait for 3 resources with timeout of 10m0s
2025/02/27 00:20:17 updating status for upgraded release for elastic-exporter
2025/02/27 00:20:17 <---------------------------------------------------------->
2025/02/27 00:20:17 NAME: elastic-exporter
2025/02/27 00:20:17 LAST DEPLOYED: 2025-02-27 00:20:15.836231 +0530 IST m=+24.652970043
2025/02/27 00:20:17 NAMESPACE: default
2025/02/27 00:20:17 Chart: prometheus-elasticsearch-exporter
2025/02/27 00:20:17 Chart VERSION: 6.6.1
2025/02/27 00:20:17 Home: https://github.com/prometheus-community/elasticsearch_exporter
2025/02/27 00:20:17 STATUS: deployed
2025/02/27 00:20:17 REVISION: 12
2025/02/27 00:20:17 <---------------------------------------------------------->
2025/02/27 00:20:17 Deployed ElasticSearch Exporter
2025/02/27 00:20:17 Deploying Alertmanager rules
2025/02/27 00:20:17 Applied Alertmanager rules
2025/02/27 00:20:17 Deploying Dashboard
2025/02/27 00:20:17 Applied Dashboard

# Elasticsearch monitoring Setup

## Building the project

The project is orchaestrated in golang, to build the project do
```sh
go build
```

## Prequisite
Kubeconfig should be Setup

```sh
➜  elasticsearch-monitoring-stack git:(main) k get nodes
NAME                   STATUS   ROLES    AGE     VERSION
pool-lhftzph7u-a5lx3   Ready    <none>   4h33m   v1.32.1
pool-lhftzph7u-a5lx8   Ready    <none>   4h33m   v1.32.1
pool-lhftzph7u-a5lxu   Ready    <none>   4h33m   v1.32.1
```

## To deploy kubeprometheus stack

After the building, project you can run 
```sh
➜  elasticsearch-monitoring-stack git:(main) ./elasticsearch-monitoring-stack deployKubeProm
deployKubeProm called
2025/02/27 00:16:44 <---------------------------------------------------------->
2025/02/27 00:16:44 NAME: prom-stack
2025/02/27 00:16:44 LAST DEPLOYED: 2025-02-26 19:42:24.498245 +0530 IST
2025/02/27 00:16:44 NAMESPACE: monitoring
2025/02/27 00:16:44 Chart: kube-prometheus-stack
2025/02/27 00:16:44 Chart VERSION: 69.5.2
2025/02/27 00:16:44 Home: https://github.com/prometheus-operator/kube-prometheus
2025/02/27 00:16:44 STATUS: deployed
2025/02/27 00:16:44 REVISION: 1
2025/02/27 00:16:44 <---------------------------------------------------------->
2025/02/27 00:16:48 preparing upgrade for prom-stack
2025/02/27 00:16:51 performing update for prom-stack
2025/02/27 00:16:53 creating upgraded release for prom-stack
2025/02/27 00:16:53 Starting delete for "prom-stack-kube-prometheus-admission" ServiceAccount
2025/02/27 00:16:53 Ignoring delete failure for "prom-stack-kube-prometheus-admission" /v1, Kind=ServiceAccount: serviceaccounts "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:16:53 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:16:54 creating 1 resource(s)
2025/02/27 00:16:54 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRole
2025/02/27 00:16:54 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=ClusterRole: clusterroles.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:16:54 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:16:55 creating 1 resource(s)
2025/02/27 00:16:55 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRoleBinding
2025/02/27 00:16:55 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=ClusterRoleBinding: clusterrolebindings.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:16:55 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:16:55 creating 1 resource(s)
2025/02/27 00:16:55 Starting delete for "prom-stack-kube-prometheus-admission" Role
2025/02/27 00:16:55 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=Role: roles.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:16:55 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:16:56 creating 1 resource(s)
2025/02/27 00:16:56 Starting delete for "prom-stack-kube-prometheus-admission" RoleBinding
2025/02/27 00:16:56 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=RoleBinding: rolebindings.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:16:56 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:16:57 creating 1 resource(s)
2025/02/27 00:16:57 Starting delete for "prom-stack-kube-prometheus-admission-create" Job
2025/02/27 00:16:57 Ignoring delete failure for "prom-stack-kube-prometheus-admission-create" batch/v1, Kind=Job: jobs.batch "prom-stack-kube-prometheus-admission-create" not found
2025/02/27 00:16:57 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:16:57 creating 1 resource(s)
2025/02/27 00:16:57 Watching for changes to Job prom-stack-kube-prometheus-admission-create with timeout of 10m0s
2025/02/27 00:16:57 Add/Modify event for prom-stack-kube-prometheus-admission-create: ADDED
2025/02/27 00:16:57 prom-stack-kube-prometheus-admission-create: Jobs active: 1, jobs failed: 0, jobs succeeded: 0
2025/02/27 00:17:00 Add/Modify event for prom-stack-kube-prometheus-admission-create: MODIFIED
2025/02/27 00:17:00 prom-stack-kube-prometheus-admission-create: Jobs active: 0, jobs failed: 0, jobs succeeded: 0
2025/02/27 00:17:01 Add/Modify event for prom-stack-kube-prometheus-admission-create: MODIFIED
2025/02/27 00:17:01 Starting delete for "prom-stack-kube-prometheus-admission-create" Job
2025/02/27 00:17:01 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:01 Starting delete for "prom-stack-kube-prometheus-admission" RoleBinding
2025/02/27 00:17:01 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:01 Starting delete for "prom-stack-kube-prometheus-admission" Role
2025/02/27 00:17:01 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:01 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRoleBinding
2025/02/27 00:17:02 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:02 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRole
2025/02/27 00:17:02 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:02 Starting delete for "prom-stack-kube-prometheus-admission" ServiceAccount
2025/02/27 00:17:02 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:02 checking 116 resources for changes
2025/02/27 00:17:02 Looks like there are no changes for ServiceAccount "prom-stack-grafana"
2025/02/27 00:17:02 Looks like there are no changes for ServiceAccount "prom-stack-kube-state-metrics"
2025/02/27 00:17:03 Looks like there are no changes for ServiceAccount "prom-stack-prometheus-node-exporter"
2025/02/27 00:17:03 Looks like there are no changes for ServiceAccount "prom-stack-kube-prometheus-alertmanager"
2025/02/27 00:17:03 Looks like there are no changes for ServiceAccount "prom-stack-kube-prometheus-operator"
2025/02/27 00:17:03 Looks like there are no changes for ServiceAccount "prom-stack-kube-prometheus-prometheus"
2025/02/27 00:17:03 Looks like there are no changes for Secret "prom-stack-grafana"
2025/02/27 00:17:04 Looks like there are no changes for Secret "alertmanager-prom-stack-kube-prometheus-alertmanager"
2025/02/27 00:17:04 Looks like there are no changes for ConfigMap "prom-stack-grafana-config-dashboards"
2025/02/27 00:17:04 Looks like there are no changes for ConfigMap "prom-stack-grafana"
2025/02/27 00:17:04 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-grafana-datasource"
2025/02/27 00:17:05 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-alertmanager-overview"
2025/02/27 00:17:05 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-apiserver"
2025/02/27 00:17:05 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-cluster-total"
2025/02/27 00:17:05 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-controller-manager"
2025/02/27 00:17:06 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-etcd"
2025/02/27 00:17:06 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-grafana-overview"
2025/02/27 00:17:06 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-coredns"
2025/02/27 00:17:06 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-resources-cluster"
2025/02/27 00:17:06 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-resources-multicluster"
2025/02/27 00:17:07 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-resources-namespace"
2025/02/27 00:17:07 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-resources-node"
2025/02/27 00:17:07 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-resources-pod"
2025/02/27 00:17:08 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-resources-workload"
2025/02/27 00:17:08 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-k8s-resources-workloads-namespace"
2025/02/27 00:17:08 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-kubelet"
2025/02/27 00:17:08 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-namespace-by-pod"
2025/02/27 00:17:09 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-namespace-by-workload"
2025/02/27 00:17:09 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-node-cluster-rsrc-use"
2025/02/27 00:17:09 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-node-rsrc-use"
2025/02/27 00:17:10 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-nodes-aix"
2025/02/27 00:17:11 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-nodes-darwin"
2025/02/27 00:17:11 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-nodes"
2025/02/27 00:17:12 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-persistentvolumesusage"
2025/02/27 00:17:12 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-pod-total"
2025/02/27 00:17:13 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-prometheus"
2025/02/27 00:17:13 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-proxy"
2025/02/27 00:17:13 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-scheduler"
2025/02/27 00:17:13 Looks like there are no changes for ConfigMap "prom-stack-kube-prometheus-workload-total"
2025/02/27 00:17:14 Looks like there are no changes for ClusterRole "prom-stack-grafana-clusterrole"
2025/02/27 00:17:14 Looks like there are no changes for ClusterRole "prom-stack-kube-state-metrics"
2025/02/27 00:17:14 Looks like there are no changes for ClusterRole "prom-stack-kube-prometheus-operator"
2025/02/27 00:17:14 Looks like there are no changes for ClusterRole "prom-stack-kube-prometheus-prometheus"
2025/02/27 00:17:14 Looks like there are no changes for ClusterRoleBinding "prom-stack-grafana-clusterrolebinding"
2025/02/27 00:17:15 Looks like there are no changes for ClusterRoleBinding "prom-stack-kube-state-metrics"
2025/02/27 00:17:15 Looks like there are no changes for ClusterRoleBinding "prom-stack-kube-prometheus-operator"
2025/02/27 00:17:15 Looks like there are no changes for ClusterRoleBinding "prom-stack-kube-prometheus-prometheus"
2025/02/27 00:17:16 Patch Role "prom-stack-grafana" in namespace monitoring
2025/02/27 00:17:16 Looks like there are no changes for RoleBinding "prom-stack-grafana"
2025/02/27 00:17:16 Looks like there are no changes for Service "prom-stack-grafana"
2025/02/27 00:17:16 Looks like there are no changes for Service "prom-stack-kube-state-metrics"
2025/02/27 00:17:16 Looks like there are no changes for Service "prom-stack-prometheus-node-exporter"
2025/02/27 00:17:17 Looks like there are no changes for Service "prom-stack-kube-prometheus-alertmanager"
2025/02/27 00:17:17 Looks like there are no changes for Service "prom-stack-kube-prometheus-coredns"
2025/02/27 00:17:17 Looks like there are no changes for Service "prom-stack-kube-prometheus-kube-controller-manager"
2025/02/27 00:17:17 Looks like there are no changes for Service "prom-stack-kube-prometheus-kube-etcd"
2025/02/27 00:17:18 Looks like there are no changes for Service "prom-stack-kube-prometheus-kube-proxy"
2025/02/27 00:17:18 Looks like there are no changes for Service "prom-stack-kube-prometheus-kube-scheduler"
2025/02/27 00:17:18 Looks like there are no changes for Service "prom-stack-kube-prometheus-operator"
2025/02/27 00:17:18 Patch Service "prom-stack-kube-prometheus-prometheus" in namespace monitoring
2025/02/27 00:17:18 Patch DaemonSet "prom-stack-prometheus-node-exporter" in namespace monitoring
2025/02/27 00:17:19 Patch Deployment "prom-stack-grafana" in namespace monitoring
2025/02/27 00:17:19 Patch Deployment "prom-stack-kube-state-metrics" in namespace monitoring
2025/02/27 00:17:19 Patch Deployment "prom-stack-kube-prometheus-operator" in namespace monitoring
2025/02/27 00:17:19 Patch Alertmanager "prom-stack-kube-prometheus-alertmanager" in namespace monitoring
2025/02/27 00:17:20 Patch MutatingWebhookConfiguration "prom-stack-kube-prometheus-admission" in namespace 
2025/02/27 00:17:20 Patch Prometheus "prom-stack-kube-prometheus-prometheus" in namespace monitoring
2025/02/27 00:17:20 Patch PrometheusRule "prom-stack-kube-prometheus-alertmanager.rules" in namespace monitoring
2025/02/27 00:17:20 Patch PrometheusRule "prom-stack-kube-prometheus-config-reloaders" in namespace monitoring
2025/02/27 00:17:21 Patch PrometheusRule "prom-stack-kube-prometheus-etcd" in namespace monitoring
2025/02/27 00:17:21 Patch PrometheusRule "prom-stack-kube-prometheus-general.rules" in namespace monitoring
2025/02/27 00:17:21 Patch PrometheusRule "prom-stack-kube-prometheus-k8s.rules.container-cpu-usage-second" in namespace monitoring
2025/02/27 00:17:21 Patch PrometheusRule "prom-stack-kube-prometheus-k8s.rules.container-memory-cache" in namespace monitoring
2025/02/27 00:17:22 Patch PrometheusRule "prom-stack-kube-prometheus-k8s.rules.container-memory-rss" in namespace monitoring
2025/02/27 00:17:22 Patch PrometheusRule "prom-stack-kube-prometheus-k8s.rules.container-memory-swap" in namespace monitoring
2025/02/27 00:17:22 Patch PrometheusRule "prom-stack-kube-prometheus-k8s.rules.container-memory-working-s" in namespace monitoring
2025/02/27 00:17:22 Patch PrometheusRule "prom-stack-kube-prometheus-k8s.rules.container-resource" in namespace monitoring
2025/02/27 00:17:23 Patch PrometheusRule "prom-stack-kube-prometheus-k8s.rules.pod-owner" in namespace monitoring
2025/02/27 00:17:23 Patch PrometheusRule "prom-stack-kube-prometheus-kube-apiserver-availability.rules" in namespace monitoring
2025/02/27 00:17:23 Patch PrometheusRule "prom-stack-kube-prometheus-kube-apiserver-burnrate.rules" in namespace monitoring
2025/02/27 00:17:23 Patch PrometheusRule "prom-stack-kube-prometheus-kube-apiserver-histogram.rules" in namespace monitoring
2025/02/27 00:17:24 Patch PrometheusRule "prom-stack-kube-prometheus-kube-apiserver-slos" in namespace monitoring
2025/02/27 00:17:24 Patch PrometheusRule "prom-stack-kube-prometheus-kube-prometheus-general.rules" in namespace monitoring
2025/02/27 00:17:24 Patch PrometheusRule "prom-stack-kube-prometheus-kube-prometheus-node-recording.rules" in namespace monitoring
2025/02/27 00:17:24 Patch PrometheusRule "prom-stack-kube-prometheus-kube-scheduler.rules" in namespace monitoring
2025/02/27 00:17:24 Patch PrometheusRule "prom-stack-kube-prometheus-kube-state-metrics" in namespace monitoring
2025/02/27 00:17:25 Patch PrometheusRule "prom-stack-kube-prometheus-kubelet.rules" in namespace monitoring
2025/02/27 00:17:25 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-apps" in namespace monitoring
2025/02/27 00:17:25 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-resources" in namespace monitoring
2025/02/27 00:17:25 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-storage" in namespace monitoring
2025/02/27 00:17:26 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-system-apiserver" in namespace monitoring
2025/02/27 00:17:26 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-system-controller-manager" in namespace monitoring
2025/02/27 00:17:27 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-system-kube-proxy" in namespace monitoring
2025/02/27 00:17:27 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-system-kubelet" in namespace monitoring
2025/02/27 00:17:27 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-system-scheduler" in namespace monitoring
2025/02/27 00:17:27 Patch PrometheusRule "prom-stack-kube-prometheus-kubernetes-system" in namespace monitoring
2025/02/27 00:17:28 Patch PrometheusRule "prom-stack-kube-prometheus-node-exporter.rules" in namespace monitoring
2025/02/27 00:17:28 Patch PrometheusRule "prom-stack-kube-prometheus-node-exporter" in namespace monitoring
2025/02/27 00:17:28 Patch PrometheusRule "prom-stack-kube-prometheus-node-network" in namespace monitoring
2025/02/27 00:17:28 Patch PrometheusRule "prom-stack-kube-prometheus-node.rules" in namespace monitoring
2025/02/27 00:17:28 Patch PrometheusRule "prom-stack-kube-prometheus-prometheus-operator" in namespace monitoring
2025/02/27 00:17:29 Patch PrometheusRule "prom-stack-kube-prometheus-prometheus" in namespace monitoring
2025/02/27 00:17:29 Patch ServiceMonitor "prom-stack-grafana" in namespace monitoring
2025/02/27 00:17:29 Patch ServiceMonitor "prom-stack-kube-state-metrics" in namespace monitoring
2025/02/27 00:17:29 Patch ServiceMonitor "prom-stack-prometheus-node-exporter" in namespace monitoring
2025/02/27 00:17:30 Patch ServiceMonitor "prom-stack-kube-prometheus-alertmanager" in namespace monitoring
2025/02/27 00:17:30 Patch ServiceMonitor "prom-stack-kube-prometheus-coredns" in namespace monitoring
2025/02/27 00:17:30 Patch ServiceMonitor "prom-stack-kube-prometheus-apiserver" in namespace monitoring
2025/02/27 00:17:30 Patch ServiceMonitor "prom-stack-kube-prometheus-kube-controller-manager" in namespace monitoring
2025/02/27 00:17:31 Patch ServiceMonitor "prom-stack-kube-prometheus-kube-etcd" in namespace monitoring
2025/02/27 00:17:31 Patch ServiceMonitor "prom-stack-kube-prometheus-kube-proxy" in namespace monitoring
2025/02/27 00:17:31 Patch ServiceMonitor "prom-stack-kube-prometheus-kube-scheduler" in namespace monitoring
2025/02/27 00:17:31 Patch ServiceMonitor "prom-stack-kube-prometheus-kubelet" in namespace monitoring
2025/02/27 00:17:31 Patch ServiceMonitor "prom-stack-kube-prometheus-operator" in namespace monitoring
2025/02/27 00:17:32 Patch ServiceMonitor "prom-stack-kube-prometheus-prometheus" in namespace monitoring
2025/02/27 00:17:32 Patch ValidatingWebhookConfiguration "prom-stack-kube-prometheus-admission" in namespace 
2025/02/27 00:17:32 waiting for release prom-stack resources (created: 0 updated: 116  deleted: 0)
2025/02/27 00:17:32 beginning wait for 116 resources with timeout of 10m0s
2025/02/27 00:17:33 Starting delete for "prom-stack-kube-prometheus-admission" ServiceAccount
2025/02/27 00:17:33 Ignoring delete failure for "prom-stack-kube-prometheus-admission" /v1, Kind=ServiceAccount: serviceaccounts "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:17:33 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:34 creating 1 resource(s)
2025/02/27 00:17:34 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRole
2025/02/27 00:17:35 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=ClusterRole: clusterroles.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:17:35 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:35 creating 1 resource(s)
2025/02/27 00:17:35 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRoleBinding
2025/02/27 00:17:35 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=ClusterRoleBinding: clusterrolebindings.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:17:35 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:36 creating 1 resource(s)
2025/02/27 00:17:36 Starting delete for "prom-stack-kube-prometheus-admission" Role
2025/02/27 00:17:36 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=Role: roles.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:17:36 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:37 creating 1 resource(s)
2025/02/27 00:17:37 Starting delete for "prom-stack-kube-prometheus-admission" RoleBinding
2025/02/27 00:17:37 Ignoring delete failure for "prom-stack-kube-prometheus-admission" rbac.authorization.k8s.io/v1, Kind=RoleBinding: rolebindings.rbac.authorization.k8s.io "prom-stack-kube-prometheus-admission" not found
2025/02/27 00:17:37 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:38 creating 1 resource(s)
2025/02/27 00:17:38 Starting delete for "prom-stack-kube-prometheus-admission-patch" Job
2025/02/27 00:17:38 Ignoring delete failure for "prom-stack-kube-prometheus-admission-patch" batch/v1, Kind=Job: jobs.batch "prom-stack-kube-prometheus-admission-patch" not found
2025/02/27 00:17:38 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:39 creating 1 resource(s)
2025/02/27 00:17:39 Watching for changes to Job prom-stack-kube-prometheus-admission-patch with timeout of 10m0s
2025/02/27 00:17:39 Add/Modify event for prom-stack-kube-prometheus-admission-patch: ADDED
2025/02/27 00:17:39 prom-stack-kube-prometheus-admission-patch: Jobs active: 1, jobs failed: 0, jobs succeeded: 0
2025/02/27 00:17:47 Add/Modify event for prom-stack-kube-prometheus-admission-patch: MODIFIED
2025/02/27 00:17:47 prom-stack-kube-prometheus-admission-patch: Jobs active: 0, jobs failed: 0, jobs succeeded: 0
2025/02/27 00:17:47 Add/Modify event for prom-stack-kube-prometheus-admission-patch: MODIFIED
2025/02/27 00:17:47 Starting delete for "prom-stack-kube-prometheus-admission-patch" Job
2025/02/27 00:17:47 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:48 Starting delete for "prom-stack-kube-prometheus-admission" RoleBinding
2025/02/27 00:17:48 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:48 Starting delete for "prom-stack-kube-prometheus-admission" Role
2025/02/27 00:17:48 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:48 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRoleBinding
2025/02/27 00:17:48 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:48 Starting delete for "prom-stack-kube-prometheus-admission" ClusterRole
2025/02/27 00:17:48 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:48 Starting delete for "prom-stack-kube-prometheus-admission" ServiceAccount
2025/02/27 00:17:48 beginning wait for 1 resources to be deleted with timeout of 10m0s
2025/02/27 00:17:49 updating status for upgraded release for prom-stack
2025/02/27 00:17:49 <---------------------------------------------------------->
2025/02/27 00:17:49 NAME: prom-stack
2025/02/27 00:17:49 LAST DEPLOYED: 2025-02-27 00:16:49.080578 +0530 IST m=+6.293736876
2025/02/27 00:17:49 NAMESPACE: monitoring
2025/02/27 00:17:49 Chart: kube-prometheus-stack
2025/02/27 00:17:49 Chart VERSION: 69.5.2
2025/02/27 00:17:49 Home: https://github.com/prometheus-operator/kube-prometheus
2025/02/27 00:17:49 STATUS: deployed
2025/02/27 00:17:49 REVISION: 2
2025/02/27 00:17:49 <---------------------------------------------------------->
```

## To deploy elasticsearch eck operator and elasticsearch

The default elasticsearch.yaml is given below

```yaml
apiVersion: elasticsearch.k8s.elastic.co/v1
kind: Elasticsearch
metadata:
  name: es-cluster
spec:
  version: 8.17.2
  nodeSets:
  - name: default
    count: 3
    config:
      # Tweak to avoid using mmap-based fields if desired
      node.store.allow_mmap: false
      # Roles for each node (master, data, ingest)
      node.roles: ["master", "data", "ingest"]
```

To deploy everything please do

```
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
```

This will deploy the alert manager rules for elasticsearch & deploy dashboard for elasticsearch.

kubectl delete -f manifests/mysql/mysql-deployment.yaml
kubectl delete -f manifests/mysql/mysql-clusterip-service.yaml 
kubectl delete -f manifests/mysql/mysql-pvc.yaml
kubectl delete -f manifests/mysql/mysql-storageclass.yaml
kubectl delete -f manifests/mysql/mysql-configmap.yaml
kubectl delete -f manifests/mysql/mysql-secret.yaml
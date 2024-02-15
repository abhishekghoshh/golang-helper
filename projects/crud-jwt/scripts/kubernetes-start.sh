kubectl apply -f manifests/mysql/mysql-configmap.yaml
kubectl apply -f manifests/mysql/mysql-secret.yaml
kubectl apply -f manifests/mysql/mysql-storageclass.yaml
kubectl apply -f manifests/mysql/mysql-pvc.yaml
kubectl apply -f manifests/mysql/mysql-deployment.yaml
kubectl apply -f manifests/mysql/mysql-clusterip-service.yaml 
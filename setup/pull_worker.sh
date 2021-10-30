# 查看需要的镜像:
# kubeadm config images list  
#
# root@master:~/k8s_scripts# kubeadm config images list
# k8s.gcr.io/kube-apiserver:v1.22.3
# k8s.gcr.io/kube-controller-manager:v1.22.3
# k8s.gcr.io/kube-scheduler:v1.22.3
# k8s.gcr.io/kube-proxy:v1.22.3
# k8s.gcr.io/pause:3.5
# k8s.gcr.io/etcd:3.5.0-0
# k8s.gcr.io/coredns/coredns:v1.8.4
# 

set -o errexit
set -o nounset
set -o pipefail

##这里定义版本，按照上面得到的列表自己改一下版本号

KUBE_VERSION=v1.22.3
KUBE_PAUSE_VERSION=3.5
ETCD_VERSION=3.5.0-0
DNS_VERSION=1.8.4

##这是原始仓库名，最后需要改名成这个
GCR_URL=k8s.gcr.io

##这里就是写你要使用的仓库
DOCKERHUB_URL=gotok8s

##这里是镜像列表，新版本要把coredns改成coredns/coredns
images=(
kube-proxy:${KUBE_VERSION}
pause:${KUBE_PAUSE_VERSION}
)

##这里是拉取和改名的循环语句
for imageName in ${images[@]} ; do
  docker pull $DOCKERHUB_URL/$imageName
  docker tag $DOCKERHUB_URL/$imageName $GCR_URL/$imageName
  docker rmi $DOCKERHUB_URL/$imageName
done

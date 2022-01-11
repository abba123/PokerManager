# PokerManager
PokerManagr is a software which can manage and analysis your hands played on GGPoker and Natural8

# Related Installation
* Frontend : vue
* Backend : golang(gin)
* Web server : Nginx
* DB : mysql
* Cache : Redis
* Message Queue : Kafka
* deployment : docker, k8s

# Start
you can use k8s to run the program

## for mac
start the minikube
    
    minikube start

create the pod (all for one)

    kubectl create -f pod.yaml
  
expose the port

    sudo kubectl port-forward my-pod 80:80 8000:8000
  
connect the web

    connect web with url http://127.0.0.1:80

# What we can do

Details in the [user guide](https://github.com/abba123/PokerManager/blob/master/guide.md)

# Help
connect me a0981861951@gmail.com

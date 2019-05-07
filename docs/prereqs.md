# Pre-requisites 

Please make sure the following have been completed. These steps include the 
following setups:
- Docker Desktop
- Homebrew
- go
- (optional) Github

### Install Docker Desktop

Docker Desktop includes Kubernetes. Use the appropriate link to download and install (links include installation instructions).

*macOS*

https://hub.docker.com/editions/community/docker-ce-desktop-mac

*windows*

https://hub.docker.com/editions/community/docker-ce-desktop-windows

Once Docker Desktop is installed, you need to explicitly enable Kubernetes support. Click the Docker icon in the status 
bar, go to “Preferences”, and on the “Kubernetes” tab check “Enable Kubernetes”

This will start kubernetes and install `kubectl`.  This might take a while, but the dialog will let you know once the 
Kubernetes cluster is ready.


![Kubernetes Running](screens/docker-desktop-k8s-running.png)
 
### Install Homebrew 
Homebrew is a package manager

macOS
```bash
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

Linux or Windows Subsystem for Linux (WSL)
```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/Linuxbrew/install/master/install.sh)"
```

### Github Account

If you do not already have a github account you may want to create one to push code for these guides.

https://github.com/join

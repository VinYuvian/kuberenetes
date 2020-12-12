# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do | config |
	config.vm.define "kmaster" do |node|
		node.vm.hostname = "kmaster"
		node.vm.network "private_network", ip: "192.168.10.10"
		node.vm.box = "bento/centos-7"
		node.vm.synced_folder "kube_yamls/","/home/vagrant/kube_yamls"
		node.vm.provider "virtualbox" do |v|
			v.name = "kmaster"
			v.memory = 2048
			v.cpus = 2
			v.customize ["modifyvm", :id, "--audio", "none"]
		end
		node.vm.provision "shell", path: "boot_strap.sh"
		node.vm.provision "shell", path: "boot_strap_kmaster.sh"
	end

	nodeCount = 2

	(1..nodeCount).each do |i|
		config.vm.define "kworker#{i}" do | node |
			node.vm.box = "bento/centos-7"
			node.vm.hostname = "kworker#{i}"
			node.vm.network "private_network", ip: "192.168.10.1#{i}"
			node.vm.provider "virtualbox" do |v|
     		   v.name = "kworker#{i}"
  	     	   v.memory = 1024
               v.cpus = 1
        # Prevent VirtualBox from interfering with host audio stack
               v.customize ["modifyvm", :id, "--audio", "none"]
            end
            node.vm.provision "shell", path: "boot_strap.sh"
            node.vm.provision "shell", path: "bootstrap_kworker.sh"
        end
     end

     config.vm.define "nfs" do |node|
     	node.vm.box = "bento/centos-7"
     	node.vm.network "private_network", ip:"192.168.10.15"
     	node.vm.hostname = "nfs"
     	node.vm.provider "virtualbox" do |vb|
     		vb.name = "nfs"
     		vb.memory = 1024
     		vb.cpus = 1
     		vb.customize ["modifyvm", :id, "--audio", "none"]
     	end
     	node.vm.provision "shell", path: "bootstrap_nfs.sh"
     end
end
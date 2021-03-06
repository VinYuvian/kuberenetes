# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do | config |
	config.vm.define "kmaster" do |node|
		node.vm.hostname = "kmaster"
		node.vm.network "private_network", ip: "192.168.10.10"
		node.vm.box = "bento/centos-7"
		node.vm.synced_folder "resources/","/home/vagrant/resources"
		node.vm.provider "virtualbox" do |v|
			v.name = "kmaster"
			v.memory = 2048
			v.cpus = 2
			v.customize ["modifyvm", :id, "--audio", "none"]
		end
		node.vm.provision "shell", path: "boot_strap.sh"
		node.vm.provision "shell", path: "boot_strap_kmaster.sh"
	end

	config.vm.define "jenkins" do |node|
		node.vm.hostname = "jenkins"
		node.vm.network "private_network", ip: "192.168.10.15"
		node.vm.network :forwarded_port, host: 8080, guest: 8080
		node.vm.network :forwarded_port, host: 50000, guest: 50000
		node.vm.box = "bento/centos-7"
		node.vm.synced_folder "resources/","/home/vagrant/resources"
		node.vm.provider "virtualbox" do |v|
			v.name = "jenkins"
			v.memory = 1024
			v.cpus = 1
			v.customize ["modifyvm", :id, "--audio", "none"]
		end
		node.vm.provision "shell", path: "bootstrap_jenkins.sh"
	end

	nodeCount = 2

	(1..nodeCount).each do |i|
		config.vm.define "kworker#{i}" do | node |
			node.vm.box = "bento/centos-7"
			node.vm.hostname = "kworker#{i}"
			node.vm.network "private_network", ip: "192.168.10.1#{i}"
			node.vm.provider "virtualbox" do |v|
     		   v.name = "kworker#{i}"
  	     	   v.memory = 2048
               v.cpus = 2
        # Prevent VirtualBox from interfering with host audio stack
               v.customize ["modifyvm", :id, "--audio", "none"]
            end
            node.vm.provision "shell", path: "boot_strap.sh"
            node.vm.provision "shell", path: "bootstrap_kworker.sh"
        end
     end
end

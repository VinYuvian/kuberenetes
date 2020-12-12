#!/bin/bash

echo "[TASK1] set up hostnames"
#setting up hostnames
cat >>/etc/hosts<<EOF
192.168.10.10 kmaster kmaster
192.168.10.11 kworker1 kworker1
192.168.10.12 kworker2 kworker2
192.168.10.15 nfs nfs
EOF

echo "[TASK2] Install nfs-utils"
yum install -y -q nfs-utils


echo "[TASK3] Starting and enabling nfs-server"
systemctl start nfs-server rpcbind
systemctl enable nfs-server rpcbind


echo "[TASK4] Configuring shared folder"
mkdir /home/vagrant/app
chmod -R 755 /home/vagrant/app
chown nfsnobody:nfsnobody /home/vagrant/app


echo "[TASK5] Configuring exports file"
cat >> /etc/exports <<EOL 
/home/vagrant/app 192.168.0.0/16(rw,sync,no_root_squash)
EOL


echo "[TASK6] Exporting FS"
exportfs -r

#mv /home/vagrant/reactapp.com/* /home/vagrant/app/


echo "[TASK7] Firewall Configuration "
systemctl start firewalld
firewall-cmd --permanent --add-service mountd
firewall-cmd --permanent --add-service rpc-bind
firewall-cmd --permanent --add-service nfs
firewall-cmd --reload
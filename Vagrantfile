#
#-*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  config.vm.define "lamia" do |lamia1|
    lamia1.vm.box = "bento/ubuntu-18.10"
    lamia1.vm.network "private_network", ip: "192.169.69.10"
    lamia1.vm.synced_folder ".", "/home/vagrant/files"
    lamia1.vm.network "forwarded_port", guest:80, host: 8000
    # Here we can add shell provisionner but we prefer to use playbook and ansible 
    # to configure our VM
  end 
 end



Vagrant.configure("2") do |config|

  config.vm.box = "centos7docker19"


  config.vm.define "node03" do |node03|
    node03.vm.hostname = "node03"
    node03.vm.network "private_network", ip: "192.168.33.13"
  end


  config.vm.provider "virtualbox" do |vb|
    vb.cpus = 2
    vb.memory = "4096"
  end

end

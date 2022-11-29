#!/bin/bash
cd /root
tar -xvf {{etcd_package}}.tar.gz && \
  cd {{etcd_package}} && \
  sudo cp -a etcd etcdctl /usr/bin/
nohup etcd >/tmp/etcd.log 2>&1 &
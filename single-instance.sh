#!/bin/bash

export ANSIBLE_LIBRARY=`pwd`/plugins/modules && ansible-playbook -e hostgroup=all -i inventory/hosts.yml single-instance-fs.yml

# a2z_ansible

## 功能
安装各种中间件、数据库、微服务组件

## 安装

- a2z_ansible 依赖 Ansible，所以需要在 Ansible controller（控制面）  上安装 Ansible 


```shell
export ANSIBLE_LIBRARY=`pwd`/plugins/modules && ansible-playbook -e hostgroup=oracle -i inventory/hosts.yml single-oracle.yml

export PATH=/root/.local/bin/:$PATH && ansible-playbook -e hostgroup=mysql -i inventory/hosts.yml install-mysql.yml
```
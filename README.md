# a2z-ansible

```shell
export ANSIBLE_LIBRARY=`pwd`/plugins/modules && ansible-playbook -e hostgroup=oracle -i inventory/hosts.yml single-oracle.yml

export PATH=/root/.local/bin/:$PATH && ansible-playbook -e hostgroup=mysql -i inventory/hosts.yml install-mysql.yml
```
# a2z-ansible

```shell
export ANSIBLE_LIBRARY=`pwd`/plugins/modules && ansible-playbook -e hostgroup=oracle -i inventory/hosts.yml single-instance-fs.yml

ansible-playbook -e hostgroup=mysql -i inventory/hosts.yml single-mysql.yml
```
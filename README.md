# a2z_ansible

## 功能
安装各种中间件、数据库、微服务组件

## 安装

- a2z_ansible 依赖 Ansible，所以需要在 Ansible controller（控制面）  上安装 Ansible 


## 使用

- 拷贝公钥到待安装的主机上

```shell
ssh-copy-id -i ~/.ssh/id_rsa.pub <host>
```

- 修改 Ansible controller 主机上的 /etc/hosts

```shell
# vi /etc/hosts
192.168.3.215   mysql57.db # mysql 对应的主机
192.168.3.216   app.gw # nginx redis nacos 对应的主机
192.168.3.217   app.backend # springboot 对应的主机
```

- 安装 mysql 单节点

    - 配置 MySQL 安装参数
    
    修改 roles/mysql_install/defaults/main.yml 文件
    
    ```shell
    # mysql 安装包在Ansible controller上所在的目录
    mysql_packages_dir: /root/softwares/mysql/

    # mysql 安装包的名字
    mysql_package: mysql-5.7.39-linux-glibc2.12-x86_64.tar.gz
    # mysql_package: mysql-8.0.30-linux-glibc2.17-x86_64-minimal.tar.xz

    # mysql 安装目录
    mysql_base_dir: /usr/local/mysql/

    # mysql 真正的 datadir 就会是 mysql_data_dir_base + mysql_port
    mysql_data_dir_base: /glzt/mysql/data/

    # mysql 端口
    mysql_port: 3306

    # root 密码
    mysql_root_password: mtls0352

    ```

    - 执行 mysql 安装 playbook

    ```shell
    export PATH=/root/.local/bin/:$PATH && ansible-playbook -e hostgroup=mysql -i inventory/hosts.yml install-mysql.yml
    ```

- 安装 nacos redis nginx

    - 配置安装参数

    roles/nacos_install/defaults/main.yml（安装包在本地）

    ```shell
    # nacos 安装包在Ansible controller上所在的目录
    packages_dir: /root/softwares/nacos
    nacos_package: nacos-server-2.1.2.zip

    # 是否安装 jdk
    java_install: true

    # jdk 包名称
    java_package: 
    - java-11-openjdk.x86_64
    - java-11-openjdk-devel.x86_64
    
    # nacos 安装目录
    nacos_home: /opt
    ```

    roles/redis_install/defaults/yml（安装包在网上）

    ```shell
    # redis 版本
    redis_version: 5.0.0
    # redis 安装目录
    redis_install_dir: /opt/redis
    # redis 端口
    redis_dir: /var/lib/redis/{{ redis_port }}

    ```

    roles/nginx_install/defaults/yml（安装包在本地）

    ```shell
    packages_dir: /root/softwares/nginx
    nginx_package: nginx-1.23.0

    ```

    - 执行安装 playbook

    ```shell
    export PATH=/root/.local/bin/:$PATH && ansible-playbook -e hostgroup=application -i inventory/hosts.yml install-application.yml
    ```

- 安装 springboot 后端服务

    - 配置 安装参数
    
    修改 roles/springboot_install/defaults/main.yml 文件
    
    ```shell
    # 是否需要安装 jdk
    springboot_java_install: false

    # 需要安装后端服务的信息
    springboot_applications:
    - { deploy_folder: '/opt/record', src_file: '/root/softwares/application/Record-1.0.0-SNAPSHOT.jar', application_name: 'Record-1.0.0-SNAPSHOT.jar', configuration_template: false, startup: 'startup.sh' }
    springboot_group: springboot
    springboot_user: springboot

    ```

    - 执行 mysql 安装 playbook

    ```shell
    export PATH=/root/.local/bin/:$PATH && ansible-playbook -e hostgroup=mysql -i inventory/hosts.yml install-mysql.yml
    ```

<!-- ```shell
export ANSIBLE_LIBRARY=`pwd`/plugins/modules && ansible-playbook -e hostgroup=oracle -i inventory/hosts.yml single-oracle.yml
``` -->
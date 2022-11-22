{% if 'mysql-5.6' in mysql_package %}
set sql_log_bin=0;
    update mysql.user set password=password('{{ mysql_root_password }}') where user='root';
    create user {{mysql_monitor_user}}@'127.0.0.1' identified by '{{ mysql_monitor_password }}' ;
    grant replication client on *.* to {{mysql_monitor_user}}@'127.0.0.1';
    grant process on *.* to {{mysql_monitor_user}}@'127.0.0.1';

    create user {{mysql_monitor_user}}@'localhost' identified by '{{ mysql_monitor_password }}' ;
    grant replication client on *.* to {{mysql_monitor_user}}@'localhost';
    grant process on *.* to {{mysql_monitor_user}}@'localhost';

    create user {{mysql_backup_user}}@'127.0.0.1' identified by '{{mysql_backup_password}}';
    create user {{mysql_backup_user}}@'localhost' identified by '{{mysql_backup_password}}';
    
    grant reload,lock tables on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant replication client on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant create tablespace  on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant process            on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant super              on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant create,insert,select      on percona_schema.xtrabackup_history to {{mysql_backup_user}}@'127.0.0.1';
    
    grant reload,lock tables on *.* to {{mysql_backup_user}}@'localhost';
    grant replication client on *.* to {{mysql_backup_user}}@'localhost';
    grant create tablespace  on *.* to {{mysql_backup_user}}@'localhost';
    grant process            on *.* to {{mysql_backup_user}}@'localhost';
    grant super              on *.* to {{mysql_backup_user}}@'localhost';
    grant create,insert,select      on percona_schema.xtrabackup_history to {{mysql_backup_user}}@'localhost';

    grant create,insert,drop,update               on mysql.backup_progress to backup@'127.0.0.1';
    grant create,insert,drop,update,select,alter  on mysql.backup_history  to backup@'127.0.0.1';
    
    grant create,insert,drop,update               on mysql.backup_progress to backup@'localhost';
    grant create,insert,drop,update,select,alter  on mysql.backup_history  to backup@'localhost';

    flush privileges;

set sql_log_bin=1;

{% else %}

set sql_log_bin=0;
    alter user root@'localhost' identified by '{{ mysql_root_password }}' ;
    create user root@'127.0.0.1' identified by '{{ mysql_root_password }}';
    grant all on *.* to root@'127.0.0.1' with grant option;

    create user {{mysql_monitor_user}}@'127.0.0.1' identified by '{{ mysql_monitor_password }}' ;
    grant replication client on *.* to {{mysql_monitor_user}}@'127.0.0.1';
    grant process on *.* to {{mysql_monitor_user}}@'127.0.0.1';

    create user {{mysql_monitor_user}}@'localhost' identified by '{{ mysql_monitor_password }}' ;
    grant replication client on *.* to {{mysql_monitor_user}}@'localhost';
    grant process on *.* to {{mysql_monitor_user}}@'localhost';

    create user {{mysql_backup_user}}@'127.0.0.1' identified by '{{mysql_backup_password}}';
    create user {{mysql_backup_user}}@'localhost' identified by '{{mysql_backup_password}}';
    
    grant reload,lock tables on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant replication client on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant create tablespace  on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant process            on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant super              on *.* to {{mysql_backup_user}}@'127.0.0.1';
    grant create,insert,select      on percona_schema.xtrabackup_history to {{mysql_backup_user}}@'127.0.0.1';
    
    grant reload,lock tables on *.* to {{mysql_backup_user}}@'localhost';
    grant replication client on *.* to {{mysql_backup_user}}@'localhost';
    grant create tablespace  on *.* to {{mysql_backup_user}}@'localhost';
    grant process            on *.* to {{mysql_backup_user}}@'localhost';
    grant super              on *.* to {{mysql_backup_user}}@'localhost';
    grant create,insert,select      on percona_schema.xtrabackup_history to {{mysql_backup_user}}@'localhost';

    grant create,insert,drop,update               on mysql.backup_progress to backup@'127.0.0.1';
    grant create,insert,drop,update,select,alter  on mysql.backup_history  to backup@'127.0.0.1';
    
    grant create,insert,drop,update               on mysql.backup_progress to backup@'localhost';
    grant create,insert,drop,update,select,alter  on mysql.backup_history  to backup@'localhost';
    
    grant select on performance_schema.replication_group_members to backup@'127.0.0.1';
    grant select on performance_schema.replication_group_members to backup@'localhost';
set sql_log_bin=1;

{% endif %}


This project has the sample go code to connect to the elastic search with shard.


Go steps to initialze the project and run

    go mod init Elastic-search-example
    go mod tidy
    go build
    go run main.go





Steps followed to install server

https://www.elastic.co/guide/en/elasticsearch/reference/8.13/rpm.html#install-rpm

History

[root@localhost ~]# history
    1  rpm --import https://artifacts.elastic.co/GPG-KEY-elasticsearch
    2  mkdir -p /etc/zypp/repos.d/
    3  ls
    4  cd /etc/zypp/repos.d/
    5  ls
    6  vi elasticsearch.repo
    7  cd ~
    8  sudo yum install --enablerepo=elasticsearch elasticsearch
    9  cd /etc/zypp/repos.d/
   10  sudo yum install --enablerepo=elasticsearch elasticsearch
   11  cd ~
   12  mkdir -p /opt/dev/
   13  cd /opt/dev/
   14  ls
   15  wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.13.0-x86_64.rpm
   16  yum install wget
   17  wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.13.0-x86_64.rpm
   18  ls -ltr
   19  mkdir software
   20  mv elasticsearch-8.13.0-x86_64.rpm software/
   21  cd software/
   22  ls
   23  sudo rpm --install elasticsearch-8.13.0-x86_64.rpm 
   24  sudo /bin/systemctl daemon-reload
   25  sudo /bin/systemctl enable elasticsearch.service
   26  sudo systemctl start elasticsearch.service
   27  sudo systemctl stop elasticsearch.service
   28  echo "keystore_password" > /opt/dev/software/my_pwd_file.tmp
   29  cat /opt/dev/software/my_pwd_file.tmp
   30  chmod 600 /opt/dev/software/my_pwd_file.tmp 
   31  sudo systemctl set-environment ES_KEYSTORE_PASSPHRASE_FILE=/opt/dev/software/my_pwd_file.tmp
   32  sudo systemctl start elasticsearch.service
   33  systemctl status elasticsearch.service
   34  chmod 600 /opt/dev/software/my_pwd_file.tmp
   35  ls -ltr
   36  chmod 777 /opt/dev/software/my_pwd_file.tmp
   37  ls -ltr
   38  sudo systemctl start elasticsearch.service
   39  curl http://192.168.0.133:9200/
   40  curl http://192.168.0.133:9200
   41  curl http://192.168.0.133
   42  sudo journalctl -f
   43  curl http://192.168.0.133:9060
   44  sudo journalctl --unit elasticsearch
   45  sudo journalctl --unit elasticsearch --since  "2024-03-30 10:28:45"
   46  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200
   47  ls
   48  cat my_pwd_file.tmp 
   49  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200 pretty
   50  export ELASTIC_PASSWORD="your_password"
   51  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200
   52  sudo systemctl stop elasticsearch.service
   53  sudo systemctl start elasticsearch.service
   54  sudo journalctl -f
   55  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200
   56  curl --cacert /etc/elasticsearch/certs/http_ca.crt https://localhost:9200
   57  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:your_password https://localhost:9200
   58  https://localhost:9200
   59  curl https://localhost:9200
   60  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:your_password https://localhost:9200 pretty"
   61  curl -XGET 'http://localhost:9200/_cluster/state?pretty'
   62  curl -XGET 'https://localhost:9200/_cluster/state?pretty'
   63  curl --cacert /etc/elasticsearch/certs/http_ca.crt  -XGET 'https://localhost:9200/_cluster/state?pretty'
   64  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:your_password  -XGET 'https://localhost:9200/_cluster/state?pretty'
   65  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:your_password  -XGET 'https://localhost:9200/_xpack/license?pretty'
   66  /usr/share/elasticsearch/bin/elasticsearch-reset-password -u elastic
   67  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:fM3dfFJ6Vn6R8GGJo=Ee  -XGET 'https://localhost:9200/_xpack/license?pretty'
   68  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:fM3dfFJ6Vn6R8GGJo=Ee https://localhost:9200
   69  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:fM3dfFJ6Vn6R8GGJo=Ee  -XGET 'https://localhost:9200/_xpack/license?pretty'
   70  curl --cacert /etc/elasticsearch/certs/http_ca.crt -u elastic:fM3dfFJ6Vn6R8GGJo=Ee https://localhost:9200
   71  sudo systemctl status firewalld
   72  sudo systemctl disable firewalld
   73  sudo systemctl status firewalld
   74  sudo systemctl stop firewalld
   75  sudo systemctl status firewalld
   76  history
[root@localhost ~]#

Make file is not used now, its a place holder

All the best in future
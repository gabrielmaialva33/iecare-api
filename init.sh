echo "
  [Unit]
  Description=IECare
  After=network.target
  [Service]
  User=ubuntu
  WorkingDirectory=/home/ubuntu/iecare-api/bin
  ExecStart=/home/ubuntu/iecare-api/bin/iecare
  Restart=always
  RestartSec=3
  StartLimitInterval=0
  [Install]
  WantedBy=multi-user.target
  " > /etc/systemd/system/iecare.service


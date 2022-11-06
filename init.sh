if [ ! -f /etc/systemd/system/iecare.service ]
            then
              echo "[Unit]
              Description=iecare
              After=network.target

              [Service]
              User=ubuntu
              WorkingDirectory=/home/ubuntu/iecare-api
              ExecStart=/home/ubuntu/iecare-api/bin/iecare
              Restart=always
              RestartSec=3
              StartLimitInterval=0

              [Install]
              WantedBy=multi-user.target" > /etc/systemd/system/iecare.service
            fi
            sudo systemctl daemon-reload
            sudo systemctl enable iecare.service
            sudo systemctl start iecare.service



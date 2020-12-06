#!/bin/bash

cp -R /tmp/.ssh/* /home/devel/.ssh
chmod 700 /home/devel/.ssh
find /home/devel/.ssh -name "id_rsa*" -not -name *.pub -exec chmod 600 {} \;
find /home/devel/.ssh -name "id_rsa*" -name *.pub -exec chmod 644 {} \;

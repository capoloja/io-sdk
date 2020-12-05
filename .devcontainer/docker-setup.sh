#!/bin/bash

cp -R /tmp/.ssh/* /root/.ssh
chmod 700 /root/.ssh
find /root/.ssh -name "id_rsa*" -not -name *.pub -exec chmod 600 {} \;
find /root/.ssh -name "id_rsa*" -name *.pub -exec chmod 644 {} \;

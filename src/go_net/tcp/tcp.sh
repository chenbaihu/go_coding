#!/bin/bash

./tcp -hostPort="127.0.0.1:2007" \
            -reqNum=0 \
            -reqNumPerConn=100 \
			-c=100 \
			-keepAlive=true

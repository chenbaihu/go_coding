#!/bin/bash

./http_post -url="http://testserver:8078/echo.php" \
            -reqNum=0 \
            -reqNumPerConn=100 \
			-c=100 \
			-keepAlive=true

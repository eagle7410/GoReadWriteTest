#!/usr/bin/env bash
docker run -it  igor/test:1 sh
docker run -v /home/igor/go/src/GoReadWriteTest/vol:/go/src/app -it  igor/test:1 sh

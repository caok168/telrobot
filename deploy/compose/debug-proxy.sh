#!/bin/bash

socat TCP4-LISTEN:5432,reuseaddr,fork TCP4:postgres:5432 &

tail -f /dev/null

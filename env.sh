#!/bin/bash

export $(cat .env)
export TIME_ZONE=Asia/Taipei

export HTTP_LISTEN_ADDR=0.0.0.0
export HTTP_LISTEN_PORT=8080

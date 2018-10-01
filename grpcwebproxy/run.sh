#!/bin/bash

grpcwebproxy --backend_tls=false --backend_addr=127.0.0.1:8888 --server_bind_address=127.0.0.1 --server_http_debug_port=7777 --run_tls_server=false

#!/bin/bash

sudo pfctl -F all -f /etc/pf.conf

echo "Current port forwarding:"

sudo pfctl -s nat

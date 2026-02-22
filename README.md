# Tailscale One-Line Installer and Exit Node Configurator for Linux
This repository provides a convenient one-liner script to install and configure Tailscale on Linux systems, including setting up an exit node. It also includes some Linux optimization commands to improve the system's performance. Tested and confirmed working in Iran using MCI and Wi-Fi networks (not Irancell).

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Optimization](#optimization)
- [Configuration](#configuration)
- [Connecting to the Exit Node](#connecting-to-the-exit-node)
- [Compatibility](#compatibility)
- [Comparison with Other Solutions](#comparison-with-other-solutions)
- [Troubleshooting](#troubleshooting)

## Introduction
Tailscale is a zero-configuration VPN that allows devices to communicate with each other securely. It is easy to set up and use, making it a great option for individuals and organizations alike. This repository provides a simple way to install and configure Tailscale on Linux systems, including setting up an exit node.

### Architecture Overview
The following Mermaid diagram shows the high-level architecture of the Tailscale setup:
```mermaid
graph LR
    A[Client Device] -->|Secure Connection|> B[Tailscale Server]
    B -->|Exit Node|> C[Internet]
    C -->|Secure Connection|> D[Server]
    D -->|Secure Connection|> B
    B -->|Secure Connection|> A
```
This diagram shows how the client device connects to the Tailscale server, which then routes the traffic through the exit node to the internet.

## Installation
To install Tailscale and configure an exit node on your Linux machine, run the following command:
```bash
curl -sSL https://raw.githubusercontent.com/mehrshud/tailscale/master/install.sh | sudo bash
```
This command will download and run the installation script, which will install Tailscale and configure the exit node.

### Example Installation Script
Here is an example of what the installation script might look like:
```bash
#!/bin/bash

# Install Tailscale
curl -sSL https://pkgs.tailscale.com/install.sh | sudo bash

# Configure the exit node
export EXIT_NODE_NAME="my-exit-node"
export EXIT_NETS="0.0.0.0/0, ::/0"

# Run the Tailscale installation script
tailscale up --exit-node=$EXIT_NODE_NAME --exit-nets=$EXIT_NETS
```
This script installs Tailscale, configures the exit node, and starts the Tailscale service.

## Optimization
The script includes some Linux optimization commands to improve the system's performance. These optimizations include:

- Enabling TCP BBR congestion control
- Adjusting system file limits and kernel parameters for better network performance
- Disabling Transparent Huge Pages (THP) for reduced CPU usage

### Optimization Example
Here is an example of how to optimize the system for better network performance:
```bash
# Enable TCP BBR congestion control
sysctl -w net.ipv4.tcp_congestion_control=bbr

# Adjust system file limits
ulimit -n 100000

# Disable Transparent Huge Pages (THP)
echo never > /sys/kernel/mm/transparent_hugepage/enabled
```
These optimizations can help improve the system's performance and reduce latency.

## Configuration
To configure the exit node, you need to set some environment variables before running the installation script. Below is an example of how to set these variables:
```bash
export EXIT_NODE_NAME="my-exit-node"
export EXIT_NETS="0.0.0.0/0, ::/0"
curl -sSL https://raw.githubusercontent.com/mehrshud/tailscale/master/install.sh | sudo bash
```
- `EXIT_NODE_NAME`: The name of your exit node.
- `EXIT_NETS`: The CIDR ranges for the exit node (default is to allow all traffic).

### Configuration Options
Here are some additional configuration options that you can use:

- `EXIT_NODE_PORT`: The port number to use for the exit node (default is 51820).
- `EXIT_NODE_PROTOCOL`: The protocol to use for the exit node (default is UDP).
- `EXIT_NODE_AUTH`: The authentication method to use for the exit node (default is none).

## Connecting to the Exit Node
After setting up the exit node, you can connect to it using Tailscale clients on different platforms. Instructions for each platform can be found in the [Tailscale documentation](https://tailscale.com/kb/). Make sure to set the exit node IP address to the configured exit node on your server.

### Connection Example
Here is an example of how to connect to the exit node using the Tailscale client:
```bash
tailscale up --exit-node=my-exit-node
```
This command will connect to the exit node and establish a secure connection.

## Compatibility
This script is compatible with most Linux distributions, including Ubuntu, Debian, and CentOS.

### Compatibility Matrix
Here is a compatibility matrix showing which Linux distributions are supported:
| Distribution | Version | Supported |
| --- | --- | --- |
| Ubuntu | 18.04, 20.04 | |
| Debian | 9, 10, 11 | |
| CentOS | 7, 8 | |

## Comparison with Other Solutions
Here is a comparison table showing how Tailscale stacks up against other VPN solutions:
| Solution | Ease of Use | Security | Performance |
| --- | --- | --- | --- |
| Tailscale | Easy | High | High |
| OpenVPN | Medium | High | Medium |
| WireGuard | Easy | High | High |
| StrongSwan | Hard | High | Medium |

### Comparison Example
Here is an example of how Tailscale compares to OpenVPN:
```bash
# Tailscale installation script
curl -sSL https://raw.githubusercontent.com/mehrshud/tailscale/master/install.sh | sudo bash

# OpenVPN installation script
sudo apt-get install openvpn
```
As you can see, Tailscale is much easier to install and configure than OpenVPN.

## Troubleshooting
If you encounter any issues during installation or configuration, here are some troubleshooting tips:

- Check the system logs for errors
- Run the installation script with the `--debug` flag to get more verbose output
- Check the Tailscale documentation for known issues and workarounds

### Troubleshooting Example
Here is an example of how to troubleshoot an issue with the Tailscale client:
```bash
# Run the Tailscale client with the --debug flag
tailscale up --exit-node=my-exit-node --debug

# Check the system logs for errors
sudo journalctl -u tailscale
```
By following these troubleshooting steps, you should be able to resolve any issues you encounter.
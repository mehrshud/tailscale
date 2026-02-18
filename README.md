# Tailscale One-Line Installer and Exit Node Configurator for Linux 🚀

[![CI](https://github.com/mehrshud/tailscale/actions/workflows/ci.yml/badge.svg)](https://github.com/mehrshud/tailscale/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/language-Go-blue)]()
[![Stars](https://img.shields.io/github/stars/mehrshud/tailscale?style=social)](https://github.com/mehrshud/tailscale)



This repository provides a convenient one-liner script to install and configure Tailscale on Linux systems, including setting up an exit node. It also includes some Linux optimization commands to improve the system's performance. Tested and confirmed working in Iran using MCI and Wi-Fi networks (not Irancell).

## Table of Contents

- [Installation](#installation)
- [Optimization](#optimization)
- [Configuration](#configuration)
- [Connecting to the Exit Node](#connecting-to-the-exit-node)
- [Compatibility](#compatibility)

## Installation 💻

To install Tailscale and configure an exit node on your Linux machine, run the following command:

```bash
curl -sSL https://raw.githubusercontent.com/mehrshud/tailscale/master/install.sh | sudo bash
```

## Optimization ⚡

The script includes some Linux optimization commands to improve the system's performance. These optimizations include:

- Enabling TCP BBR congestion control
- Adjusting system file limits and kernel parameters for better network performance
- Disabling Transparent Huge Pages (THP) for reduced CPU usage

## Configuration 🛠️

To configure the exit node, you need to set some environment variables before running the installation script. Below is an example of how to set these variables:

```bash
export EXIT_NODE_NAME="my-exit-node"
export EXIT_NETS="0.0.0.0/0, ::/0"
curl -sSL https://raw.githubusercontent.com/mehrshud/tailscale/master/install.sh | sudo bash
```

- `EXIT_NODE_NAME`: The name of your exit node.
- `EXIT_NETS`: The CIDR ranges for the exit node (default is to allow all traffic).

## Connecting to the Exit Node 🌐

After setting up the exit node, you can connect to it using Tailscale clients on different platforms. Instructions for each platform can be found in the [Tailscale documentation](https://tailscale.com/kb/). Make sure to set the exit node IP address to the configured exit node on your server.

## Compatibility 🌍

This script has been tested and confirmed working in Iran using MCI and Wi-Fi networks. Note that it does not work with Irancell.

## License

This project is licensed under the terms of the MIT License.

## Support

If you encounter any issues or have any questions, please create an issue in this repository.

---

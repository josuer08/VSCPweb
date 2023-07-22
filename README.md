## VSCPweb

VSCPweb is a web based software defined networking playground that uses openvswitches (OVS) and a modified version of [mininet](https://github.com/mininet/mininet).
It is a tool that allows you to:
+ Create different OVS topologies
+ Connect SDN controllers to the virtual network
+ Use linux namespace based lightweight containers as hosts or servers in the network
+ Monitor the health of all of the worker servers
+ Have CLI access to all of the hosts of the virtual network
+ Have CLI access to mininet

## Table of Contents

* Introduction: #introduction
* Installation: #installation
* Usage: #usage
* Contributing: #contributing
* License: #license

## Introduction

VSCPweb is a web based network simulator that can be used to learn about and experiment with different network topologies and configurations. It is a valuable tool for anyone who wants to learn more about networking or who needs to test out new network configurations.

VSCPweb is made with golang and htmx. It uses openvswitches as the underlying network virtualization layer. This allows VSCPweb to create realistic and scalable network simulations.

## Installation

To install VSCPweb, you will need to have the following dependencies installed:

+ golang
+ HTMX
+ mininet

We are making all of the effort possible into later making distributuion packages that are as distro agnostic as possible but we woul likely not work on window or mac versions (you can still install in a linux VM and access the web console with your browser)

Once you have the dependencies installed, you can install VSCPweb by running the following command:

```
THIS IS NOT READY YET
```

## Usage

To start VSCPweb, run the following command:

```
vscpweb
```

VSCPweb will start a web server on port 8080. You can access the VSCPweb interface at http://localhost:8080.

## Contributing

VSCPweb is an open source project and we welcome contributions from everyone. If you find a bug or have an idea for a new feature, please open an issue or submit a pull request.

## Demo

We plan to have a small limited demo running in https://vscp.josue.xyz but it is not yet public due to the extent of control that vscp can excert over the host system.

## License

VSCPweb is licensed under the GPLv3 License.
+ The GPLv3 license is published by the Free Software Foundation (FSF).
+ The GPLv3 license is version 3 of the GNU General Public License (GPL).
+ The GPLv3 license is a copyleft license.
+ The GPLv3 license prohibits patent claims on the software.
+ The GPLv3 license has stronger copyleft provisions than the GPLv2 license.
+ The GPLv3 license prohibits the use of DRM technologies in the software.
See the LICENSE file for more information.

## Contact

If you have any questions or concerns, please contact leave an issue on GitHub.
## Thor Chain

Official thor implementation.

## Building the source


Building `thor` requires both a Go (version 1.15 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
cd thorlinkgo\thor\cmd\thor
go build
```



## Executables

The thor project comes with an executable binary file found in the `cmd\thor`
directory.

|    Command    | Description                                                                                                                                                                                                                                      |
| :-----------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  **`thor`**   | Our main thor node binary file. It is the entry point into the thor network, capable of running as a full node (default), archive node (retaining all historical state). It can be used by other processes as a gateway into the thor network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports.          |
                                                                                                                                                

## License

The thor library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The tbinaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.

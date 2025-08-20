# go-grpc

There is a `.devcontainer` folder for quickly getting up and running when using Visual Studio Code.
If you're not using Visual Studio Code you can still look at that folder as a reference for how to setup a development environment.
The `.devcontainer/devcontainer.json` file uses environment variables to set the username and password for the package repository (`build.args`). If you don't want to set these environment variables you can set these values manually in that file instead.

To start smtpd, go into `.devcontainer/smtpd` and run `halonconfig` followed by `supervisorctl start smtpd`

To look at the `greeter_server` logs run `supervisorctl tail -f greeter_server stderr`.

To test the `greeter_client`  logs run `/grpc-go/examples/helloworld/greeter_client/main`.
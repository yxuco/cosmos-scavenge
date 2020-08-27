# cosmos-scavenge

Detailed steps of implementation are described in [Cosmos SDK Tutorial](https://tutorials.cosmos.network/scavenge/tutorial/03-scaffold.html#). You may go directly to `Build app` below to build and start the app.

## Download and build `scaffold`

```bash
mkdir -p ~/work/cosmos-tutorial
cd ~/work/cosmos-tutorial
git clone git@github.com:cosmos/scaffold.git
cd scaffold
make
```

## Generate app and module

```bash
cd ~/work/cosmos-tutorial
scaffold app lvl-1 yxuco cosmos-scavenge
cd cosmos-scavenge/x
scaffold module yxuco cosmos-scavenge scavenge
cd .. && go mod tidy
```

## Define messages

Define messages using the generated template `x/scavenge/types/msg.go`

- implement sdk.Msg interface: <https://github.com/cosmos/cosmos-sdk/blob/master/types/tx_msg.go>
- Register messages in `x/scavenge/types/codec.go`
- Create `x/scavenge/alias.go` to make nested definitions of types and keeper available at module's top level

## Define key store and keeper functions

- In `x/scavenge/types/types.go`, define data struct for key store
- In `x/scavenge/types/key.go`, define key prefixes for each key store data
- In `x/scavenge/keeper/keeper.go`, define functions for accessing key store

## Implement message handler

Implement handler functions to processe each message and interact with key store.

- Edit generated templates in `x/scavenge/handler.go`
- Define event types in `x/scavenge/types/events.go` (events help client apps to know what happened in a transaction)

## Implement querier to query key store

- Implement query functions in `x/scavenge/keeper/querier.go`
- Define query types in `x/scavenge/types/querier.go`

## Implement CLI

Implement CLI app to send transaction and query.

- Implement functions for creating transactions in `x/scavenge/client/cli/tx.go`
- Implement functions for creating queries in `x/scavenge/client/cli/query.go`

## Make module reusable by any app

- Edit `x/scavenge/module.go` to add bank.Keeper to AppModule struct and NowAppModule function

## Use module in app

- In `app/app.go`, follow comments of `TODO` to change name of appd appcli and add module refs
- Change folder name of `cmd/appd` and `cmd/appcli`
- Edit rootCmd description in `cmd/appcli/main.go`
- Edit rootCmd description in `cmd/appd/main.go`

## Build app

- At app root, `go mod tidy`
- Edit Makefile to set `appd`and `appcli`
- make
- Verify executable: `scavengeD -h`, and `scavengeCLI -h`

## Config and start a node

```bash
./init.sh
scavengeD start
```

## Play scavenge

Send transaction to create scavenge:

```bash
scavengeCLI tx scavenge createScavenge 69foo "A stick" "What's brown and sticky?" --from me
```

Send query to verify that the scavenge of the resulting `txhash` is created:

```bash
scavengeCLI q tx <txhash>
```

Send transaction to commit a solution:

```bash
scavengeCLI tx scavenge commitSolution "A stick" --from you
```

Send query to verify that a solution is committed:

```bash
scavengeCLI q scavenge commited "A stick" $(scavengeCLI keys show you -a)
```

Send transaction to reveal the solution:

```bash
scavengeCLI tx scavenge revealSolution "A stick" --from you
```

Send query to verify that the scavenger account received 69foo coins:

```bash
scavengeCLI q account $(scavengeCLI keys show you -a)
```

Send query to list scavenges and show the solutions:

```bash
scavengeCLI q scavenge list
scavengeCLI q scavenge get <solutionHash>
```

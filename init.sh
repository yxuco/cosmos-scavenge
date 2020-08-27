#!/bin/bash
rm -r ~/.scavengeCLI
rm -r ~/.scavengeD

# initialize application node as mynode, create ~/.scavengeD
scavengeD init mynode --chain-id scavenge

scavengeCLI config keyring-backend test

# create 2 test users in mac keyring
scavengeCLI keys add me
scavengeCLI keys add you

# add validator account 'me' with stakes, update ~/.scavengeD/config/genesis.json
scavengeD add-genesis-account $(scavengeCLI keys show me -a) 1000foo,100000000stake
# add user account 'you' with only token 'foo', no stakes
scavengeD add-genesis-account $(scavengeCLI keys show you -a) 1foo

# config CLI to talk to only scavenge, create ~/.scavengeCLI
scavengeCLI config chain-id scavenge
scavengeCLI config output json
scavengeCLI config indent true
scavengeCLI config trust-node true

scavengeD gentx --name me --keyring-backend test

# create genesis transaction to crate a validator for 'me', tx written to ~/.scavengeD/config/gentx/
scavengeD gentx --name me

# collect all config, update genesis.json, prepare to start app
scavengeD collect-gentxs

# start app
# scavengeD start
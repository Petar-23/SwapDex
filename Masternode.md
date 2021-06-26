
# SwapDex Masternode Guide

This guide is for developer don't use if you don't know what are you doing.

## Requirements for running masternode:-
1} Need ubuntu server to run masternode.Atleast 16gb ram 120gb ssd & 2core cpu.

2} 10000sdx for Masternode

3} create metamask wallet & transfer 10010 sdx from exchange to that wallet.

## You are now ready to start installation process.



# Install Go & Gvm. 

For building sdx you need go & gvm. This guide is for ubuntu 20.
For proper installation guide follow this guide:-https://github.com/moovweb/gvm

```bash 
sudo apt update
sudo apt upgrade
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
  
```
# Note:- Restart the terminal after first command.

## Debian/Ubuntu
```bash
sudo apt-get install curl git mercurial make binutils bison gcc build-essential
```
List Go Versions
To list all installed Go versions (The current version is prefixed with "=>"):
```bash
gvm list
```
To list all Go versions available for download:
```bash
gvm listall
```
Compiling Go 1.5+

```bash
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.5
```

# Installing Go 1.13.8+
```bash
gvm install go1.13.8
gvm use go1.13.8
```

# Lets start building.
```bash
git clone https://github.com/69th-byte/SmartDex-Chain.git sdxchain
cd sdxchain 
make all
mv /root/sdxchain/build/bin/sdx /usr/bin/sdx
```
# Last step of running Masternode.
```bash

sdx account new
$ give secure password

sdx init swapdex.json
``` 
```bash
nano pass.txt
```

## $ Put your password here,which you used to create coinbase account.
ctrl + x & y enter to save

```bash
sdx  --syncmode "full" --networkid 7879 --port 10303 --rpc --rpccorsdomain "*" --rpcaddr 0.0.0.0 --rpcport 8501 --rpcvhosts "*"   --rpcapi "db,eth,net,web3,personal,debug" --gcmode "archive" --identity "SwapDex Whale" --etherbase "coinbase address here" --unlock "coinbase account here" --password pass.txt console
```
Enter

## Congratulations!! you have run your own Masternode.

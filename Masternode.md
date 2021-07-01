# SwapDEX Masternode Guide

This guide will explain the steps required to set up a Masternode for the SmartDEX Chain.

## Requirements for running a masternode:-

The following are the <b> minimum recommended requirements </b> to run a Masternode smoothly. If you use lower capacity device or VPS, we cannot guarantee your Masternode will run smoothly and you may not receive the full allotment of rewards. Using a device or VPS with higher specifications is fine, but is not required.

1) Device/VPS running Ubuntu operating system with the following minimum specifications:

* 16GB RAM
* 120GB SSD
* 6 Core CPU

2) A MetaMask wallet with at least 10,010 SDX in the wallet. While you only need 10,000 SDX to stake, we recommend the wallet holding a little more to cover gas fees.

# You are now ready to start installation process

This guide is written for the Ubuntu Operating System version 20. All commands should be performed in the terminal application - the shortcut to open terminal is ``Ctrl + Alt + T``.
For those who are unfamiliar to using terminal commands in Ubuntu, you should write the commands written ``in these text boxes`` exactly as they are written in the guide and at the end of each line press enter.
Where entries are provided in ``[Square Brackets]`` the description in the square bracket gives an indication of what you should enter, but the entry will be unique to each user. For these entries, the square brackets should not be included in your entry.
Where entries are provided in ``"quotation marks"`` you need to keep the quotation marks.
There will be commands which require you confirm you want to proceed and times you are prompted to enter your password.
This guide assumes you will confirm all processes and enter your password as required when prompted.

To copy within Ubuntu press ``Ctrl + Shift + C`` and to paste press ``Ctrl + Shift + V``.

## Install tmux, Go & Gvm

To run a Masternode, you will need tmux, Go & GVM.
For full details on GVM, you can refer to https://github.com/moovweb/gvm however this guide includes all the steps you will need.
Start by opening terminal and entering the following 

```bash 
sudo apt update
sudo apt upgrade
sudo apt install git
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer) 
```
Restart the terminal (close terminal and then reopen). Once the terminal is back up, get back to the action.

```bash
sudo apt-get install curl git mercurial make binutils bison gcc build-essential
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.13.8
gvm use go1.13.8
sudo apt-get install tmux
```

## Lets start building!
```bash
git clone https://github.com/69th-byte/SmartDex-Chain.git sdxchain
cd sdxchain 
make all
sudo mv build/bin/sdx  /usr/bin/sdx
```

## Create your Masternode Account
Now you have the required software installed, you are ready to set up your Masternode Account. 

```bash
sdx account new
[ENTER SECURE PASSWORD - you will need this again soon]
```
Take a note of the Address that is created (we recommend copying it into a separate file for ease of reference, you will need it on two further occasions when setting up your Masternode), this is your *SDX Masternode Server Account Address* which you will need shortly and which you will need to connect your wallet to your Masternode.
**Please note you need to add 0x to the start of the SDX Masternode Server Account Address**

```bash
sdx init swapdex.json
nano pass.txt
```

## Save your SDX Masternode Account password
You will now need to enter the Password which you used to create server generated SDX Masternode Account and then you will save and exit
```bash
[ENTER SECURE PASSWORD]
ctrl + x & y
```
Press Enter to exit nano.

## Open tmux to start running your Masternode
If you do not run your masternode within tmux, it will stop running when you exit and you will not receive rewards.
```bash
tmux
```

## Run the script to run your Masternode
Please note there are a number of entries in the code below where you will need to provide unique information. Please ensure all lines of code are entered in the one command.
If you enter tmux and forgot to copy your SDX Masternode Server Account Address you can enter the command ``exit`` to end the tmux session to find your SDX Masternode Account address.

``
sdx  --syncmode "full" --networkid 7879 --port 10303 --rpc --rpccorsdomain "*" --rpcaddr 0.0.0.0 --rpcport 8501 --rpcvhosts "*"   --rpcapi "db,eth,net,web3,personal,debug" --gcmode "archive" --identity "[Name for node]" --etherbase "0x[SDX Masternode Server Account Address]" --unlock "0x[SDX Masternode Server Account Address]" --mine --gasprice 2500 --password pass.txt console
``

## Congratulations, you have now set up your own Masternode!

### But wait, what's next?
You have now set up the Masternode software, but before your node starts providing you rewards, you need to go to the SDX Master site to stake your 10,000 SDX and attach your wallet to your node so you can start receiveing rewards.

These details will be updated once SDXMaster is finished...

# SwapDEX Masternode Guide (Draft)

This guide will explain the steps required to set up a Masternode for the SmartDEX Chain.

## Requirements for Running a Masternode

The following are the <b> minimum recommended requirements </b> to run a Masternode smoothly. If you use lower capacity device or VPS, we cannot guarantee your Masternode will run smoothly and you may not receive the full allotment of rewards. Using a device or VPS with higher specifications is fine, but is not required.

1) Device/VPS running Ubuntu operating system with the following minimum specifications:

* 16GB RAM
* 120GB SSD *(higher the better - this size will increase as the chain ages so it is worth considering getting more storage to future-proof the storage requirements)*
* 6 Core CPU

2) A MetaMask wallet with at least 10,001 SDX in the wallet. While you only need 10,000 SDX to stake, we recommend the wallet holding a little more to cover gas fees.

# You are now Ready to Start Installation Process

This guide is written for the Ubuntu Operating System version 20. All commands should be performed in the terminal application - the shortcut to open terminal is ``Ctrl + Alt + T``.
For those who are unfamiliar to using terminal commands in Ubuntu, you should write the commands written ``in these text boxes`` exactly as they are written in the guide and at the end of each line press enter.
Where entries are provided in ``[Square Brackets]`` the description in the square bracket gives an indication of what you should enter, but the entry will be unique to each user. For these entries, the square brackets should not be included in your entry.
>Examples:
>* identity "\[Name for node]\"
>   * identity "MyNode"
>* etherbase 0x\[MSAA]
>   * etherbase "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B"
>* unlock 0x\[MSAA]
>   * unlock "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B"
   
Where entries are provided in ``"quotation marks"`` you need to keep the quotation marks.
This guide assumes you are logged in as root.
There will be commands which may require you confirm you want to proceed and times you are prompted to enter your password. This guide assumes you will confirm all processes and enter your password as required when prompted.

To copy within Ubuntu press ``Ctrl + Shift + C`` and to paste press ``Ctrl + Shift + V``.

## Install All Required Programs

To run a Masternode, you will need a number of programs including tmux, Go & GVM.
For full details on GVM, you can refer to https://github.com/moovweb/gvm however this guide includes all the steps you will need.
Start by opening terminal and entering the following 

```bash 
apt-get update && apt-get upgrade -y
apt-get install curl git nano mercurial make binutils bison gcc build-essential tmux
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer) 
```
Restart the terminal (close terminal and then reopen). Once the terminal is back up, get back to the action.

```bash
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.13.8
gvm use go1.13.8
```

## Let's Start Building!
```bash
git clone https://github.com/69th-byte/SmartDex-Chain.git sdxchain
cd sdxchain 
make all
mv build/bin/sdx  /usr/bin/sdx
```

## Create your Masternode Account
Please note you will not see the cursor move while you enter your password. This is normal.
```bash
sdx account new
[ENTER A SECURE PASSWORD - you will need this again soon]
```
Take a note of the 40 character address that is created (we recommend copying it into a separate file for ease of reference, you will need it on two further occasions when setting up your Masternode), this is your server generated *SDX Masternode Server Account Address (MSAA)* which you will need shortly and which you will need to connect your wallet to your Masternode.
Please note, you need to **_add 0x_ to the start of the MSAA**.

```bash
sdx init swapdex.json
nano pass.txt
```

## Save your SDX Masternode Account Password
You will now need to enter the Password which you used to create the MSAA.
```bash
[ENTER YOUR SECURE PASSWORD]
ctrl + x & y
```
Press Enter to exit nano.

## Open tmux and Start Running your Masternode
If you do not run your masternode within tmux, it will stop running when you exit and you will not receive rewards. Once you are running your masternode within tmux, you can exit tmux and it will continue running. If you need to stop the miner from running, you can close it by opening tmux and entering ``Ctrl + d``.
If you enter tmux and forgot to copy your MSAA you can enter the command ``exit`` to end the tmux session to find your MSAA address.
```bash
tmux new -s SDX
```
Please note there are a number of entries in the code below in ``[Square Brackets]`` where you will need to provide unique information. Don't forget to delete the ``[Square brackets]`` and if you added the 0x to your MSAA, do not include the 0x twice at the etherbase and unlock address.
We have changed the way the code displays below so you can see the full code on screen, however please ensure **all lines of code below are entered in one command**.

``sdx  --syncmode "full" --networkid 7879 --port 10303 --rpc --rpccorsdomain "*" --rpcaddr 0.0.0.0 --rpcport 8501 --rpcvhosts "*" --rpcapi "db,eth,net,web3,personal,debug"  --identity "[Name for node]" --etherbase "0x[MSAA]" --unlock "0x[MSAA]" --mine --gasprice 2500 --password pass.txt console
``

## Finish Setup on the SDX Master Site
You need to register your masternode on the SDX Master site *(to be provided on final version of guide)* in order for it to be activated on the Smart DEX Chain.
When you start your masternode using the code from the previous step, it will start syncing with the chain. This means it will start downloading all previous transactions so it has a full and accurate record of the ledger. Your masternode needs to be synced with the live chain in order for your registration to proceed successfully. If you register your masternode before it has synced, it will get slashed for 5 epochs. Providing it has synced by the end of hose 5 epochs, it will then activate. In order to see the progress of your sync, cross reference the current block (shows as **number** in the mining script) with the Current Block on the landing page for the SDX Master website.

## Register your Masternode & Stake your SDX
To register your Masternode, visit WEBSITE and complete the following steps:
* Connect your MetaMask account to the website by clicking the <b>Login</b> button. Before you log in, please ensure your MetaMask account is connected to the Smart DEX Chain and the account you wish your masternode rewards to go to is active. This account will need to hold at least 10,001 SDX. You can register multiple masternodes against one wallet address.
* On the login screen, select MetaMask then click <b>Save</b>. If it's the first time you are logging in, you may need to grant access in MetaMask.
* Click on <b>Register a Masternode</b>.
* The Stake field will automatically be filled with 10,000 SDX. This amount is independent from the standard staking so you gain no benefit from staking a higher amount than 10,000 SDX at this point so we do not recommend you increase this figure.
* Enter your MSAA (including 0x at the start) in the SDX Masternode Address field then click <b>Apply</b>.
* Confirm the transaction in MetaMask.

# Congratulations, you have now set up your own Masternode!
Once your masternode is active, you rewards will appear in your wallet automatically at the end of each epoch where your masternode was selected to mine.
If your masternode status is Slashed, this means it has been rejected by the chain.
A masternode can be Slashed for a number of reasons such as incorrect information being provided when registering, an error in the code to run the masternode, the code to run the masternode being activated before the MSAA has been registered on the SDX Master or the device/VPS having insufficient specifications to keep up with the requirements of the chain.
If the cause of the slashing is something that you can fix, there is no need to restart the whole process, you can simply correct the issue and the masternode will kick back in once the chain registers the issue has been corrected. This can take a few epochs, so don't panic if it takes a few hours for the status to update.

Good luck!

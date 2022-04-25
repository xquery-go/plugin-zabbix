# Plugin-Zabbix-FE

## Zabbix version 5.0
## Add plugin into Zabbix

First, clone the repo of Zabbix :

```sh
git clone https://git.zabbix.com/scm/zbx/zabbix.git --depth 1 zabbix-agent2
```

After, clone this repo :
```sh
git clone https://gitlab.ypsi.cloud/melissa.bertin/plugin-zabbix-fe.git
```

Finally, copy the directory flexibleengine into the **zabbix-agent2** directory:
```sh
cp -r plugin-zabbix-fe/flexibleengine zabbix-agent2/src/go/plugins/flexibleengine
```

And modify the three files in zabbix-agent2/src/go/plugins (**plugins_linux.go**, **plugins_windows.go**, **plugins_darwin.go**) by adding this line in import list:

_ "zabbix.com/plugins/flexibleengine"


## Build Agent

First, go to the zabbix source
```sh
cd zabbix-agent2
```

Run this command to build the agent with the new plugin:
```sh
./bootstrap.sh; ./configure --enable-agent2 --enable-static; make
```

## Modify agent 2

If the zabbix-agent2 isn't already running use this doc: https://docs.google.com/document/d/1Yvqh_r_nXwbnwLPYnbtdnJ5UOofiR6f5GvoOOrUdjJM/edit?usp=sharing
If the zabbix-agent2 is already running make this three commands to modify agent2:
```sh
systemctl stop zabbix-agent2
cp <zabbix-source>/src/go/bin/zabbix_agent2 /usr/sbin/
systemctl start zabbix-agent2
```

## Use the new plugin

I describe the procedure for the **NAT template** but is the same procedure for other element.
First, go to your Zabbix application in *Configuration > Templates*. Choose the *button import* in upper right of your screen. Select a template **Cloud-FlexibleEngine-NAT.xml** file in the **templates** directory.

After, you must change **MACROS values** in the template for this go to *MACROS menu* and set your value for {$ACCESS_KEY}, {$PROJECT_ID}, {$SECRET_KEY}. At this, moment don't set the {$INSTANCE_ID} value.

Once your template created, you can create a new host in *Configuration > Hosts*. In *Templates menu* choose **Cloud-FlexibleEngine-NAT** template and in *Macros menu*, go to *Inherited and host macros* and modify the value for {$INSTANCE_ID} with your **NAT ID**. 

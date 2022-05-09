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

For template EVS, you must define {$INSTANCE_ID} value which is the ID of the EVS and define {$DISK_NAME} which is the ID of the ECS with the of disk device after like this: 93503d16-53a0-41ec-985f-ae6eee18a3b6-vda

For template DDS, you must define {$INSTANCE_ID} value which is the ID of the DDS and define {$ROLE} which is the role of the DDS (primary or secondary)

For template DCS, you must define {$INSTANCE_ID} value which is the ID of the DCS and define {$ENGINE} which is the engine of the DCS
For template OBS, you must define {$BUCKET_NAME} value which is the name of the OBS

## Use discovery plugin
To begin, create a host group which has the domain name in FE. <br> Import the template "Cloud-FlexibleEngine-Discovery" and all templates of object you want. 

After, create one host by project. For this, create host and give a name. <br>
Add this host to the host group created previously.<br>
Link the template "Cloud-FlexibleEngine-Discovery" to this host. <br>
Add agent interface.<br>
In tags section, define two tags. The first is to define the name of a project (name: project ; value: NAME_OF_YOUR_PROJECT). The second is to define the region of a projet (name: region ; value: NAME_OF_YOUR_REGION).<br>
Finally, in macros section, you must define values of each inherited macros : 
* {$ACCESS_KEY} : The access key to connect to FE
* {$SECRET_KEY} : The secret key to connect to FE
* {$DOMAIN_NAME} : The domain name on FE
* {$PROJECT_ID} : The project ID where elements are stored
* {$PROJECT_NAME} : The project name where elements are stored
* {$REGION} : The name of the region
* {$TOKEN_API} : The token API generate in Zabbix (Administration > General > Token API)
* {URL_ZABBIX} : The zabbix URL to make request API (example: http://SERVER_IP)

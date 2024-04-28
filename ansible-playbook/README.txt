This document specifies how to configure and run ansible automation playbooks to deploy stemcell-GO server on AWS

Pre-requisite:
1. Active AWS account wit IAM user to create EC2 instances.
2. Installed python, pip.
3. Install boto, boto3 and ansible suing pip on local machine to facilitate ansible communication with AWS. 

Generate ssh keys to be used for SSH login to EC2 instance.
	ssh-keygen -t rsa -b 4096 -f ~/.ssh/my_aws


Setup ansible vault to store AWS access key and secret key used for communication
	1. create vault password file to be used to run playbooki or edit pass.yml. pass.yml will store all ansible secrets i.e. aws_access_key and aws_secret_key	
		openssl rand -base64 2048 > vault.pass
	2. create password yml file (subtree to pwd)
		ansible-vault create group_vars/all/pass.yml --vault-password-file vault.pass
	This should open editor add access key and secret key here
		ec2_access_key: AAAAAAAAAAAAAABBBBBBBBBBBB                                      
		ec2_secret_key: afjdfadgf$fgajk5ragesfjgjsfdbtirhf
	3. edit pass.yml (optional)
		ansible-vault edit group_vars/all/pass.yml --vault-password-file vault.pass

Create inventory file for storing host information
	touch inventory
Hosts contianed in this inventory file will be used to spawn stemcell-GO server(s) 


Run ansible playbook ec_instance_create.yml
	ansible-playbook -i ./inventory ec2_instance_create.yml --tags create_ec2 --vault-password-file vault.pass


After successful execution of the above playbook, an EC2 instance should be in "Running" state on your AWS account.


To start stemcell-GO server on the created instance, run ec2_instance_config.yml playbook
	ansible-playbook -i ./inventory ec2_instance_config.yml --vault-password-file vault.pass


After successful execution of the above playbook, the server should be accessible on http://<ec2-public-ip>:8080.


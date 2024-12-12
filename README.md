# FaceMatch AI

## How it works

First step: you provide your Instagram account. From here, all your 1st degree connections will be scanned using the LinkedIn API. You will be required to sign in using LinkedIn so that the APU has access to your profile.

Then, your connections' profile pictures and yours will be serailized, creating embeddings which are used for the face matching. The image is standardized to a 1x1 480px format and converted to base64. This value is sent to AWS Bedrock using the Titan model.

This embedding and basic profile data is added to the database.

## Deploying in AWS

Firstly, ensure you have the AWS CLI installed and have configured your IAM credentials. Also have a VPC, subnet, and security group ready for use.

```sh
aws ec2 describe-images --owners amazon # select desired AMI
aws ec2 create-key-pair --key-name FMAI_KeyPair --query '<PASSWORD>' --output text > KeyPair.pem # create key-pair for SSH connection

aws ec2 create-security-group --group-name MySecurityGroup --description "My security group" --vpc-id vpc-fmai-main
aws ec2 authorize-security-group-ingress --group-id sg-fmai-main --protocol tcp --port 80 --cidr 0.0.0.0/0

# TODO: create subnet

aws ec2 run-instances --image-id '<IMAGE_ID>' --count 1 --instance-type t2.micro --key-name MyKeyPair --security-group-ids sg-fmai-main --subnet-id subnet-fmai-main
aws ec2 describe-instances # check if it worked; optional --instance-ids param

# TODO:
# - connect to EC2 instance via SSH
# - clone git repo
# - configure application
# - run server

# TODO: automate the last process with a user data script
```

Remember to follow security best practices:
- Use the principle of least privilege when setting up IAM roles and security group rules.
- Regularly update and patch your EC2 instance.
- Use secure methods for storing and accessing your Git credentials.

For detailed CLI commands and up-to-date best practices, I recommend consulting the official AWS documentation. Also, consider using AWS CodeDeploy for more robust application deployment options.
Sources
[1] [Create an Amazon EC2 instance for CodeDeploy (AWS CLI or Amazon EC2 console) - AWS CodeDeploy] (https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-ec2-create.html)
[3] [How to connect to a private EC2 instance from a local Visual Studio Code IDE with Session Manager and AWS SSO (CLI) | AWS re:Post] (https://repost.aws/articles/AR8Gk1UngsTpmpu7azMUiNvw/how-to-connect-to-a-private-ec2-instance-from-a-local-visual-studio-code-ide-with-session-manager-and-aws-sso-cli)

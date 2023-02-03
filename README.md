# Word Scrambler

This program is intended to learn about *goroutines** and **AWS SQS** to handle multi-threads and communication between different processes.
There are two main programs:

- **Scrambler:** Will randomize the position of the word in the phrase and push it into the AWS queue.

- **Assember:** Will fetch every word and it will assemble back to the original phrase

## Required:

AWSCLIv2 Needs to be installed before using this feature, because we need to have your AWS credentials configured in order to acess **AWS SQS**:

```shell
	$ aws configure
```


## Usage:

- Provision SQS Queue in AWS as follows, a queue named `phrase-scrambler-queue` will be created:

```shell
	$ cd cdk-deploy 
	$ cdk synth 	# verify how the queue will be created
	$ cdk deploy	# to provision the queue
```

- Once the queue is provisioned successfully, run the following commands in separate windows to see how consumer and producer processes are running.

```shell
	$ make scrambler # to scramble the Phrase
	$ make assember # to rebuild the Phrase
```


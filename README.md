# **Phrase Scrambler** ![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)


The initial service will scramble a phrase by rearranging its words and then send each word to a queue. A different service will then fetch each word from the queue and unscramble them, restoring the original phrase.


This program is intended to learn about *goroutines* and **AWS SQS** to handle multi-threads and communication between different processes.
There are two main components:

- **Scrambler:** Randimizes the position of the words in the phrase and pushes it into the AWS queue.

- **Assember:** Fetches every word from the queue and reassembles them into the original phrase

## **Prerequisites:**

- **AWSCLIv2** Needs to be installed before using this feature, because we need to have your AWS credentials configured in order to acess **AWS SQS**:

```shell
	$ aws configure
```

- You need to install **CDKv2** in your system

```shell
$ npm install -g aws-cdk
$ cdk --version # Make sure the version (2+) is installed
```

- Follow the [CDKv2](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html) guide for more information

## **Usage:**

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


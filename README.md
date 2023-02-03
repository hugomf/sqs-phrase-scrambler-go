# Word Scrambler

This program is intended to learn about goroutines and sqs to handle multi-threads and communication between different processes.
There are two main programs:

- Scrambler: Will randomize the position of the word in the phrase and push it into the AWS queue.

- Assember: Will fetch every word and it will assemble back to the original phrase

## Usage:

- Provision SQS Queue in AWS as follows, a queue named `phrase-scrambler-queue` will be created:

```shell
	$ cd cdk-deploy 
	$ cdk deploy
```

- Run the following commands in separate windows to see how the consuming and producing processes are running.

```shell
	$ make scrambler # to scramble the Phrase
	$ make assember # to rebuild the Phrase
```


# Word Scrambler

This program is intended to learn about goroutines and sqs to handle multi-threads and communication between different processes.

## Usage:

```shell
	$ make scrambler # to scramble the Phrase
	$ make assember # to rebuild the Phrase
```

- Scrambler: Will randomize the position of the word in the phrase and push it into the AWS queue.

- Assember: Will fetch every word and it will assemble back to the original phrase
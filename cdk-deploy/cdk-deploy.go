package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkDeployStackProps struct {
	awscdk.StackProps
}

func NewCdkDeployStack(scope constructs.Construct, id string, props *CdkDeployStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// producerQueue := awssqs.NewQueue(stack, jsii.String("phrase-producer-queue"), &awssqs.QueueProps{
	// 	QueueName:         jsii.String("phrase-producer-queue"),
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	// awscdk.NewCfnOutput(stack, jsii.String("PhraseProducerQueueUrl"), &awscdk.CfnOutputProps{
	// 	Value:      producerQueue.QueueUrl(),
	// 	ExportName: jsii.String("PhraseProducerQueueUrl"),
	// })

	scramblerQueue := awssqs.NewQueue(stack, jsii.String("phrase-scrambler-queue"), &awssqs.QueueProps{
		QueueName:         jsii.String("phrase-scrambler-queue"),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	})

	awscdk.NewCfnOutput(stack, jsii.String("PhraseScramblerQueueUrl"), &awscdk.CfnOutputProps{
		Value:      scramblerQueue.QueueUrl(),
		ExportName: jsii.String("PhraseScramblerQueueUrl"),
	})

	// fmt.Print()

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkDeployStack(app, "CdkDeployStack", &CdkDeployStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}

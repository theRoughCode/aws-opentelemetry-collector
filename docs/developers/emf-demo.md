### Using AWS-Observability-Collector to send Embedded Metrics Format Logs

This example will introduce how to use AOC Beta to send EMF logs with metrics on AWS ECS/EKS. This instruction provided the data emitter image that will generate OTLP format of metrics and traces data to AWS CloudWatch and X-Ray consoles.  Please follow the steps below to try AWS Observability Collector Beta.

#### Steps,
**Prerequisite**  
AOC docker image is not yet available for Beta release. You may have to follow [build docker](build-docker.md) instruction to build your own AOC docker image to run the example.

#### Create ECS-AWSObservability IAM Policy 
1. Open the IAM console at https://console.aws.amazon.com/iam/.
2. In the navigation pane, choose **Policies**.
3. Choose **Create policy, JSON**.
4. Enter the following policy:
5. Choose Review policy.
6. On the Review policy page, enter ```ECS-AWSObservability``` for the Name and choose Create policy.
```{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:DescribeLogStreams",
                "xray:*",
                "s3:*"
            ],
            "Resource": [
                "arn:aws:logs:*:*:*"
            ]
        }
    ]
}
```

#### Install AOC on Amazon ECS as sidecar
The easiest way to deploy AOC on Amazon ECS is to run it as a sidecar, defining it in the same task definition as your application.

**Configure The Task Definition**  
The following two examples are provided to run AWS-Observability-Collector as sidecar to send application metrics and traces on Amazon ECS.  

Notice that in the task definition templates we provided, you have to fill ```{{aocImage}}``` with the aoc docker image you created. Replace ```{{region}}``` with the region your want to send the data to. Fill ```{{ecsTaskRoleArn}}``` and ```{{ecsExecutionRoleArn}}``` with the IAM role you created above.
You can find ECS Fargate AOC [Installing template](https://github.com/wyTrivail/aocintegtest/blob/master/src/main/resources/templates/ecs/aoc-sidecar-fargate.mustache).

Another example is installing AOC on ECS EC2 instance as sidecar to send telemetry data.
You can find ECS EC2 AOC [Installing template](https://github.com/wyTrivail/aocintegtest/blob/master/src/main/resources/templates/ecs/aoc-sidecar-fargate.mustache).


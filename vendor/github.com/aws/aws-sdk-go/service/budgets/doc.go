// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package budgets provides the client and types for making API
// requests to AWS Budgets.
//
// Budgets enable you to plan your service usage, service costs, and your RI
// utilization. You can also track how close your plan is to your budgeted amount
// or to the free tier limits. Budgets provide you with a quick way to see your
// usage-to-date and current estimated charges from AWS and to see how much
// your predicted usage accrues in charges by the end of the month. Budgets
// also compare current estimates and charges to the amount that you indicated
// you want to use or spend and lets you see how much of your budget has been
// used. AWS updates your budget status several times a day. Budgets track your
// unblended costs, subscriptions, and refunds. You can create the following
// types of budgets:
//
//    * Cost budgets allow you to say how much you want to spend on a service.
//
//    * Usage budgets allow you to say how many hours you want to use for one
//    or more services.
//
//    * RI utilization budgets allow you to define a utilization threshold and
//    receive alerts when RIs are tracking below that threshold.
//
// You can create up to 20,000 budgets per AWS master account. Your first two
// budgets are free of charge. Each additional budget costs $0.02 per day. You
// can set up optional notifications that warn you if you exceed, or are forecasted
// to exceed, your budgeted amount. You can have notifications sent to an Amazon
// SNS topic, to an email address, or to both. For more information, see Creating
// an Amazon SNS Topic for Budget Notifications (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-sns-policy.html).
// AWS Free Tier usage alerts via AWS Budgets are provided for you, and do not
// count toward your budget limits.
//
// Service Endpoint
//
// The AWS Budgets API provides the following endpoint:
//
//    * https://budgets.amazonaws.com
//
// For information about costs associated with the AWS Budgets API, see AWS
// Cost Management Pricing (https://aws.amazon.com/aws-cost-management/pricing/).
//
// See budgets package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/budgets/
//
// Using the Client
//
// To contact AWS Budgets with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Budgets client Budgets for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/budgets/#New
package budgets
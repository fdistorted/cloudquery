---
title: Keyless access to AWS in GitHub Actions with OIDC
tag: security
date: 2022/03/28
description: A guide to configuring OpenID Connect access to AWS from GitHub Actions
author: mikeelsmore
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


With this blog, we will show you how to access your AWS environment without storing IAM credentials in GitHub by using OpenID Connect (OIDC).



## What is OpenID Connect?

OpenID Connect has been around since 2014, and in reality, it’s a simplified identity layer on top of the OAuth 2.0 protocol. This gives your clients and applications an easy process to authenticate against an Authorization Server and get information about the end-users (in this case AWS) but in a process flow that’s easier for applications to do so. For a deeper look into OpenID Connect check out [https://openid.net/connect/](https://openid.net/connect/).

So using OpenID Connect we can remove the need to have keys stored in GitHub Actions, saving the headache of [rotating the keys](https://github.com/cloudquery/cq-provider-aws/tree/main/policies/cis_v1.2.0) or the other headaches.

## Setting up AWS OpenID Connect Identity Provider

The first step to this whole process is to configure AWS to allow GitHub to communicate with it via OpenID Connect. Here are a few different ways to do it.

We've got three different ways to do this [AWS Console](#aws-console), [AWS CLI](#aws-cli), and [Terraform](#terraform).

### AWS Console

Once you are logged into the AWS Console, head to IAM and select `Identity providers`:

![Identity providers](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image4.png)

Select the `Add provider` button in the top right corner

![Add provider](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image5.png)

On the `Add an Identity provider` screen, you will want to select `OpenID Connect` as the `Provider type`, and then add the following information to the fields

- Provider URL: `https://token.actions.githubusercontent.com`
- Audience: `https://github.com/(org name)` (so in our case `https://github.com/cloudquery`)

![Add an Identity provider](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image6.png)

Make sure to select `Get thumbprint` to get the certificates, and then `Add provider` in the bottom right.

The next step is to add a role for the provider to assume when the GitHub Action accesses AWS. head to `Roles` within `IAM`:

![Roles](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image2.png)

Select `Create role` in the top right.

![Create role](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image1.png)

Now we need to identify our OpenID Connect provider. Select `Web identity` from the `Trusted entity types`, select the `token.actions.githubusercontent.com` from the `Identity provider` drop-down, and then select the `Audience` you added a moment ago. Then hit `next`.

![Select trusted entity](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image7.png)

The next step is to either assign or create the permissions the role will need, so what permissions do your GitHub Actions need to work within AWS. This is entirely up to your needs, so best to do reading around AWS Roles and service permissions.

![Add permissions](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image3.png)

The final step in creating the Role is to name it and provide a relevant description. You can also add more permissions or even adapt the `trusted entities` ( which Identity providers have access to the role).

![Review role](/images/blog/keyless-access-to-aws-in-github-actions-with-oidc/image8.png)

### AWS CLI

When using the AWS Console, the console itself will attempt to do some of the steps for you, however, when using the AWS CLI you’ll have an initial step to do yourself. This first step is getting the `thumbprint` for the GitHub certificates used to authenticate the requests, that process is an entire blog itself so I would suggest following the guide provided by AWS called [Obtaining the thumbprint for an OpenID Connect Identity Provider](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html).

Once you have the `thumbprint`, that should look something like `15E29108718111E59B3DAD31954647E3C344A231` we can get on to actually creating the Identity provider. The command to create the OpenID Connect Identity provider is as follows:

```bash
aws iam create-open-id-connect-provider --url https://token.actions.githubusercontent.com --client-id-list https://github.com/(org name) --thumbprint-list 15E29108718111E59B3DAD31954647E3C344A231
```

Using the IAM suite of commands we run `create-open-id-connect-provider` and provide it with a few pieces of information.

- `--url [https://token.actions.githubusercontent.com](https://token.actions.githubusercontent.com/)` which is the provider URL
- `--client-id-list https://github.com/(org name)` is the `Audience` in the console, and essentially is who can use it so `https://github.com/cloudquery` for us
- `--thumbprint-list 15E29108718111E59B3DAD31954647E3C344A231`

This will give you a response object that looks like this:

```json
{
  "OpenIDConnectProviderArn": "arn:aws:iam::000000000000:oidc-provider/token.actions.githubusercontent.com"
}
```

Now the OpenID Connect Identity provider exists, you need to create a role for it to assume to access AWS services with a trust policy and permissions.

Let’s start with the trust policy `githubactionawstrustpolicy.json`

```json
{
  "Version": "2012-10-17",
  "Statement": {
    "Sid": "RoleForGitHubActions",
    "Effect": "Allow",
    "Principal": {
      "Federated": "arn:aws:iam::000000000000:oidc-provider/token.actions.githubusercontent.com"
    },
    "Action": "sts:AssumeRoleWithWebIdentity",
    "Condition": {
      "StringLike": {
        "token.actions.githubusercontent.com:sub": "repo:github.com/(org name)"
      }
    }
  }
}
```

This trust policy is allowing a connection from GitHub via the OpenID Connect we created earlier to assume the role, and to set this up we use the following command:

```bash
aws iam create-role --role-name GitHub-Action-Role --assume-role-policy-document <pathto>/githubactionawstrustpolicy.json
```

Using the `aws iam create-role` command we are creating the `GitHub-Action-Role` and pointing it to our trust policy. And once this role exists, we need to give permissions to actually access AWS Services.

Here is an example permissions JSON in the form of `permissionsforgithubactions.json`

```json
{
  "Version": "2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": "s3:ListBucket",
    "Resource": "arn:aws:s3:::example_bucket"
  }
}
```

The permission we are granting here is that the GitHub Action may run `ListBuckets` on only the `example_bucket` S3 bucket.

To attach this to our new role you will need to run the command:

```bash
aws iam put-role-policy --role-name GitHub-Action-Role --policy-name Perms-Policy-For-GitHub-Actions --policy-document <pathto>/permissionsforgithubactions.json
```

With this command, we’re attaching the `Perms-Policy-For-GitHub-Actions` to our `GitHub-Action-Role` with the `ListBucket` action.

### Terraform

If you’re using [Terraform](https://www.terraform.io/) already, it’s quite a straightforward terraform script to create an OpenID Connect Identity provider with its required role and permissions. Luckily, the Terraform registry contains a complete [AWS provider](https://registry.terraform.io/providers/hashicorp/aws/latest) to use. The script is as follows:

```hcl
provider "aws" {}

resource "aws_iam_openid_connect_provider" "default" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "https://github.com/(org name)",
  ]

  thumbprint_list = [
    "15E29108718111E59B3DAD31954647E3C344A231",
  ]
}

resource "aws_iam_role" "github_action_role" {
  name = "github_actions_role"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRoleWithWebIdentity",
        Effect = "Allow"
        Sid    = "RoleForGitHubActions",
        Principal = {
          Federated = aws_iam_openid_connect_provider.default.arn
        }
        Condition = {
          StringLike = {
            "token.actions.githubusercontent.com:sub" = "repo:github.com/(org name)"
          }
        }
      },
    ]
  })
}

resource "aws_iam_role_policy" "github_actions_policy" {
  name = "github-actions-policy"
  role = aws_iam_role.github_action_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:Describe*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}
```

Let’s break this down, as we’re using `AWS` you will need to make sure you have `provider "aws" {}` configured to actually interact with the AWS account in question.

The next block is actually configuring the OpenID Connect Identity provider

```hcl
resource "aws_iam_openid_connect_provider" "default" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "https://github.com/(org name)",
  ]

  thumbprint_list = [
    "15E29108718111E59B3DAD31954647E3C344A231",
  ]
}
```

The individual parts of this are

- `url = “https://token.actions.githubusercontent.com”` which is the provider URL
- `client_id_list = ["https://github.com/(org name)"]` is the `Audience` in the AWS console, and essentially is who can use it so `https://github.com/cloudquery` for us
- `thumbprint_list = ["15E29108718111E59B3DAD31954647E3C344A231"]` which is the certificates that OpenID Connect uses to authenticate. The process is an entire blog itself so I would suggest following the guide provided by AWS called [Obtaining the thumbprint for an OpenID Connect Identity Provider](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html).

Once we have an OIDC we will need to grant it a role to assume within AWS, the block for this is:

```hcl
resource "aws_iam_role" "github_action_role" {
  name = "github_actions_role"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRoleWithWebIdentity",
        Effect = "Allow"
        Sid    = "RoleForGitHubActions",
        Principal = {
          Federated = aws_iam_openid_connect_provider.default.arn
        }
        Condition = {
          StringLike = {
            "token.actions.githubusercontent.com:sub" = "repo:github.com/(org name)"
          }
        }
      },
    ]
  })
}
```

Here we are simply naming the role `github_actions_role` and then giving it a role policy, once we have this role policy the final block is used to grant said role permissions:

```hcl
resource "aws_iam_role_policy" "github_actions_policy" {
  name = "github-actions-policy"
  role = aws_iam_role.github_action_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:Describe*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}
```

This permissions block is allowing the GitHub Action to Describe all EC2, which could be handy.

## Setting up GitHub to use the OpenID Connect provider

Now that we have a working OpenID Connect provider within AWS, we need to add the configuration to GitHub for use in our GitHub Actions. To do this, we simply add another step to the desired `yml` workflow.

The relevant blocks look like so:

```yaml
permissions:
  id-token: write
  contents: read # This is required for actions/checkout@v2
```

Adding the `permissions` to the job allows the action that gets the credentials from AWS to store them for use in further steps. The permission that is specifically required is `id-token: write`.

The next `step` is where the credential retrieving magic actually happens

```yaml
    steps:
      […]
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v1
        with:
          role-to-assume: arn:aws:iam::000000000000:role/cq-provider-aws-github-action
          aws-region: us-east-1
      […]
```

An AWS action already exists, so we point to `aws-actions/configure-aws-credentials@v1` then within the `with` section we add the role we want it to assume (like `arn:aws:iam::000000000000:role/GitHubActionRole`), which AWS region it operates within, and we can optionally add the Role session name.

## Summary

Now, no matter which way you’ve chosen to configure the AWS component your GitHub Actions are able to communicate without the need for stored keys and the access can be managed entirely on the AWS side of things.

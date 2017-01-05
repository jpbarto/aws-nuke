package resources

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetListers(sess *session.Session) []ResourceLister {
	var (
		autoscaling = AutoScalingNuke{autoscaling.New(sess)}
		ec2         = EC2Nuke{ec2.New(sess)}
		elb         = ElbNuke{elb.New(sess)}
		iam         = IamNuke{iam.New(sess)}
		rds         = RDSNuke{rds.New(sess)}
		route53     = Route53Nuke{route53.New(sess)}
		s3          = S3Nuke{s3.New(sess)}
		ecr         = ECRNuke{ecr.New(sess)}
		elasticache = ElasticacheNuke{elasticache.New(sess)}
	)

	return []ResourceLister{
		autoscaling.ListGroups,
		ec2.ListAddresses,
		ec2.ListCustomerGateways,
		ec2.ListDhcpOptions,
		ec2.ListInstances,
		ec2.ListInternetGatewayAttachements,
		ec2.ListInternetGateways,
		ec2.ListKeyPairs,
		ec2.ListNatGateways,
		ec2.ListNetworkACLs,
		ec2.ListRouteTables,
		ec2.ListSecurityGroups,
		ec2.ListSubnets,
		ec2.ListVpcs,
		ec2.ListVpnConnections,
		ec2.ListVpnGatewayAttachements,
		ec2.ListVpnGateways,
		ecr.ListRepos,
		elb.ListELBs,
		iam.ListGroupPolicyAttachements,
		iam.ListGroups,
		iam.ListInstanceProfileRoles,
		iam.ListInstanceProfiles,
		iam.ListPolicies,
		iam.ListRolePolicyAttachements,
		iam.ListRoles,
		iam.ListUserAccessKeys,
		iam.ListUserGroupAttachements,
		iam.ListUserGroupAttachements,
		iam.ListUserPolicyAttachements,
		iam.ListUsers,
		rds.ListInstances,
		rds.ListSnapshots,
		rds.ListParameterGroups,
		rds.ListSubnetGroups,
		route53.ListHostedZones,
		route53.ListResourceRecords,
		s3.ListBuckets,
		s3.ListObjects,
		elasticache.ListCacheClusters,
		elasticache.ListSubnetGroups,
	}
}

package envarAws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
)

func MakeDefaultParameterStoreMapResolver(ctx context.Context, prefix string, recursive bool) envarResolvers.MapResolver {
	return envarResolvers.MapResolver{
		Contents: GetAllParameterStoreItems(ctx, MakeParameterStoreClient(ctx), prefix, recursive),
	}
}

func MakeParameterStoreClient(ctx context.Context) *ssm.Client {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	return ssm.NewFromConfig(cfg)
}

func GetAllParameterStoreItems(ctx context.Context, client *ssm.Client, prefix string, recursive bool) (output map[string]string) {
	output = make(map[string]string)

	response, err := client.GetParametersByPath(ctx, &ssm.GetParametersByPathInput{
		Path:       aws.String(prefix),
		Recursive:  aws.Bool(recursive),
		MaxResults: aws.Int32(1000),
	})

	for {
		if err != nil {
			panic(err)
		}

		for _, parameter := range response.Parameters {
			output[*parameter.Name] = *parameter.Value
		}

		if response.NextToken != nil {
			response, err = client.GetParametersByPath(ctx, &ssm.GetParametersByPathInput{
				Path:       aws.String(prefix),
				Recursive:  aws.Bool(recursive),
				MaxResults: aws.Int32(1000),
				NextToken:  response.NextToken,
			})
		} else {
			return
		}
	}
}

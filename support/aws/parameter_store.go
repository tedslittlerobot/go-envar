package envarAws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	"strings"
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
	// sanitise prefix - ensure that it starts and ends with a slash as per aws api docs
	prefix = fmt.Sprintf("/%s/", strings.Trim(prefix, "/"))

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
			// sanitise paths - removes the prefix from the key
			output[strings.Replace(*parameter.Name, prefix, "", 1)] = *parameter.Value
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
